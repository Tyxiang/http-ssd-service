package handler

import (
	"errors"
	"strings"

	"github.com/tidwall/gjson"
)

func uri_to_path(uri string) string {
	path_1 := strings.Trim(uri, "/")
	path_2 := strings.Replace(path_1, "/", ".", -1)
	path_3 := strings.Replace(path_2, "()", ".-1", -1)
	path_4 := strings.Replace(path_3, "(", ".", -1)
	path := strings.Replace(path_4, ")", "", -1)
	//fmt.Println(path)
	return path
}

func validJson(data []byte) (err error) {
	//err := fastjson.ValidateBytes(data)
	if !gjson.Valid(string(data)) {
		err := errors.New("wrong data type")
		return err
	}
	return nil
}
