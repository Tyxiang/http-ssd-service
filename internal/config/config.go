package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Service struct {
		Name string `json:"name"`
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"service"`
}

var Values Config

var configDir = "configs/"
var defaultConfigFileName = "default.json"
var currentConfigFileName = "current.json"

func LoadDefault() error {
	var err error
	Values, err = read(configDir + defaultConfigFileName)
	return err
}

func LoadCurrent() error {
	var err error
	Values, err = read(configDir + currentConfigFileName)
	return err
}

func read(configPath string) (c Config, err error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		// fmt.Println(err)
		return
	}
	defer configFile.Close()
	configString, err := ioutil.ReadAll(configFile)
	if err != nil {
		// fmt.Println(err)
		return
	}
	err = json.Unmarshal([]byte(configString), &c)
	if err != nil {
		// fmt.Println(err)
		return
	}
	return c, err
}

func save(configPath string, c Config) (e error) {
	configString, err := json.Marshal(c)
	if err != nil {
		// fmt.Println(err)
		e = err
		return
	}
	err = ioutil.WriteFile(configPath, configString, 0644) //存在就覆盖；不存在创建。
	if err != nil {
		// fmt.Println(err)
		e = err
		return
	}
	return e
}

func Get(c *gin.Context) {
	//uri := c.Param("uri")
	//q := c.Query("q")
	//s := c.Query("s")
	c.JSON(200, gin.H{
		"success": true,
		"data":    Values,
	})
}

func Put(c *gin.Context) {
	//uri := c.Param("uri")
	configString, err := c.GetRawData()
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	err = json.Unmarshal([]byte(configString), &Values)
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	err = save(configDir+currentConfigFileName, Values)
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
}
