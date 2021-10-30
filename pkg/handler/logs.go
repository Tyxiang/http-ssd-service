package handler

import (
	"http-ssd-service/pkg/log"

	"github.com/gofiber/fiber/v2"
)

func GetLogs(c *fiber.Ctx) error {
	list, err := log.List()
	if err != nil {
		log.Trace(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    list,
	})
}

func GetLog(c *fiber.Ctx) error {
	name := c.Params("name")
	data, err := log.Read(name)
	if err != nil {
		log.Trace(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Send(data)
}

func DeleteLog(c *fiber.Ctx) error {
	name := c.Params("name")
	err := log.Remove(name)
	if err != nil {
		log.Trace(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
	})
}
