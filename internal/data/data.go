package data

type Config struct {
	Service struct {
		Name string `json:"name"`
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"service"`
}

var Configs Config
