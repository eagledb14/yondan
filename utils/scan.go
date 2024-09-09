package utils

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Ullaakut/nmap/v3"
)

type Scan struct {
	Ip string
	Ports []nmap.Port
	Hostname string
	Timestamp string
}

func NewScan(host nmap.Host, hostname string) *Scan {
	return &Scan {
		Ip: host.Addresses[0].String(),
		Ports: host.Ports,
		Hostname: hostname,
		Timestamp: time.Now().Format("2006-01-02"),
	}
}

//Split the /24 into 16 subnets
//scan each subnet every half hour
//store results in db
func Poll(ranges []string, db *ConcurrentMap) {
	tempMap := make(map[string][]*Scan)
	for _, cidr := range ranges {
		run, _ := nmapScan(cidr)

		for _, host := range run.Hosts {
			if len(host.Addresses) == 0 {
				continue
			}

			hostname, _ := Lookup(host.Addresses[0].String())
			hostname = strings.ToLower(hostname)
			tempMap[host.Addresses[0].String()] = []*Scan{NewScan(host, hostname)}

			addPortsIpToMap(&tempMap, NewScan(host, hostname))
			addDomainToMap(&tempMap, NewScan(host, hostname))
			addServiceToMap(&tempMap, NewScan(host, hostname))
			addProtocolToMap(&tempMap, NewScan(host, hostname))
		}

	}
	db.MassWrite(&tempMap)
}

func addPortsIpToMap(tempMap *map[string][]*Scan, scan *Scan) {
	for _, port := range scan.Ports {
		portStr := strconv.Itoa(int(port.ID))
		(*tempMap)[portStr] = append((*tempMap)[portStr], scan)
	}
}

func addDomainToMap(tempMap *map[string][]*Scan, scan *Scan) {
	if scan.Hostname == "" {
		return
	}
	(*tempMap)[scan.Hostname] = append((*tempMap)[scan.Hostname], scan)
}

func addServiceToMap(tempMap *map[string][]*Scan, scan *Scan) {
	for _, port := range scan.Ports {
		service := port.Service
		if service.Name == "" {
			continue
		}
		(*tempMap)[service.Name] = append((*tempMap)[service.Name], scan)
	}
}

func addProtocolToMap(tempMap *map[string][]*Scan, scan *Scan) {
	for _, port := range scan.Ports {
		protocol := port.Protocol
		if protocol == "" {
			continue
		}
		(*tempMap)[protocol] = append((*tempMap)[protocol], scan)
	}
}

func nmapScan(target ...string) (nmap.Run, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets(target...),
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
