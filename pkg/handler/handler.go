package handler

import (
	"strings"
)

func parse(uri string) (string, string) {
	println(uri)
	// remove "%20"(space)
	for strings.HasSuffix(uri, "%20") {
		uri = strings.TrimSuffix(uri, "%20")
	}
	//println(uri)
	var path string
	var property string
	if strings.HasSuffix(uri, ".type") {
		path = strings.TrimSuffix(uri, ".type")
		property = "type"
	} else {
		path = uri
	}
	path = strings.Trim(path, "/") // remove "/"
	words := [][]string{
		{".", "\\."}, // input "fav.movie" output "fav\.movie" for {"fav.movie": 123}
		{"/", "."},
		{"$", "#"},
		{"_", "?"},
		{"%22", "\""},
		{"%3E", ">"},
		{"%3C", "<"},
		// {"%7B", "{"},
		// {"%7D", "}"},
		// {"()", ".-1"},
	}
	for i := range words {
		path = strings.ReplaceAll(path, words[i][0], words[i][1])
	}
	println(path)
	return path, property
}
