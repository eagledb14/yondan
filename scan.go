package main

import (
	"context"
	// "fmt"
	"time"
	"strconv"

	"github.com/Ullaakut/nmap/v3"
)

type Scan struct {
	Ip string
	Ports []nmap.Port
	Hostname string
	Timestamp time.Time
}

func NewScan(host nmap.Host, hostname string) Scan {
	return Scan {
		Ip: host.Addresses[0].String(),
		Ports: host.Ports,
		Hostname: hostname,
		Timestamp: time.Now(),
	}
}

//Split the /24 into 16 subnets
//scan each subnet every half hour
//store results in db
func Poll(db *ConcurrentMap) {
	ranges := []string{
		"127.0.0.0/28",
		"127.0.0.16/28",
		"127.0.0.32/28",
		"127.0.0.48/28",
		"127.0.0.64/28",
		"127.0.0.80/28",
		"127.0.0.96/28",
		"127.0.0.112/28",
		"127.0.0.128/28",
		"127.0.0.144/28",
		"127.0.0.160/28",
		"127.0.0.176/28",
		"127.0.0.192/28",
		"127.0.0.208/28",
		"127.0.0.224/28",
		"127.0.0.240/28",
	}
	
	for {
		for _, cidr := range ranges {
			tempMap := make(map[string][]Scan)
			run, _ := nmapScan(cidr)

			for _, host := range run.Hosts {
				if len(host.Addresses) == 0 {
					continue
				}

				// hostname, _ := Lookup(host.Addresses[0].String())
				hostname := ""
				tempMap[host.Addresses[0].String()] = []Scan{NewScan(host, hostname)}

				addPortsIpToMap(&tempMap, NewScan(host, hostname))
				addDomainToMap(&tempMap, NewScan(host, hostname))
				addServiceToMap(&tempMap, NewScan(host, hostname))
			}

			db.MassWrite(&tempMap)
			time.Sleep(30 * time.Minute)
		}
	}
}

func addPortsIpToMap(tempMap *map[string][]Scan, scan Scan) {
	for _, port := range scan.Ports {
		portStr := strconv.Itoa(int(port.ID))
		if _, ok := (*tempMap)[portStr]; !ok {
			(*tempMap)[portStr] = []Scan{}
		} else {
			(*tempMap)[portStr] = append((*tempMap)[portStr], scan)
		}
	}
}

func addDomainToMap(tempMap *map[string][]Scan, scan Scan) {
	if _, ok := (*tempMap)[scan.Hostname]; !ok {
		(*tempMap)[scan.Hostname] = []Scan{}
	} else {
		(*tempMap)[scan.Hostname] = append((*tempMap)[scan.Hostname], scan)
	}
}

func addServiceToMap(tempMap *map[string][]Scan, scan Scan) {
	for _, port := range scan.Ports {
		service := port.Service
		if _, ok := (*tempMap)[service.Name]; !ok {
			(*tempMap)[service.Name] = []Scan{}
		} else {
			(*tempMap)[service.Name] = append((*tempMap)[service.Name], scan)
		}
	}
}

func nmapScan(target ...string) (nmap.Run, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets(target...),
		// nmap.WithCustomArguments()
		// nmap.WithMostCommonPorts(1000),
		// nmap.WithCustomArguments("-p-"),
		// nmap.WithFastMode(),
		nmap.WithPorts("80,443,843"),
	)
	
	if err != nil {
		return nmap.Run{}, err
	}

	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		return nmap.Run{}, err
	}
	if err != nil {
		return nmap.Run{}, err
	}

	return *result, err
}
