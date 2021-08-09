package Logs

import (
	"github.com/spf13/viper"
)

func init(){
	config := viper.New()
	config.AddConfigPath("./configs")
	config.SetConfigType("json")
	//load current config file
	config.SetConfigName("current") // name of config file (without extension)
	err := config.ReadInConfig()
	if err != nil {
		if _, no := err.(viper.ConfigFileNotFoundError); no {
			//config file not found
			//load default config
			config.SetConfigName("default")
			err := config.ReadInConfig()
			if err != nil {
				fmt.Fprintln(gin.DefaultWriter, err.Error())
				panic(err)
			}
		} else {
			//config file was found, but another error was produced
			//program end
			fmt.Fprintln(gin.DefaultWriter, err.Error())
			panic(err)
		}
	}
}

