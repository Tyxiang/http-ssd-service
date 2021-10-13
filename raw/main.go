package main

import (
	// "errors"

	"fmt"
	"http-ssd-service/pkg/config"
	"io"
	"os"
	"strings"

	//"github.com/tidwall/sjson"
	//"github.com/tidwall/gjson"
	"github.com/gin-gonic/gin"
)

var logDir = "logs/"

func main() {
	//init log
	logFilePath := logDir + "last.log"
	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	//gin.DefaultWriter = io.MultiWriter(logFile) // writer file only, do not write console
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout) // writer file and console
	fmt.Fprintln(gin.DefaultWriter, "service start ... ")
	//load config
	err = config.Load("last")
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		err = config.Load("default")
	}
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		panic(err)
	}
	//make router
	router := gin.Default()
	system := router.Group("/system")
	system.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"data": ["configs", "logs", "persistences", "scripts"],
		  })
	})
	configs := system.Group("/configs")
	{
		// configs.GET("", func(c *gin.Context) {
		// 	value, err := config.List()
		// 	if err != nil {
		// 		fmt.Fprintln(gin.DefaultWriter, err.Error())
		// 		c.JSON(500, gin.H{
		// 			"success": false,
		// 			"message": err.Error(),
		// 		})
		// 	}
		// 	c.JSON(200, gin.H{
		// 		"success": true,
		// 		"data":    value,
		// 	})
		// })
		configs.GET("/*uri", func(c *gin.Context) {
			uri := c.Param("uri")
			//fmt.Println(uri)
			//q := c.Query("q")
			//s := c.Query("s")
			path := uri_to_path(uri)
			//fmt.Println(path)
			value := config.Get(path)
			c.JSON(200, gin.H{
				"success": true,
				"data":    value,
			})
		})
		configs.PUT("/:name", func(c *gin.Context) {
			name := c.Param("name")
			data, err := c.GetRawData()
			if err != nil {
				fmt.Fprintln(gin.DefaultWriter, err.Error())
				c.JSON(500, gin.H{
					"success": false,
					"message": err.Error(),
				})
			}
			err = config.New(name, data)
			if err != nil {
				fmt.Fprintln(gin.DefaultWriter, err.Error())
				c.JSON(500, gin.H{
					"success": false,
					"message": err.Error(),
				})
			}
			c.JSON(200, gin.H{
				"success": true,
			})
		})
		configs.DELETE("/:name", func(c *gin.Context) {
			name := c.Param("name")
			err = config.Remove(name)
			if err != nil {
				fmt.Fprintln(gin.DefaultWriter, err.Error())
				c.JSON(500, gin.H{
					"success": false,
					"message": err.Error(),
				})
			}
			c.JSON(200, gin.H{
				"success": true,
			})
		})
	}

	//run router
	host := config.Get("host").(string)
	port := config.Get("port").(string)
	router.Run(host + ":" + port)
}

func uri_to_path(uri string) string {
	path_1 := strings.Trim(uri, "/")
	path_2 := strings.Replace(path_1, "/", ".", -1)
	path_3 := strings.Replace(path_2, "(", ".", -1)
	path := strings.Replace(path_3, ")", "", -1)
	//fmt.Println(path)
	return path
}
