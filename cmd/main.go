package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var object string = "object"

func main() {
	//log
	log_file_name := time.Now().Format("2006-01-02T15-04-05")
	log_file, _ := os.Create("./logs/" + log_file_name + ".log")
	//gin.DefaultWriter = io.MultiWriter(log_file) // writer file only, do not write console
	gin.DefaultWriter = io.MultiWriter(log_file, os.Stdout) // writer file and console
	//config
	config := viper.New()
	config.AddConfigPath("./configs")
	config.SetConfigType("json")
	//load current config file
	config.SetConfigName("current") // name of config file (without extension)
	err := config.ReadInConfig()
	if err != nil {
		if _, no := err.(viper.ConfigFileNotFoundError); no {
			//config file not found
			//load default config
			config.SetConfigName("default")
			err := config.ReadInConfig()
			if err != nil {
				fmt.Fprintln(gin.DefaultWriter, err.Error())
				panic(err)
			}
		} else {
			//config file was found, but another error was produced
			//program end
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			panic(err)
		}
	}
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

func post_admin(c *gin.Context) {
	//uri
	uri := c.Param("uri")
	//get data
	data, err := c.GetRawData()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//post to ovl
	if strings.HasSuffix(uri, "()") { //ovl 新增以空括号结尾
		uri = strings.TrimSuffix(uri, "()")
		json_path := uri_to_json_path(uri)
		r, err := redis_clinet.Do("json.arrappend", key_object, json_path, string(data)).Result()
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			//"index": strconv.FormatInt(r.(int64)-1, 10),
			"index": r.(int64) - 1,
		})
		return
	}
	//post to else
	if true {
		json_path := uri_to_json_path(uri)
		_, err = redis_clinet.Do("json.set", key_object, json_path, string(data), "NX").Result() // 只在不存在时设置。也就是POST前必须删除，已存在的话只能用PUT
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
		})
		return
	}
}
func post_object(c *gin.Context) {
	//uri
	uri := c.Param("uri")
	//get data
	data, err := c.GetRawData()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//can only post to ovl
	if !strings.HasSuffix(uri, "()") { //ovl 新增必须以空括号结尾
		err = errors.New("can not do this")
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	}
	//post to ovl
	uri = strings.TrimSuffix(uri, "()")
	json_path := uri_to_json_path(uri)
	r, err := redis_clinet.Do("json.arrappend", key_object, json_path, string(data)).Result()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//default response success
	c.JSON(200, gin.H{
		"success": true,
		//"index": strconv.FormatInt(r.(int64)-1, 10),
		"index": r.(int64) - 1,
	})
	return
}

func get_admin(c *gin.Context) {
	//uri
	uri := c.Param("uri")
	//.type
	if strings.HasSuffix(uri, ".type") {
		uri = strings.TrimSuffix(uri, ".type")
		json_path := uri_to_json_path(uri)
		r, err := redis_clinet.Do("json.type", key_object, json_path).Result()
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    r,
		})
		return
	}
	//.length
	if strings.HasSuffix(uri, ".length") {
		uri = strings.TrimSuffix(uri, ".length")
		json_path := uri_to_json_path(uri)
		r, err := redis_clinet.Do("json.type", key_object, json_path).Result()
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		var cmd string
		switch r {
		case "string":
			cmd = "json.strlen"
		case "array":
			cmd = "json.arrlen"
		case "object":
			cmd = "json.objlen"
		default:
			err = errors.New("no length attribute")
			if err != nil {
				fmt.Fprintln(gin.DefaultWriter, err.Error())
				c.JSON(500, gin.H{
					"success": false,
					"message": err.Error(),
				})
				return
			}
			return
		}
		r, err = redis_clinet.Do(cmd, key_object, json_path).Result()
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    r,
		})
		return
	}
	//.keys
	if strings.HasSuffix(uri, ".keys") {
		uri = strings.TrimSuffix(uri, ".keys")
		json_path := uri_to_json_path(uri)
		r, err := redis_clinet.Do("json.objkeys", key_object, json_path).Result()
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    r,
		})
		return
	}
	//.memory
	if strings.HasSuffix(uri, ".memory") {
		uri = strings.TrimSuffix(uri, ".memory")
		json_path := uri_to_json_path(uri)
		r, err := redis_clinet.Do("json.debug", "memory", key_object, json_path).Result()
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    r,
		})
		return
	}
	//else
	if true {
		json_path := uri_to_json_path(uri)
		r, err := redis_clinet.Do("json.get", key_object, json_path).Result()
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(200, "{\"success\": true, \"data\":"+r.(string)+"}")
		return
	}
}
func get_object(c *gin.Context) {
	//uri
	uri := c.Param("uri")
	//query
	q := c.Query("q")
	s := c.Query("s")
	//get object
	json_path := uri_to_json_path(uri)
	r, err := redis_clinet.Do("json.get", key_object, json_path).Result()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//default response success
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(200, "{\"success\": true, \"data\":"+r.(string)+"}")
	return
}

func put(c *gin.Context) {
	//uri
	uri := c.Param("uri")
	//get data
	json_path := uri_to_json_path(uri)
	data, err := c.GetRawData()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//put
	_, err = redis_clinet.Do("json.set", key_object, json_path, string(data), "XX").Result() //只在已存在时设置
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//default response success
	c.JSON(200, gin.H{
		"success": true,
	})
}

func delete_admin(c *gin.Context) {
	//uri
	uri := c.Param("uri")
	//delete
	json_path := uri_to_json_path(uri)
	r, err := redis_clinet.Do("json.del", key_object, json_path).Result()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if r == int64(0) {
		err = errors.New("may not exist")
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	}
	//default response success
	c.JSON(200, gin.H{
		"success": true,
	})
}
func delete_object(c *gin.Context) {
	//uri
	uri := c.Param("uri")
	//can only delete array item
	if !strings.HasSuffix(uri, ")") { //如果 uri 不是以括号结尾，代表要删除的是 kvs
		err := errors.New("can not do this")
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	}
	//delete array item
	json_path := uri_to_json_path(uri)
	r, err := redis_clinet.Do("json.del", key_object, json_path).Result()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if r == int64(0) {
		err = errors.New("may not exist")
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	}
	//default response success
	c.JSON(200, gin.H{
		"success": true,
	})
}

func uri_to_json_path(uri string) string {
	var r string
	uri = strings.TrimSuffix(uri, "/")
	uri = strings.TrimPrefix(uri, "/")
	r = strings.ReplaceAll(uri, "/", ".")
	r = strings.ReplaceAll(r, "(", "[")
	r = strings.ReplaceAll(r, ")", "]")
	r = "." + r
	return r
}
