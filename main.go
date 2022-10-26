package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var AngkaWater int
var AngkaWind int

func main() {
	rand.Seed(time.Now().UnixNano())
	max := 20
	go func() {
		for {
			AngkaWater = rand.Intn(max)
			time.Sleep(300 * time.Millisecond)
		}
	}()
	go func() {
		for {
			AngkaWind = rand.Intn(max)
			time.Sleep(300 * time.Millisecond)
		}
	}()
	engine := html.New("./public", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		arr := DataStatus()
		return c.Render("index", arr)
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(DataStatus())
	})

	log.Fatal(app.Listen(":8080"))
}

type Status struct {
	Water       string `json:"water"`
	Wind        string `json:"wind"`
	StatusWater string `json:"status_water"`
	StatusWind  string `json:"status_wind"`
}

func DataStatus() Status {
	var Status Status
	if AngkaWater <= 5 {
		Status.Water = fmt.Sprintf("%dM", AngkaWater)
		Status.StatusWater = "status aman"
	} else if AngkaWater > 5 && AngkaWater <= 8 {
		Status.Water = fmt.Sprintf("%dM", AngkaWater)
		Status.StatusWater = "status siaga"
	} else {
		Status.Water = fmt.Sprintf("%dM", AngkaWater)
		Status.StatusWater = "status bahaya"
	}
	if AngkaWind <= 6 {
		Status.Wind = fmt.Sprintf("%dm/s", AngkaWind)
		Status.StatusWind = "status aman"
	} else if AngkaWind > 6 && AngkaWind <= 15 {
		Status.Wind = fmt.Sprintf("%dm/s", AngkaWind)
		Status.StatusWind = "status siaga"
	} else {
		Status.Wind = fmt.Sprintf("%dm/s", AngkaWind)
		Status.StatusWind = "status bahaya"
	}
	return Status
}
