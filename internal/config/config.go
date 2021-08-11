package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"http-object/internal/data"

	"github.com/gin-gonic/gin"
)

var configDir = "configs/"
var defaultConfigFileName = "default.json"
var currentConfigFileName = "current.json"

func load(configPath string) (c data.Config, e error) {
	var config data.Config
	configFile, err := os.Open(configPath)
	if err != nil {
		// fmt.Println(err)
		e = err
		return
	}
	defer configFile.Close()
	configString, err := ioutil.ReadAll(configFile)
	if err != nil {
		// fmt.Println(err)
		e = err
		return
	}
	err = json.Unmarshal([]byte(configString), &config)
	if err != nil {
		// fmt.Println(err)
		e = err
		return
	}
	return config, nil
}

func Load() (c data.Config, e error) {
	var config data.Config
	config, err := load(configDir + currentConfigFileName)
	if err != nil {
		// fmt.Println(err)
		config, err = load(configDir + defaultConfigFileName)
	}
	if err != nil {
		// fmt.Println(err)
		e = err
		return
	}
	return config, nil
}

func Save(c data.Config) (e error) {
	return e
}

func Get(c *gin.Context) {
	//uri
	//uri := c.Param("uri")
	//query
	//q := c.Query("q")
	//s := c.Query("s")
	//default response success
	//c.Header("Content-Type", "application/json; charset=utf-8")
	//c.String(200, "{\"success\": true, \"data\":"+data.Configs.(string)+"}")
	c.JSON(200, gin.H{
		"success": true,
		"data":    data.Configs,
	})
}

func Put(c *gin.Context) {
	//uri
	//uri := c.Param("uri")
	//get data
	//json_path := uri_to_json_path(uri)
	//data, err := c.GetRawData()
	// if err != nil {
	// 	fmt.Fprintln(gin.DefaultWriter, err.Error())
	// 	c.JSON(500, gin.H{
	// 		"success": false,
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }
	// //put
	// _, err = redis_clinet.Do("json.set", key_object, json_path, string(data), "XX").Result() //只在已存在时设置
	// if err != nil {
	// 	fmt.Fprintln(gin.DefaultWriter, err.Error())
	// 	c.JSON(500, gin.H{
	// 		"success": false,
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }
	//default response success
	c.JSON(200, gin.H{
		"success": true,
	})
}
