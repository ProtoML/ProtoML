package protomlserver

import (
	"testing"
	"os"
	"github.com/ProtoML/ProtoML/utils"
	"github.com/ProtoML/ProtoML/utils/osutils"
	"path"
	"time"
	"github.com/ProtoML/ProtoML-persist/persist/persistparsers"
	"github.com/ProtoML/ProtoML/tests"
	"fmt"
)

func createTestDir(protomlSrcPath string) (testPath string, err error) {
	protomlDir, err := utils.ProtoMLDir()
	srcPath := path.Join(protomlDir, protomlSrcPath)
	testPath = path.Join(os.TempDir(),"ProtoML",fmt.Sprintf("%d",time.Now().Unix()))
	err = osutils.TouchDir(testPath)
	if err != nil {
		return
	}
	err = osutils.CopyDirectory(srcPath, testPath)
	return
}

func destroyTestDir(testPath string) (err error) {
//	err = os.RemoveAll(testPath)
	return
}

const (
	CONFIG_ERR_MSG = "Cannot use configuration file %s\nerror: %s"
	EMPTY_DIR = "ProtoML/tests/testsets/empty"
	SYNTHETIC_DIR = "ProtoML/tests/testsets/synthetic"
)

func DatasetTestBase(t *testing.T, protomlDir, protomlJson string) {
	tests.SetupLogger(t)

	// setup test directory and config
	testDir, err := createTestDir(protomlDir)
	configFilePath := path.Join(testDir, protomlJson)
	t.Logf("Test original dir: %s", protomlDir)
	t.Logf("Test config at: %s", configFilePath)
	defer destroyTestDir(testDir)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	// load config
	config, err := persistparsers.LoadConfig(configFilePath)
	if err != nil {
		t.Fatalf(CONFIG_ERR_MSG, path.Base(configFilePath), err)
		return
	}
	config.LocalPersistStorage.RootDir = testDir
	config.LocalPersistStorage.DatasetDirectory = testDir
	config.LocalPersistStorage.ElasticPort = 9400

	// start server
	err = ProtoMLServer(config)
	if err != nil {
		t.Fatalf("ProtoML Server Error\nerror: %s",err)
	}
}
/*
func TestEmpty(t *testing.T) {
	DatasetTestBase(t, EMPTY_DIR, "ProtoML.json")
}*/

func SyntheticTestBase(t *testing.T, protomlJson string) {
	DatasetTestBase(t, SYNTHETIC_DIR, protomlJson)
}

func TestSyntheticStartup(t *testing.T) {
	SyntheticTestBase(t, "ProtoML_Startup.json")
}

