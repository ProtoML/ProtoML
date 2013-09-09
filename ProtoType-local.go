package main

import (
	"flag"
	"fmt"
	"github.com/ProtoML/ProtoML-persist/local"
	"github.com/ProtoML/ProtoML-persist/persist"
	"github.com/ProtoML/ProtoML/types"
	"os"
)

func parseRunRequest() (runRequest types.RunRequest, ok bool) {
	var runRequest RunRequest
	runRequest.DataNamespace = flag.String("data-namespace", "", "Data namespace.")
	runRequest.TransformName = flag.String("transform", "", "Filename of transform.")
	runRequest.JsonParameters = flag.String("parameters", "", "JSON file containing parameters for the transform.")
	dataIds := flag.String("data", "", "Comma separated list of data ids in the proper ordering.")
	flag.Parse()
	// TODO parse dataIds into a data list

	if runRequest.DataNamespace || runRequest.TransformName || runRequest.JsonParameters == "" {
		ok = false
		return
	} else {
		ok = true
	}
	return
}

func main() {
	runRequest, ok := parseRunRequest()
	if !ok {
		os.Exit(1)
	}

	// setup PersistStorage
	var storage persist.PersistStorage
	storage = new(local.LocalStorage)

	// TODO some call taking in runRequest and storage
}
