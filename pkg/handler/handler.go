package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

func uri_to_path(uri string) string {
	//fmt.Println(uri)
	path := strings.Trim(uri, "/")
	words := [][]string{
		{"/", "."},
		{"[]", ".-1"},
		{"[", "."},
		{"]", ""},
		{"$", "#"},
		{"_", "?"},
		{"..", "\\."}, // {"fav\.movie":"Deer Hunter"}
		// {"%7B", "{"},
		// {"%7D", "}"},
		{"%22", "\""},
		{"%3E", ">"},
		{"%3C", "<"},
	}
	for i := range words {
		path = strings.Replace(path, words[i][0], words[i][1], -1)
	}
	fmt.Println(path)
	return path
}

func validJson(data []byte) error {
	//err := fastjson.ValidateBytes(data)
	if !gjson.Valid(string(data)) {
		err := errors.New("wrong data type")
		return err
	}
	return nil
}
