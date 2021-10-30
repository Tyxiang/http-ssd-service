package ssd

import (
	"testing"
)

func init() {
	buffer =
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
	  }`
}

// func TestGet(t *testing.T) {
// 	data, err := Get("")
// 	fmt.Println(data, err)
// }

func BenchmarkAdd(b *testing.B) {
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Add("test", []byte("abc"))
	}
}

func BenchmarkGet(b *testing.B) {
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Get("system")
	}
}

func BenchmarkSet(b *testing.B) {
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Set("test", []byte("abc"))
	}
}
