package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Ullaakut/nmap/v3"
	"net"

	"github.com/eagledb14/shodan-clone/template"
	"github.com/eagledb14/shodan-clone/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
	// "math/rand"
	// "time"
)

func main() {
	db := utils.NewConcurrentMap()
	createTestData("127.0.0.0/20", db)
	// testPoll(db)
	
	// go Poll(db)
	serv(":3000", db)
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

	app.Static("/favicon.ico", "./resources/favicon.ico")
	app.Static("/htmx", "./resources/htmx.js")
	app.Static("/styles.css", "./resources/styles.css")
	app.Static("/logo", "./resources/logo.png")

	app.Listen(port)
}

func testPoll(db *utils.ConcurrentMap) {
	// Poll([]string{"google.com", "facebook.com", "netflix.com"}, db)
	go func() {
		utils.Poll([]string{"127.0.0.1"}, db, 0)
	}()
	time.Sleep(10 * time.Second)
	d := db.ReadAll()

	for k, v := range d {
		fmt.Println("'" + k + "'")
			for _, i := range v {
				fmt.Println("\t" + i.Ip)
			}
	}
}

func createTestData(cidr string, db *utils.ConcurrentMap) {
	ip, ipNet, _ := net.ParseCIDR(cidr)

	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		portNum := rand.Intn(10)

		ports := []nmap.Port{}

		for range portNum {
			num := rand.Intn(65655)
			ports = append(ports, nmap.Port{ID: uint16(num), State: nmap.State{State: "open"}, Service: nmap.Service{Name: "http"}})
		}
		
		newScan := utils.Scan{Ip: ip.String(), Ports: ports, Hostname: "example.com", Timestamp: time.Now().Format("2006-01-02")}

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
