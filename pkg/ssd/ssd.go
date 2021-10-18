package ssd

var bufferBytes []byte
var Dir = "ssds/"
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
	return nil
}
