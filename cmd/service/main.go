package main

import (
	"errors"
	"http-ssd-service/pkg/config"
	"http-ssd-service/pkg/handler"
	"http-ssd-service/pkg/log"
	"http-ssd-service/pkg/ssd"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// load config
	var earlyErr error // for early err before log init
	err := config.Load("last")
	if err != nil {
		earlyErr = err
		err = config.Load("default")
	}
	if err != nil {
		// default config file lost
		panic(err)
	}
	// init log
	log.Dir = config.Pick("log.dir").String()
	err = log.Init(config.Pick("log.level").String())
	if err != nil {
		panic(err)
	}
	// save early log
	log.Info("service start ... ")
	if earlyErr != nil {
		log.Warn(earlyErr) //trace debug info warn error fatal panic
		log.Info(errors.New("load default config"))
	}
	// load ssd
	ssd.Dir = config.Pick("ssd.dir").String()
	err = ssd.Load("last")
	if err != nil {
		// last ssd file lost
		log.Warn(err)
		log.Info(errors.New("load default ssd"))
		err = ssd.Load("default")
	}
	if err != nil {
		// default ssd file lost
		log.Panic(err)
	}

	// sys
	sys := fiber.New()
	sys.Use(logger.New())
	if config.Pick("service.system.cors").Bool() == true {
		log.Info("system cors on")
		sys.Use(cors.New())
	}
	// sys-router
	sys.Get("/", func(c *fiber.Ctx) error {
		data := []string{"config", "logs", "ssds", "scripts"}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    data,
		})
	})
	//// configs
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
	// sys-listen
	go func() {
		log.Fatal(sys.Listen(config.Pick("service.system.host").String() + ":" + config.Pick("service.system.port").String()))
	}()

	// ssd
	ssd := fiber.New()
	ssd.Use(logger.New())
	if config.Pick("service.system.cors").Bool() == true {
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
