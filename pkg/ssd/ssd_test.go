package ssd

import "testing"

func init() {
	bufferBytes = []byte(
		`{
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
	  }`)
}

func BenchmarkGet(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		Get("users.0.name")
	}
}
