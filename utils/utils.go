package utils

import (
	"os"
	"errors"
)

func ProtoMLDir() (protomlDir string, err error) {
    // check protoml directory existance
	protomlDir = os.Getenv("PROTOMLDIR")
	if len(protomlDir) == 0 {
		err = errors.New("Cannot find PROTOML_DIR enviroment variable")
		return
	}
	_, dirErr := os.Stat(protomlDir)
	protomlDirExists := os.IsNotExist(dirErr)
	if !protomlDirExists {
		err = dirErr
		return	
	}
	return
}

