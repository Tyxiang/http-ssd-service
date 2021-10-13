package main

import (
	"http-ssd-service/pkg/config"
	"http-ssd-service/pkg/handler"
	"http-ssd-service/pkg/log"
	"http-ssd-service/pkg/persistence"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	err := log.Init()
	if err != nil {
		panic(err)
	}
	log.Save("info", "service start ... ") //trace debug info warn error fatal panic
	////load config
	err = config.Load("last")
	if err != nil {
		log.Save("warn", err)
		err = config.Load("default")
	}
	if err != nil {
		log.Save("panic", err)
	}
	////load persistence
	err = persistence.Load("last")
	if err != nil {
		log.Save("warn", err)
		err = persistence.Load("default")
	}
	if err != nil {
		log.Save("panic", err)
	}
	////sys
	sys := fiber.New()
	sys.Use(logger.New())
	sysCors, _ := config.Get("system.cors")
	if sysCors == true {
		log.Save("info", "system cors on")
		sys.Use(cors.New())
	}
	////sys-router
	sys.Get("/", func(c *fiber.Ctx) error {
		data := []string{"configs", "logs", "persistences", "scripts"}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    data,
		})
	})
	//////configs
	sys.Post("/configs/*", handler.PostConfigs)
	sys.Get("/configs/*", handler.GetConfigs)
	sys.Put("/configs/*", handler.PutConfigs)
	sys.Delete("/configs/*", handler.DeleteConfigs)
	//////logs
	sys.Get("/logs", handler.GetLogs)
	sys.Get("/logs/:name", handler.GetLog)
	sys.Delete("/logs/:name", handler.DeleteLog)
	//////persistences
	sys.Post("/persistences", handler.PostPersistences)
	sys.Get("/persistences", handler.GetPersistences)
	sys.Get("/persistences/:name", handler.GetPersistence)
	sys.Put("/persistences/:name", handler.PutPersistence)
	sys.Delete("/persistences/:name", handler.DeletePersistence)
	//////scripts
	////sys-listen
	sysHost, _ := config.Get("system.host")
	sysPort, _ := config.Get("system.port")
	go func() {
		log.Save("fatal", sys.Listen(sysHost.(string)+":"+sysPort.(string)))
	}()

	////ssd
	ssd := fiber.New()
	ssd.Use(logger.New())
	ssdCors, _ := config.Get("system.cors")
	if ssdCors == true {
		log.Save("info", "ssd cors on")
		ssd.Use(cors.New())
	}
	////ssd-router
	ssd.Get("/*", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    "ssd",
		})
	})
	////ssd-listen
	ssdHost, _ := config.Get("ssd.host")
	ssdPort, _ := config.Get("ssd.port")
	log.Save("fatal", ssd.Listen(ssdHost.(string)+":"+ssdPort.(string)))
}
