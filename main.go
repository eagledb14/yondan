package main

import (
	"fmt"
	"sort"
	"strings"
	"net"
	"net/url"

	"github.com/eagledb14/shodan-clone/template"
	"github.com/eagledb14/shodan-clone/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := utils.NewConcurrentMap()
	Populate(db)

	serv(":8080", db)
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
		params = strings.ToLower(params)

		scans, _ := utils.Query(params, db)

		if len(scans) == 0 {
			c.Redirect("/missing/" + url.QueryEscape(params))
		} else if len(scans) == 1 {
			c.Redirect("/host/" + scans[0].Ip)
		}

		sort.Slice(scans, func(i, j int) bool {
			ip1 := net.ParseIP(scans[i].Ip).To4()
			ip2 := net.ParseIP(scans[j].Ip).To4()
		
			for k := 0; k < 4; k++ {
				if ip1[k] != ip2[k] {
					return ip1[k] < ip2[k]
				}
			}
				return false
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

	app.Get("/missing/:params", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		params := c.Params("params")
		params, _ = url.QueryUnescape(params)
		return c.SendString(template.BuildPage(template.Missing(), params))
	})

	app.Get("/missing", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString(template.BuildPage(template.Missing(), ""))
	})

	app.Static("/favicon.png", "./resources/favicon.png")
	app.Static("/htmx", "./resources/htmx.js")
	app.Static("/styles.css", "./resources/styles.css")
	app.Static("/logo", "./resources/logo.png")

	err := app.Listen(port)
	fmt.Println(err)
}

