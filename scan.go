package main

import (
	"context"
	// "fmt"
	"time"

	"github.com/Ullaakut/nmap/v3"
)

type Scan struct {
	Address nmap.Address
	Ports []nmap.Port
	Hostname string
	Timestamp time.Time
}

func NewScan(host nmap.Host, hostname string) Scan {
	return Scan {
		Address: host.Addresses[0],
		Ports: host.Ports,
		Hostname: hostname,
		Timestamp: time.Now(),
	}
}

//Split the /24 into 16 subnets
//scan each subnet every half hour
//store results in db
func Poll(db *ConcurrentMap[Scan]) {
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
			ipMap := make(map[string]Scan)
			run, _ := nmapScan(cidr)

			for _, host := range run.Hosts {
				if len(host.Addresses) == 0 {
					continue
				}

				hostname, _ := Lookup(host.Addresses[0].String())
				ipMap[host.Addresses[0].String()] = NewScan(host, hostname)
				// fmt.Println(host.Addresses)
			}

			db.MassWrite(&ipMap)
			time.Sleep(30 * time.Minute)
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
		// nmap.WithPorts("80,443,843"),
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
