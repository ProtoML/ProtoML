package main

import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"github.com/ProtoML/ProtoML/parsers"
	"github.com/ProtoML/ProtoML/types"
	"github.com/ProtoML/ProtoML/utils"
	"github.com/ProtoML/ProtoML-persist/persist"
	"github.com/ProtoML/ProtoML-persist/local"
	"github.com/ProtoML/ProtoML/pipeline"
	"bytes"
)

const (
	PROTOML_PERSIST_DIR = "ProtoML"
	PROTOML_PERSIST_PIPELINE = "pipeline.json"
)

func InitPersistState(storage persist.PersistStorage) (pipes *types.Pipeline, err error) {
	// touch ProtoML state in persist
	dirs, err := storage.ListDirectories() 
	utils.ErrorPanic(err)
	// setup main directory
	stateDir := false
	for _, dir := range(dirs) {
		if dir == PROTOML_PERSIST_DIR {
			stateDir = true
			break
		}
	}
	if !stateDir {
		err := storage.CreateDirectory(PROTOML_PERSIST_DIR)
		utils.PanicOnError(err, "Unable to create ProtoML state folder in PersistStorage")
		pipes = pipeline.NewPipeline()
		err = storage.CreateFile(PROTOML_PERSIST_DIR,PROTOML_PERSIST_PIPELINE)
	} else {
		// setup pipeline
		statePipeline := false
		files, err := storage.ListFiles(PROTOML_PERSIST_DIR)
		if err != nil {
			return pipes, err
		}
		for _, file := range(files) {
			if file == PROTOML_PERSIST_PIPELINE {
				statePipeline = true
				break
			}
		}
		if statePipeline {
			pipelineReader, err := storage.Load(PROTOML_PERSIST_DIR, PROTOML_PERSIST_PIPELINE)
			if err != nil {
				return pipes, err
			}
			pipelineBlob, err := ioutil.ReadAll(pipelineReader)
			if err != nil {
				return pipes, err
			}
			pipes, err = pipeline.LoadPipeline(pipelineBlob)
		} else {
			pipes = pipeline.NewPipeline()
			err = storage.CreateFile(PROTOML_PERSIST_DIR,PROTOML_PERSIST_PIPELINE)
		}
	}	  
	return
}
 
func main() {
	// execution flags
	var rootDir, configFile string
	flag.StringVar(&rootDir, "d", ".", "Directory of execution")
	flag.StringVar(&configFile, "config", "ProtoML.json", "Configuration file")
	flag.Parse()

	// load and parse config
	configFileReader, err := os.Open(configFile)
	utils.PanicOnError(err,
		fmt.Sprintf("Unable to open configuration file %v", configFile))
	defer configFileReader.Close()

	jsonBlob, err := ioutil.ReadAll(configFileReader)
	utils.PanicOnError(err,
		fmt.Sprintf("Unable to read configuration file %v", configFile))

	config, err := parsers.LoadConfig(rootDir, configFile, jsonBlob)
	utils.PanicOnError(err,
		fmt.Sprintf("Unable to parse configuration file %v", configFile))
	
	// setup PersistStorage
	var storage persist.PersistStorage
	storage = new(local.LocalStorage)
	err = storage.Init(config)
	utils.PanicOnError(err,
		"Cannot initilize PersistStorage")
	defer func() {
		err := storage.Close()
		if err != nil {
			utils.PanicOnError(err, "Unable to close PersistStorage properly")
		}
	}()	
		
	pipe, err := InitPersistState(storage)
	
	// store pipeline on close
	defer func() {
		pipelineBlob, err := pipeline.StorePipeline(pipe)
		utils.PanicOnError(err, "Unable to parse initial pipeline")
		if err != nil {
			utils.PanicOnError(err, "Unable to convert pipeline to json")
		}
		pipelineBuffer := bytes.NewBuffer(pipelineBlob)
		err = storage.Store(PROTOML_PERSIST_DIR,PROTOML_PERSIST_PIPELINE, pipelineBuffer)		
		if err != nil {
			utils.PanicOnError(err, "Unable to save pipeline to PersistStorage")
		}
	}()
	
	fmt.Println(config)
	fmt.Println(pipe)
	fmt.Println(storage)
}
