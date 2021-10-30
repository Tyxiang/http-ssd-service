package handler

import (
	"http-ssd-service/pkg/config"
	"http-ssd-service/pkg/log"

	"github.com/gofiber/fiber/v2"
)

func PostConfig(c *fiber.Ctx) error {
	u := c.Params("*")
	path, _ := parse(u)
	data := c.Body()
	err := config.Add(path, data)
	if err != nil {
		log.Trace(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	err = config.Save()
	if err != nil {
		log.Error(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
	})
}
func GetConfig(c *fiber.Ctx) error {
	u := c.Params("*")
	path, _ := parse(u)
	data := config.Get(path)
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    data,
	})
}
func PutConfig(c *fiber.Ctx) error {
	u := c.Params("*")
	path, _ := parse(u)
	data := c.Body()
	err := config.Set(path, data)
	if err != nil {
		log.Trace(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	err = config.Save()
	if err != nil {
		log.Error(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
	})
}
func DeleteConfig(c *fiber.Ctx) error {
	u := c.Params("*")
	path, _ := parse(u)
	err := config.Del(path)
	if err != nil {
		log.Trace(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	err = config.Save()
	if err != nil {
		log.Error(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
	})
}
