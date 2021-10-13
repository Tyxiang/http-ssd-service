package main

import (
	"fmt"
	"http-ssd-service/pkg/config"
	"http-ssd-service/pkg/log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	// logrus.Trace("trace msg")
	// logrus.Debug("debug msg")
	// logrus.Info("info msg")
	// logrus.Warn("warn msg")
	// logrus.Error("error msg")
	// logrus.Fatal("fatal msg")
	// logrus.Panic("panic msg")
	logrus.Info("service start ... ")
	////load config
	err := config.Load("last")
	if err != nil {
		logrus.Warn(err)
		err = config.Load("default")
	}
	if err != nil {
		logrus.Panic(err)
	}
	////sys
	sys := fiber.New()
	sys.Use(logger.New())
	sysCors, _ := config.Get("system.cors")
	if sysCors == true {
		logrus.Info("system cors on")
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
	sys.Post("/configs/*", func(c *fiber.Ctx) error {
		uri := c.Params("*")
		path := uri_to_path(uri)
		data := c.Body()
		err = config.Add(path, data)
		if err != nil {
			logrus.Trace(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		err = config.Save()
		if err != nil {
			logrus.Error(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})
	sys.Get("/configs/*", func(c *fiber.Ctx) error {
		uri := c.Params("*")
		path := uri_to_path(uri)
		data, err := config.Get(path)
		if err != nil {
			logrus.Trace(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    data,
		})
	})
	sys.Put("/configs/*", func(c *fiber.Ctx) error {
		uri := c.Params("*")
		path := uri_to_path(uri)
		data := c.Body()

		fmt.Println(data)

		err = config.Set(path, data)
		if err != nil {
			logrus.Trace(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		err = config.Save()
		if err != nil {
			logrus.Error(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})
	sys.Delete("/configs/*", func(c *fiber.Ctx) error {
		uri := c.Params("*")
		path := uri_to_path(uri)
		err = config.Del(path)
		if err != nil {
			logrus.Trace(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		err = config.Save()
		if err != nil {
			logrus.Error(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})
	//////logs
	sys.Get("/logs", func(c *fiber.Ctx) error {
		list, err := log.List()
		if err != nil {
			logrus.Trace(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    list,
		})
	})
	sys.Get("/logs/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		data, err := log.Read(name)
		if err != nil {
			logrus.Trace(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		return c.Send(data)
	})
	sys.Delete("/logs/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		err := log.Remove(name)
		if err != nil {
			logrus.Trace(err)
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})

	////sys-listen
	sysHost, _ := config.Get("system.host")
	sysPort, _ := config.Get("system.port")
	go func() {
		logrus.Fatal(sys.Listen(sysHost.(string) + ":" + sysPort.(string)))
	}()

	////ssd
	ssd := fiber.New()
	ssd.Use(logger.New())
	ssdCors, _ := config.Get("system.cors")
	if ssdCors == true {
		logrus.Info("ssd cors on")
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
	logrus.Fatal(ssd.Listen(ssdHost.(string) + ":" + ssdPort.(string)))
}

func uri_to_path(uri string) string {
	path_1 := strings.Trim(uri, "/")
	path_2 := strings.Replace(path_1, "/", ".", -1)
	path_3 := strings.Replace(path_2, "()", ".-1", -1)
	path_4 := strings.Replace(path_3, "(", ".", -1)
	path := strings.Replace(path_4, ")", "", -1)
	//fmt.Println(path)
	return path
}
