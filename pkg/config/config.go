package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

var configBytes []byte
var dir = "configs/"

// file
func Load(n string) (err error) {
	path := dir + n + ".json"
	configBytes, err = ioutil.ReadFile(path)
	if !gjson.Valid(string(configBytes)) {
		return err
	}
	return nil
}

func List() (f []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		f = append(f, strings.TrimRight(file.Name(), ".json"))
	}
	return f, nil
}

func New(n string, d []byte) (err error) {
	path := dir + n + ".json"
	err = ioutil.WriteFile(path, d, 0644) //存在就覆盖；不存在创建。
	if err != nil {
		return err
	}
	return nil
}

func Remove(n string) (err error) {
	path := dir + n + ".json"
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func Save(n string) {

}

// data
func Add(u string) (err error) {
	return
}

func Get(p string) (r interface{}) {
	//r = gjson.ParseBytes(configBytes).Value()
	if p == "" {
		p = "@this"
	}
	r = gjson.GetBytes(configBytes, p).Value()
	return r
}

func Set(u string, d interface{}) (err error) {
	return
}

func Del(u string) (err error) {
	return
}

// func read(path string) (c []byte, err error) {
// 	c, err = ioutil.ReadFile(path)
// 	if err != nil {
// 		return
// 	}
// 	if !gjson.Valid(string(configBytes)) {
// 		fmt.Fprintln(gin.DefaultWriter, "load default config")
// 		configFilePath = configDir + "default.json"
// 		configBytes, err = ioutil.ReadFile(configFilePath)
// 		if err != nil {
// 			fmt.Fprintln(gin.DefaultWriter, err.Error())
// 			panic(err)
// 		}
// 	}
// 	return c, err
// }

// func save(configPath string, c Config) error {
// 	configString, err := json.Marshal(c)
// 	if err != nil {
// 		// fmt.Println(err)
// 		return err
// 	}
// 	err = ioutil.WriteFile(configPath, configString, 0644) //存在就覆盖；不存在创建。
// 	if err != nil {
// 		// fmt.Println(err)
// 		return err
// 	}
// 	return err
// }
