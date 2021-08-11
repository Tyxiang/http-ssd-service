package main

import (
	// "errors"
	"fmt"
	// "io"
	// "os"
	// "strings"
	// "time"
	// "github.com/gin-gonic/gin"
	// "/internal/log"
	// "/internal/config"
	"http-object/internal/authn"
)

func main() {
	fmt.Printf("main")
	authn.Hello()
	/*
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
		authns := router.Group("/authns")
		{
			authns.POST("/*uri", post_authns)
			authns.GET("/*uri", get_authns)
			authns.PUT("/*uri", put_authns)
			authns.DELETE("/*uri", delete_authns)
		}
		//run router
		router.Run(config.GetString("system.host") + ":" + config.GetString("system.port"))
	*/
}
