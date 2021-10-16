package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	bufferBytes := []byte(
		`{
		"name": "http ssd service",
		"logLevel": "trace",
		"admin": {
			"host": "127.0.0.1",
			"port": "8001",
			"cors": false
		},
		"data": {
			"host": "127.0.0.1",
			"port": "8000",
			"cors": true,
			"level": 0
		}
	}`)
	data := gjson.GetBytes(bufferBytes, "data.level.c").String()
	if data == "" {
		fmt.Println("not exist")
	}
	fmt.Println(data)
}
