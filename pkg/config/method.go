package config

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
