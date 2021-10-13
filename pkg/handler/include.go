package handler

import "strings"

func uri_to_path(uri string) string {
	path_1 := strings.Trim(uri, "/")
	path_2 := strings.Replace(path_1, "/", ".", -1)
	path_3 := strings.Replace(path_2, "()", ".-1", -1)
	path_4 := strings.Replace(path_3, "(", ".", -1)
	path := strings.Replace(path_4, ")", "", -1)
	//fmt.Println(path)
	return path
}
