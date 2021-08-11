package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Service struct {
		Name string `json:"name"`
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"service"`
}

var Values Config

var configDir = "configs/"
var defaultConfigFileName = "default.json"
var currentConfigFileName = "current.json"

func LoadDefault() error {
	var err error
	Values, err = read(configDir + defaultConfigFileName)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return err
}

func LoadCurrent() error {
	var err error
	Values, err = read(configDir + currentConfigFileName)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return err
}

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

func Get() Config {
	return Values
}

func Put() {

}
