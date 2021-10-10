package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
		//"github.com/tidwall/sjson"
	//"github.com/tidwall/gjson"
)

var Config

var configDir = "configs/"
var defaultConfigFileName = "default.json"
var currentConfigFileName = "current.json"

func read(configPath string) (c Config, err error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		// fmt.Println(err)
		return
	}
	defer configFile.Close()
	configString, err := ioutil.ReadAll(configFile)
	if err != nil {
		// fmt.Println(err)
		return
	}
	err = json.Unmarshal([]byte(configString), &c)
	if err != nil {
		// fmt.Println(err)
		return
	}
	return c, err
}

func save(configPath string, c Config) error {
	configString, err := json.Marshal(c)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(configPath, configString, 0644) //存在就覆盖；不存在创建。
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return err
}

func get() Config {
	return Values
}

func put() {

}
