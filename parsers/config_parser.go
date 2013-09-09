package parsers

import (
//	"fmt"
	"encoding/json"
	"github.com/ProtoML/ProtoML/types"
)

type ConfigParameters struct {
	TemplatePaths []string
	DisableProtoMLTemplates bool
	AddtionalDataTypes []types.DataType
}

type Config struct {
	ConfigFile		        string
	RootDir		    	    string
	Parameters              ConfigParameters
}

func LoadConfig(rootDir, configFile string, configFileBlob []byte) (config Config, err error) {
	config = Config{RootDir: rootDir, ConfigFile: configFile}
	err = json.Unmarshal(configFileBlob, &config.Parameters)
	if err != nil {
		return
	}
	return
}
