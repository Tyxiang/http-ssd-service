package main

import (
	// "errors"
	"fmt"
	"http-object/internal/config"
	"http-object/internal/data"
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
	data.Configs, err = config.Load()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
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

	// router.POST("/configs/*uri", config.Post)
	router.GET("/configs/*uri", config.Get)
	router.PUT("/configs/*uri", config.Put)
	// router.DELETE("/configs/*uri", config.Delete)

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
	router.Run(data.Configs.Service.Host + ":" + data.Configs.Service.Port)
}
