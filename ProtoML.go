package main

import (
	"flag"
	"github.com/ProtoML/ProtoML/logger"
	"github.com/ProtoML/ProtoML/protomlserver"
	"github.com/ProtoML/ProtoML-persist/persist/persistparsers"
	"os"
	"log"
)

const LOGTAG = "ProtoML-Main"

func main() {
	// execution flags
	var configFilePath string
	var validateTransforms bool
	flag.StringVar(&configFilePath, "config", "ProtoML.json", "Configuration file")
	flag.BoolVar(&validateTransforms, "validatetransforms", false, "Used to add and validate all transforms files")
	flag.Parse()
 
	// load config
	config, err := persistparsers.LoadConfig(configFilePath)
	if err != nil {
		logger.LogFatal(LOGTAG, err, "Cannot open configuration file %s", configFilePath)
		return
	}
	
	// setup logger
	logger.Debug = true
	logger.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// start server
	quitChan, err := protomlserver.ProtoMLServer(config, validateTransforms)
	defer func() { quitChan <- true }() // close server
	if err != nil {
		logger.LogFatal(LOGTAG, err, "ProtoML Server Failed")
	}
}
