package main

import (
	"bytes"
	"net/http"
)

var (
	ip = "http://127.0.0.1"
)

func main() {
	get_8001()
	post_8001()
}

func post_8001() {
	println()
	port := "8000"
	paths := []string{
		"/",
	}
	data := []byte(`{"a":1,"b":2}`)
	for _, path := range paths {
		uri := ip + ":" + port + path
		print("POST ", uri)
		response, err := http.Post(uri, "application/json", bytes.NewBuffer(data))
		if err != nil {
			println(" faild!")
		}
		defer response.Body.Close()
		println(" ......", response.StatusCode)
	}
}

func get_8001() {
	println()
	port := "8001"
	paths := []string{
		"/",
		"/config",
		"/logs",
		"/ssds",
		"/scripts",
	}
	for _, path := range paths {
		uri := ip + ":" + port + path
		print("GET ", uri)
		response, err := http.Get(uri)
		if err != nil {
			println(" faild!")
		}
		defer response.Body.Close()
		println(" ......", response.StatusCode)
	}
	// for key, value := range response.Header {
	// 	fmt.Printf("%s=>%s\n", key, value)
	// }

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	println(" faild!")
	// }
	// println(string(body))
}
