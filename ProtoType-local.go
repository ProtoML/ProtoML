package main

import (
	"flag"
	"fmt"
	"github.com/ProtoML/ProtoML-persist/local"
	"github.com/ProtoML/ProtoML-persist/persist"
	"github.com/ProtoML/ProtoML/types"
	"os"
)

type stringSlice []string

func (i *stringSlice) String() string {
	return fmt.Sprint(*i)
}

func (i *stringSlice) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func parseRunRequest(storage persist.PersistStorage) (runRequest types.RunRequest, ok bool) {
	runRequest.DataNamespace = *flag.String("namespace", "", "Data namespace.")
	runRequest.TransformName = *flag.String("transform", "", "Filename of transform.")
	runRequest.JsonParameters = *flag.String("parameters", "", "JSON file containing parameters for the transform.")
	var dataIds stringSlice
	flag.Var(&dataIds, "data", "List of data ids in the intended ordering.")
	flag.Parse()

	runRequest.Data = make([]types.Data, len(dataIds))
	for idx, dataId := range dataIds {
		data, err := storage.LoadData(dataId)
		if err != nil {
			ok = false
			return
		}
		runRequest.Data[idx] = data
	}

	if runRequest.DataNamespace == "" || runRequest.TransformName == "" || runRequest.JsonParameters == "" {
		ok = false
		return
	} else {
		ok = true
	}
	return
}

func main() {
	// setup PersistStorage
	var storage persist.PersistStorage
	storage = new(local.LocalStorage)

	// create run request
	runRequest, ok := parseRunRequest(storage)
	if !ok {
		os.Exit(1)
	}

	// TODO some call taking in runRequest and storage
	fmt.Println(runRequest)
}
