package handler

import (
	"http-ssd-service/pkg/config"
	"http-ssd-service/pkg/log"

	"github.com/gofiber/fiber/v2"
)

func PostConfigs(c *fiber.Ctx) error {
	uri := c.Params("*")
	path := uri_to_path(uri)
	data := c.Body()
	err := config.Add(path, data)
	if err != nil {
		log.Save("trace", err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	err = config.Save()
	if err != nil {
		log.Save("error", err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
	})
}
func GetConfigs(c *fiber.Ctx) error {
	uri := c.Params("*")
	path := uri_to_path(uri)
	data, err := config.Get(path)
	if err != nil {
		log.Save("trace", err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    data,
	})
}
func PutConfigs(c *fiber.Ctx) error {
	uri := c.Params("*")
	path := uri_to_path(uri)
	data := c.Body()
	err := config.Set(path, data)
	if err != nil {
		log.Save("trace", err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	err = config.Save()
	if err != nil {
		log.Save("error", err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
	})
}
func DeleteConfigs(c *fiber.Ctx) error {
	uri := c.Params("*")
	path := uri_to_path(uri)
	err := config.Del(path)
	if err != nil {
		log.Save("trace", err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	err = config.Save()
	if err != nil {
		log.Save("error", err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
	})
}
