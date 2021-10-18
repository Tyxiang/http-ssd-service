package config

var bufferBytes []byte
var Dir = "configs/"
var Warn error

func Init() error {
	err := Load("last")
	if err != nil {
		Warn = err
		err = Load("default")
		if err != nil {
			return err
		}
	}
	return err
}
