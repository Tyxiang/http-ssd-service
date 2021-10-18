package main

import "github.com/tidwall/sjson"

const json = `{"name":["abc","123"],"age":47}`

func main() {
	value, _ := sjson.Set(json, "name.-1", "Anderson")
	println(value)
}
