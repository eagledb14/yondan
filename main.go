package main

import (
	"fmt"
	"time"

	"github.com/eagledb14/shodan-clone/template"
	"github.com/gofiber/fiber/v2"
	// "math/rand"
	// "time"
)

func main() {

	db := NewConcurrentMap()
	// testPoll(db)
	db.Write("127.0.0.1", Scan{"127.0.0.1", nil, "example.com", time.Now()})
	db.Write("127.0.0.37", Scan{"127.0.0.1", nil, "example.com", time.Now()})
	db.Write("127.0.0.102", Scan{"127.0.0.1", nil, "example.com", time.Now()})
	db.Write("127.0.0.230", Scan{"127.0.0.1", nil, "example.com", time.Now()})
	db.Write("example.com", Scan{"127.0.0.1", nil, "example.com", time.Now()})
	db.Write("80", Scan{"127.0.0.1", nil, "example.com", time.Now()})

	fmt.Println(query("domain:example.com net:127.0.0.1", db))
	// fmt.Println(query("domain:example", db))

	// query("domain:example.com port:22 ip:8.8.8.8/24")
	// fmt.Println(db.Read("127.0.0.1"))
	// fmt.Println(query("net:127.0.0.0/0", db))
	// fmt.Println(query("port:80", db))
	// fmt.Println(query("domain:monkey.com", db))
	// fmt.Println(ParseCidr("142.250.9.0/24", db))
	
	// tempMap := make(map[string][]Scan)
	// addDomainToMap(&tempMap, Scan{"127.0.0.1", []nmap.Port{}, "example.com", time.Now()})
	// fmt.Println(tempMap)



	
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
	// serv(":3000", db)
}

func serv(port string, db *ConcurrentMap) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString(template.BuildPage(template.Index(),"192.168.8.3"))
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		return c.SendString("gottem")
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		params := c.Query("query")

		scans, _ := query(params, db)
		_ = scans

		if len(scans) == 0 {
			//return missing page, or not avaialable page
		} else if len(scans) == 1 {
			//return host page
		} else {
			//return search page
		}

		c.Set("Content-Type", "text/html")
		return c.SendString(template.BuildPage("", params))
	})

	app.Static("/favicon.ico", "./resources/favicon.ico")
	app.Static("/htmx", "./resources/htmx.js")
	app.Static("/styles.css", "./resources/styles.css")
	app.Static("/logo", "./resources/logo.png")

	app.Listen(port)
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
//
func testPoll(db *ConcurrentMap) {
	// Poll([]string{"google.com", "facebook.com", "netflix.com"}, db)
	go func() {
		Poll([]string{"127.0.0.1"}, db, 0)
	}()
	time.Sleep(20 * time.Second)
	d := db.ReadAll()

	for k, v := range d {
		fmt.Println("'" + k + "'")
			for _, i := range v {
				fmt.Println("\t" + i.Ip)
			}
	}
}
