package log

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	// "github.com/spf13/viper"
	// "Configs/"
)

var data string = "data"

func init() {

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
