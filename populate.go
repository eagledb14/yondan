package main

import (
	"github.com/eagledb14/shodan-clone/utils"
	"github.com/Ullaakut/nmap/v3"
	"fmt"
	"os"
	"bufio"
	"math/rand"
	crand "crypto/rand"
	"net"
	"strconv"
	"time"
)

func Populate(db *utils.ConcurrentMap) {
	dummyRanges := []string{}
	for range 200 {
		dummyRanges = append(dummyRanges, getRandomCidr())
	}

	populateExamples(db)

	ranges := readRanges()

	go func() {
		for _, cidr := range dummyRanges {
			createDummyData(cidr, db)
		}

		utils.Poll(ranges, db, 0)
		fmt.Println("Full Scan Complete: db size", db.Len())

		for {
			for _, cidr := range dummyRanges {
				createDummyData(cidr, db)
			}

			utils.Poll(ranges, db, 10)
			fmt.Println("Full Scan Complete: db size", db.Len())
		}
	}()
}

func readRanges() []string {
	file, err := os.Open("./resources/ranges.txt")

	if err != nil {
		panic("Missing file \"resources/ranges.txt\"")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ranges := []string{}

	for scanner.Scan() {
		ranges = append(ranges, scanner.Text())
	}

	return ranges
}

func getRandomCidr() string {
	ip := make([]byte, 4)
	crand.Read(ip)
	subnet := rand.Intn(6) + 22

	ipStr := fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])

	cidr := fmt.Sprintf("%s/%d", ipStr, subnet)

	return cidr
}

func createDummyData(cidr string, db *utils.ConcurrentMap, url ...string) {
	ip, ipNet, _ := net.ParseCIDR(cidr)

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
	
		newScan := utils.Scan{Ip: ip.String(), Hostname: hostname,  Ports: ports, Timestamp: time.Now().Format("2006-01-02")}

		db.Write(ip.String(), newScan)
		db.Write(newScan.Hostname, newScan)
		for _, port := range newScan.Ports {
			db.Write(strconv.Itoa(int(port.ID)), newScan)
			db.Write(port.Service.Name, newScan)
		}
	}
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

func populateExamples(db *utils.ConcurrentMap) {
	createDummyData("8.8.8.8/24", db)
	createDummyData(getRandomCidr(), db, "example.com")
}
