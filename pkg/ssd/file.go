package ssd

import (
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

// load ssd file
func Load(name string) error {
	path := Dir + name + ".json"
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fileString := string(fileBytes)
	// ssd file must be json
	if !gjson.Valid(fileString) {
		err := errors.New("ssd file format error")
		return err
	}
	buffer = fileString
	return nil
}

// save current ssd to json file
func Save() error {
	name := time.Now().Format("2006-01-02T15-04-05")
	path := Dir + name + ".json"
	fileBytes := []byte(buffer)
	err := ioutil.WriteFile(path, fileBytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

// list json files in ssd dir
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
	return names, err
}

// func Remove(name string) error {
// 	path := Dir + name + ".json"
// 	err := os.Remove(path)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func Read(name string) ([]byte, error) {
// 	path := Dir + name + ".json"
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }
