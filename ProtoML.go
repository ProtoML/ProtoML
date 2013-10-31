package main

import (
	"flag"
	"github.com/ProtoML/ProtoML/logger"
	"github.com/ProtoML/ProtoML/protomlserver"
	"github.com/ProtoML/ProtoML-persist/persist"
	"os"
	"log"
)

const LOGTAG = "ProtoML-Main"

func main() {
	// execution flags
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "ProtoML.json", "Configuration file")
	flag.Parse()
 
	// load config
	config, err := persist.LoadConfig(configFilePath)
	if err != nil {
		logger.LogFatal(LOGTAG, err, "Cannot open configuration file %s", configFilePath)
		return
	}
	
	// setup logger
	logger.Debug = true
	logger.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// start server
	err = protomlserver.ProtoMLServer(config)
	if err != nil {
		logger.LogFatal(LOGTAG, err, "ProtoML Server Failed")
	}
}
