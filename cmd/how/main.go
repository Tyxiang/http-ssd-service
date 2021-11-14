package main

import (
	"github.com/tidwall/gjson"
)

var buffer = `{
		"name": "http ssd service",
		"logLevel": "trace",
		"system": {
		"host": "127.0.0.1",
		"port": "8001",
		"cors": false
	},
	"ssd": {
		"host": "127.0.0.1",
		"port": "8000",
		"cors": true,
		"level": 0
	}
}`

func main() {
	data := gjson.Get(buffer, "system").Indexes
	println(data)
}
