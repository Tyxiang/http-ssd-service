package main

import (
	// "errors"

	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	//"http-object/internal/router"
	//"github.com/tidwall/sjson"
	//"github.com/tidwall/gjson"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

var configBytes []byte

func main() {
	//dirs
	configDir := "configs/"
	logDir := "logs/"
	// persistenceDir := "persistences/"
	// script := "script/"
	//init log
	logFilePath := logDir + "last.log"
	logFile, err := os.Create(logFilePath)
	if err != nil {
		fmt.Println(err)
	}
	//gin.DefaultWriter = io.MultiWriter(logFile) // writer file only, do not write console
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout) // writer file and console
	fmt.Fprintln(gin.DefaultWriter, "service start ... ")
	//load config
	fmt.Fprintln(gin.DefaultWriter, "load last config")
	configFilePath := configDir + "last.json"
	configBytes, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		return
	}
	if !gjson.Valid(string(configBytes)) {
		fmt.Fprintln(gin.DefaultWriter, "load default config")
		configFilePath = configDir + "default.json"
		configBytes, err = ioutil.ReadFile(configFilePath)
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			panic(err)
		}
	}
	//fmt.Println(string(configBytes))
	//make router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		name := gjson.GetBytes(configBytes, "name")
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			return
		}
		c.JSON(500, gin.H{
			"success": true,
			"name":    name.String(),
		})
	})
	configs := router.Group("/configs")
	{
		// configs.POST("/*uri", post_configs)
		configs.GET("/*uri", get_configs)
		// configs.PUT("/*uri", router.PUT)
		// configs.DELETE("/*uri", delete_configs)
	}
	// logs := router.Group("/logs")
	// {
	// 	logs.POST("/*uri", post_logs)
	// 	logs.GET("/*uri", get_logs)
	// 	logs.PUT("/*uri", put_logs)
	// 	logs.DELETE("/*uri", delete_logs)
	// }
	// objects := router.Group("/objects")
	// {
	// 	objects.POST("/*uri", post_objects)
	// 	objects.GET("/*uri", get_objects)
	// 	objects.PUT("/*uri", put_objects)
	// 	objects.DELETE("/*uri", delete_objects)
	// }
	// scripts := router.Group("/scripts")
	// {
	// 	scripts.POST("/*uri", post_scripts)
	// 	scripts.GET("/*uri", get_scripts)
	// 	scripts.PUT("/*uri", put_scripts)
	// 	scripts.DELETE("/*uri", delete_scripts)
	// }
	// authns := router.Group("/authns")
	// {
	// 	authns.POST("/*uri", post_authns)
	// 	authns.GET("/*uri", get_authns)
	// 	authns.PUT("/*uri", put_authns)
	// 	authns.DELETE("/*uri", delete_authns)
	// }

	//run router
	host := gjson.GetBytes(configBytes, "host").String()
	port := gjson.GetBytes(configBytes, "port").String()
	router.Run(host + ":" + port)
}

func get_configs(c *gin.Context) {
	uri := c.Param("uri")
	//q := c.Query("q")
	//s := c.Query("s")
	//path := uri_to_path(uri)
	var value interface{}
	if uri == "/" {
		value = gjson.ParseBytes(configBytes).Value()
	} else {
		path := uri_to_path(uri)
		value = gjson.GetBytes(configBytes, path).Value()
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    value,
	})
}

func uri_to_path(uri string) string {
	path_1 := strings.Trim(uri, "/")
	path_2 := strings.Replace(path_1, "/", ".", -1)
	path_3 := strings.Replace(path_2, "(", ".", -1)
	path := strings.Replace(path_3, ")", "", -1)
	//fmt.Println(path)
	return path
}
