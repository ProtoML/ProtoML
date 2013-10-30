package protomlserver

import (
	"github.com/ProtoML/ProtoML-persist/local"
	"github.com/ProtoML/ProtoML-persist/persist"
	"github.com/ProtoML/ProtoML/logger"
	//	"github.com/ProtoML/ProtoML/types"
	"github.com/ProtoML/ProtoML/formatadaptor"
	"github.com/ProtoML/ProtoML/utils"
	"errors"
	"fmt"
)

const LOGTAG = "ProtoML-Server"

func errorMsg(err error, msg string) error {
	return errors.New(fmt.Sprintf("%s: %v", msg, err))
}

func ProtoMLServer(config persist.Config) (err error) {
	logger.LogInfo(LOGTAG, "Starting Server")

	// check protoml directory existance
	_, err = utils.ProtoMLDir()
	if err != nil {
		err = errorMsg(err, "Cannot use enviromental variable PROTOMLDIR")
		return
	}

	// setup config
	config.FormatCollection = formatadaptor.DefaultFileFormatCollection()
	logger.LogDebug(LOGTAG, "Formats Available:")
	for _, format := range config.FormatCollection.ListAdaptors() {
		logger.LogDebug(LOGTAG, "\t"+format)
	}

	var storage persist.PersistStorage
	storage = new(local.LocalStorage)
	err = storage.Init(config)
	defer func() {
		storage.Close()
		//logger.LogFatal(LOGTAG, err, "Could not properly shutdown persistance")
	}()
	if err != nil {
		err = errorMsg(err, "Cannot create persistence layer")
		return
	}	

	return nil
}
