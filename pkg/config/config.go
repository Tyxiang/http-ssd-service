package config

var bufferBytes []byte
var Dir = "configs/"
var Warn error

func Init() (err error) {
	err = Load("last")
	if err != nil {
		Warn = err
		err = Load("default")
		if err != nil {
			return
		}
	}
	return
}
