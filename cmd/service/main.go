package main

import (
	"errors"
	"fmt"
	"http-ssd-service/pkg/config"
	"http-ssd-service/pkg/handler"
	"http-ssd-service/pkg/log"
	"http-ssd-service/pkg/ssd"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// sys ////////////////////////////////////////////////////////////////////
	// init config
	config.Dir = "configs/"
	err := config.Init()
	if err != nil {
		// panic(err)
		fmt.Println("failed to init config, program exit")
		os.Exit(1)
	}
	// init log
	log.Dir = config.Pick("log.dir").String()
	log.Level = config.Pick("log.level").String()
	err = log.Init()
	if err != nil {
		// panic(err)
		fmt.Println("failed to init log, program exit")
		os.Exit(1)
	}
	// save config warn log
	log.Info("service start ... ")
	if config.Warn != nil {
		log.Warn(config.Warn)
		log.Info(errors.New("load default config"))
	}
	// sys-service
	sys := fiber.New(fiber.Config{
		ServerHeader: "http-ssd-service",
		AppName:      "HTTP SSD Service (System)",
	})
	sys.Use(logger.New()) // http log to console
	if config.Pick("service.system.cors").Bool() {
		log.Info("system cors on")
		sys.Use(cors.New())
	}
	// sys-router
	//// root
	sys.Get("/", func(c *fiber.Ctx) error {
		data := []string{"config", "logs", "ssds", "scripts"}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    data,
		})
	})
	//// config
	sys.Post("/config/*", handler.PostConfig)
	sys.Get("/config/*", handler.GetConfig)
	sys.Put("/config/*", handler.PutConfig)
	sys.Delete("/config/*", handler.DeleteConfig)
	//// logs
	sys.Get("/logs", handler.GetLogs)
	sys.Get("/logs/:name", handler.GetLog)
	sys.Delete("/logs/:name", handler.DeleteLog)
	//// ssds
	sys.Post("/ssds", handler.PostSsds)
	sys.Get("/ssds", handler.GetSsds)
	sys.Get("/ssds/:name", handler.GetSsd)
	sys.Put("/ssds/:name", handler.PutSsd)
	sys.Delete("/ssds/:name", handler.DeleteSsd)
	//// scripts
	// ...
	// sys-listen
	go func() {
		log.Fatal(sys.Listen(config.Pick("service.system.host").String() + ":" + config.Pick("service.system.port").String()))
	}()

	// ssd ////////////////////////////////////////////////////////////////////
	// for console display correctly
	time.Sleep(15 * time.Millisecond)
	// init ssd
	ssd.Dir = config.Pick("ssd.dir").String()
	err = ssd.Init()
	if err != nil {
		log.Panic(err)
	}
	// save ssd warn log
	if ssd.Warn != nil {
		log.Warn(ssd.Warn)
		log.Info(errors.New("load default ssd"))
	}
	// ssd-service
	ssd := fiber.New(fiber.Config{
		ServerHeader: "http-ssd-service",
		AppName:      "HTTP SSD Service (SSD)",
	})
	ssd.Use(logger.New())
	if config.Pick("service.system.cors").Bool() {
		log.Info("ssd cors on")
		ssd.Use(cors.New())
	}
	// ssd-router
	ssd.Post("/*", handler.PostSsd)
	ssd.Get("/*", handler.GetSsd)
	ssd.Put("/*", handler.PutSsd)
	ssd.Delete("/*", handler.DeleteSsd)
	// ssd-listen
	log.Fatal(ssd.Listen(config.Pick("service.ssd.host").String() + ":" + config.Pick("service.ssd.port").String()))
}
