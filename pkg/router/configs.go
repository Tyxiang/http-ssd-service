package router

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	//uri := c.Param("uri")
	//q := c.Query("q")
	//s := c.Query("s")
	c.JSON(200, gin.H{
		"success": true,
		"data":    get(),
	})
}

func PUT(c *gin.Context) {
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
