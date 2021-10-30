package config

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/tidwall/gjson"
)

// load a config file
func Load(name string) error {
	path := Dir + name + ".json"
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fileString := string(fileBytes)
	// config file must be json
	if !gjson.Valid(fileString) {
		err = errors.New("config file format error")
		return err
	}
	buffer = fileString
	return nil
}

// save current config to json file
func Save() error {
	name := "last"
	path := Dir + name + ".json"
	fileBytes := []byte(buffer)
	err := ioutil.WriteFile(path, fileBytes, 0666) //存在就覆盖；不存在创建。
	if err != nil {
		return err
	}
	return nil
}

// list json files in config dir
func List() ([]string, error) {
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		return nil, err
	}
	name := ""
	names := []string{}
	for _, file := range files {
		name = strings.TrimRight(file.Name(), ".json")
		names = append(names, name)
	}
	return names, nil
}

// //new a json config file
// func New(name string, data []byte) error {
// 	path := Dir + name + ".json"
// 	err := ioutil.WriteFile(path, data, 0644) //存在就覆盖；不存在创建。
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //delete a json config file
// func Remove(name string) error {
// 	path := Dir + name + ".json"
// 	err := os.Remove(path)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
