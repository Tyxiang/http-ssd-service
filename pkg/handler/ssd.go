package handler

import (
	"http-ssd-service/pkg/log"
	"http-ssd-service/pkg/ssd"

	"github.com/gofiber/fiber/v2"
)

//file
func PostSsds(c *fiber.Ctx) error {
	err := ssd.Save()
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
func GetSsds(c *fiber.Ctx) error {
	list, err := ssd.List()
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

//data
func PostSsd(c *fiber.Ctx) error {
	u := c.Params("*")
	path, _ := parse(u)
	data := c.Body()
	err := ssd.Add(path, data)
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

func GetSsd(c *fiber.Ctx) error {
	u := c.Params("*")
	path, property := parse(u)
	var data interface{}
	if property == "type" {
		data = ssd.GetType(path)
	}
	if property == "" {
		data = ssd.Get(path)
	}
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    data,
	})
}

func PutSsd(c *fiber.Ctx) error {
	u := c.Params("*")
	path, _ := parse(u)
	data := c.Body()
	err := ssd.Set(path, data)
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

func DeleteSsd(c *fiber.Ctx) error {
	u := c.Params("*")
	path, _ := parse(u)
	err := ssd.Del(path)
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
