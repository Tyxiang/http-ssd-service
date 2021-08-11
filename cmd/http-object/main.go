package main

import (
	// "errors"
	"fmt"
	"http-object/internal/config"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//init log
	logFilePath := "logs/current.log"
	logFile, err := os.Create(logFilePath)
	if err != nil {
		fmt.Println(err)
	}
	//gin.DefaultWriter = io.MultiWriter(logFile) // writer file only, do not write console
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout) // writer file and console
	fmt.Fprintln(gin.DefaultWriter, "service start ... ")
	//load config
	err = config.LoadCurrent()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		fmt.Fprintln(gin.DefaultWriter, "load default config")
		err = config.LoadDefault()
	}
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		panic(err)
	}
	// fmt.Printf("%+v", cfg) //打印结构体
	//make router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(500, gin.H{
			"success": true,
			"name":    "http object service",
		})
		return
	})

	// router.GET("/configs/*uri", config.Get)
	// router.PUT("/configs/*uri", config.Put)

	// configs := router.Group("/configs")
	// {
	// 	configs.POST("/*uri", post_configs)
	// 	configs.GET("/*uri", get_configs)
	// 	configs.PUT("/*uri", put_configs)
	// 	configs.DELETE("/*uri", delete_configs)
	// }
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
	router.Run(config.Values.Service.Host + ":" + config.Values.Service.Port)
}

// func get_configs(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"success": true,
// 		"data":    config.Configs,
// 	})
// }
// func put_configs(c *gin.Context) {

// }
