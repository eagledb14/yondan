package main

import (
	"fmt"
	"github.com/Ullaakut/nmap/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/eagledb14/shodan-clone/template"
	// "math/rand"
	// "time"
)

func main() {
	// db := NewConcurrentMap[Scan]()
	
	// go Poll(db)
	// for {
	// 	time.Sleep(1 * time.Second)
	// 	scan, err := db.Read("127.0.0.1")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	fmt.Println(scan.Timestamp)
	// }
	// result, _ := poll("127.0.0.1")
	// printHosts(result)

	// ips := []string{"127.0.0.0/28", "127.0.0.16/28", "127.0.0.32/28", "127.0.0.48/28", "127.0.0.64/28", "127.0.0.80/28", "127.0.0.96/28", "127.0.0.112/28", "127.0.0.128/28", "127.0.0.144/28", "127.0.0.160/28", "127.0.0.176/28", "127.0.0.192/28", "127.0.0.208/28", "127.0.0.224/28", "127.0.0.240/28"}
	// for _, i := range ips {
	// 	result, _ := poll(i)
	// 	printHosts(result)
	// 	fmt.Println("Sleeping for 5 seconds...")
	// 	time.Sleep(5 * time.Second)
	// }
	serv(":3000")
}

func serv(port string) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString(template.BuildPage(template.Index("192.168.8.3")))
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("gottem")
	})


	app.Static("/favicon.ico", "./resources/favicon.ico")
	app.Static("/htmx", "./resources/htmx.js")
	app.Static("/styles.css", "./resources/styles.css")
	app.Static("/logo", "./resources/logo.png")

	app.Listen(port)
}

func printHosts(result nmap.Run) {
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}

// func test() {
// 	m := NewConcurrentMap[int]()
// 	m.Write("0", 1)
// 	for i := range 10 {
// 		go func(i int, m *ConcurrentMap[int]) {
// 			for {
// 				v, _ := m.Read("0")
// 				fmt.Println(i, v)
// 			time.Sleep(1 * time.Second)
// 			}
// 		}(i, m)
// 	}
//
// 	for i := range 50 {
// 		fmt.Println("________")
// 		var s int = rand.Intn(10)
// 		m.Write("0", i)
// 		time.Sleep(time.Duration(s) * time.Second)
// 	}
// }
