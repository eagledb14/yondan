package main

import (
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/Ullaakut/nmap/v3"
	"github.com/eagledb14/shodan-clone/utils"
)

type FlagData struct {
	Ip string `json:"ip"`
	Ports []Port `json:"ports"`
	Hostname string `json:"hostname"`
}

type Port struct {
	Id int `json:"id"`
	Service string `json:"service"`
}

func readFlagData(db *utils.ConcurrentMap) []*utils.Scan {
	file, err := os.ReadFile("./resources/flags.txt")

	if err != nil {
		return []*utils.Scan{}
	}

	flagData := []FlagData{}

	err = json.Unmarshal(file, &flagData)
	if err != nil {
		panic("Flags.txt json error" + err.Error())
	}

	flagScans := []*utils.Scan{}
	for _, flag := range flagData {
		ports := []nmap.Port{}
		for _, port := range flag.Ports {
			ports = append(ports, nmap.Port{ID: uint16(port.Id), State: nmap.State{State: "open"}, Service: nmap.Service{Name: port.Service}})
		}

		flagScans = append(flagScans, &utils.Scan{Ip: flag.Ip, Hostname: flag.Hostname, Ports: ports, Timestamp: time.Now().Format("2006-01-02")})
	}

	//adds flags to the database
	for _, flag := range flagScans {
		exists := db.DummyWrite(flag.Ip, flag)
		if exists {
			continue
		}
		db.Write(flag.Hostname, flag)
		for _, port := range flag.Ports {
			db.Write(strconv.Itoa(int(port.ID)), flag)
			db.Write(port.Service.Name, flag)
		}
	}

	return flagScans
}

func createDummyData(cidr string, db *utils.ConcurrentMap, url ...string) []*utils.Scan {
	ip, ipNet, _ := net.ParseCIDR(cidr)
	dummyData := []*utils.Scan{}

	index := 0

	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		index++
		portNum := rand.Intn(4) + 1

		ports := []nmap.Port{}

		for range portNum {
			ports = append(ports, getRandomPort())
		}

		hostname :=	""
		if len(url) > 0 {
			hostname = url[0]
		}
	
		newScan := &utils.Scan{Ip: ip.String(), Hostname: hostname,  Ports: ports, Timestamp: time.Now().Format("2006-01-02")}
		dummyData = append(dummyData, newScan)

		exists := db.DummyWrite(ip.String(), newScan)
		if exists {
			continue
		}
		db.Write(newScan.Hostname, newScan)
		for _, port := range newScan.Ports {
			db.Write(strconv.Itoa(int(port.ID)), newScan)
			db.Write(port.Service.Name, newScan)
		}
	}

	return dummyData
}

func getRandomCidr() string {
	ip := make([]byte, 4)
	crand.Read(ip)
	subnet := rand.Intn(6) + 22

	ipStr := fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])

	cidr := fmt.Sprintf("%s/%d", ipStr, subnet)

	return cidr
}

func incrementIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}

func getRandomPort() nmap.Port {
	names := []string{"http", "https", "telnet","ftp","ssh","smtp","rdp","pop3","microsoft-ds","netbios-ssn","imap","domain","msrpc","mysql","http-proxy","pptp","rpcbind","pop4s","imaps","vnc","nfs","submission","smux","smpts","http","unkown","printer","dc","nfs","ftps","BGP", "time","ntpd"}
	ports := []uint16{80,443,23,21,22,25,3389,110,445,139,143,53,135,3306,8080,1723,111,995,993,5900,1025,587,199,465,8008,49152,515,2001,2049,990, 179,37,123}

	randIndex := rand.Intn(len(names))

	return nmap.Port{ID: ports[randIndex], State: nmap.State{State: "open"}, Service: nmap.Service{Name: names[randIndex]}}
}

func updateDummyRangeTime(dummyScans []*utils.Scan) {
	for _, s := range dummyScans {
		s.Timestamp = time.Now().Format("2006-01-02")
	}
}
