package main

import (
	"bufio"
	"fmt"
	"sort"

	"math/rand"
	crand "crypto/rand"
	"os"
	"time"

	"net"

	"github.com/Ullaakut/nmap/v3"

	"strconv"

	"github.com/eagledb14/shodan-clone/template"
	"github.com/eagledb14/shodan-clone/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := utils.NewConcurrentMap()

	dummyRanges := []string{}
	for range 200 {
		dummyRanges = append(dummyRanges, getRandomCidr())
	}

	ranges := readRanges()
	for _, cidr := range dummyRanges {
		createDummyData(cidr, db)
	}

	go func() {
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

	serv(":3000", db)
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

func serv(port string, db *utils.ConcurrentMap) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString(template.BuildPage(template.Index(), ""))
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		params := c.Query("query")

		scans, _ := utils.Query(params, db)

		if len(scans) == 0 {
			return c.SendString(template.BuildPage(template.Missing(), params))
		} else if len(scans) == 1 {
			c.Redirect("/host/" + scans[0].Ip)
		}

		sort.Slice(scans, func(i, j int) bool {
			return scans[i].Ip < scans[j].Ip
		})

		return c.SendString(template.BuildPage(template.Search(scans, params), params))
	})

	app.Get("/host/:ip", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		ip := c.Params("ip")

		scan, err := db.Read(ip)
		if err != nil || len(scan) != 1 {
			return c.SendString(template.BuildPage(template.Missing(), ip))
		}

		return c.SendString(template.BuildPage(template.Host(scan[0], db), ip))
	})

	app.Static("/favicon.png", "./resources/favicon.png")
	app.Static("/htmx", "./resources/htmx.js")
	app.Static("/styles.css", "./resources/styles.css")
	app.Static("/logo", "./resources/logo.png")

	app.Listen(port)
}

func getRandomCidr() string {
	ip := make([]byte, 4)
	crand.Read(ip)
	subnet := rand.Intn(6) + 22

	ipStr := fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])

	cidr := fmt.Sprintf("%s/%d", ipStr, subnet)

	return cidr
}

func createDummyData(cidr string, db *utils.ConcurrentMap) {
	ip, ipNet, _ := net.ParseCIDR(cidr)

	index := 0

	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		index++
		portNum := rand.Intn(4) + 1

		ports := []nmap.Port{}

		for range portNum {
			ports = append(ports, getRandomPort())
		}
		
		newScan := utils.Scan{Ip: ip.String(), Ports: ports, Timestamp: time.Now().Format("2006-01-02")}

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
