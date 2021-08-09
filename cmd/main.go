package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/spf13/viper"

	"Configs/"
	"Logs"
)

func main() {
	//log
	log_file_name := time.Now().Format("2006-01-02T15-04-05")
	log_file, _ := os.Create("./logs/" + log_file_name + ".log")
	//gin.DefaultWriter = io.MultiWriter(log_file) // writer file only, do not write console
	gin.DefaultWriter = io.MultiWriter(log_file, os.Stdout) // writer file and console

	//make router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(500, gin.H{
			"success": true,
			"name":    "http object service",
		})
		return
	})
	configs := router.Group("/configs")
	{
		configs.POST("/*uri", post_configs)
		configs.GET("/*uri", get_configs)
		configs.PUT("/*uri", put_configs)
		configs.DELETE("/*uri", delete_configs)
	}
	logs := router.Group("/logs")
	{
		logs.POST("/*uri", post_logs)
		logs.GET("/*uri", get_logs)
		logs.PUT("/*uri", put_logs)
		logs.DELETE("/*uri", delete_logs)
	}
	objects := router.Group("/objects")
	{
		objects.POST("/*uri", post_objects)
		objects.GET("/*uri", get_objects)
		objects.PUT("/*uri", put_objects)
		objects.DELETE("/*uri", delete_objects)
	}
	scripts := router.Group("/scripts")
	{
		scripts.POST("/*uri", post_scripts)
		scripts.GET("/*uri", get_scripts)
		scripts.PUT("/*uri", put_scripts)
		scripts.DELETE("/*uri", delete_scripts)
	}
	secrets := router.Group("/secrets")
	{
		secrets.POST("/*uri", post_secrets)
		secrets.GET("/*uri", get_secrets)
		secrets.PUT("/*uri", put_secrets)
		secrets.DELETE("/*uri", delete_secrets)
	}
	//run router
	router.Run(config.GetString("system.host") + ":" + config.GetString("system.port"))
}