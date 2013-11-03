package osutils

import (
	"os"
	"io/ioutil"
	"io"
	"path"
	"github.com/ProtoML/ProtoML/logger"
	"crypto/md5"
	"fmt"
)

const (
	LOGTAG = "OS-Utilities"
)

func LoadBlob(filename string) (blob []byte, err error) {
	fileReader, err := os.Open(filename)
	if err != nil {
		return
	}
	defer fileReader.Close()
	blob, err = ioutil.ReadAll(fileReader)
	return
}

func MD5Hash(anything ...interface{}) string {
	// returns the md5 hash of anything that can be printed as a string
	h := md5.New()
	io.WriteString(h, fmt.Sprint(anything...))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func TouchDir(dir string) (err error) {
	err = os.MkdirAll(dir, os.ModePerm)
	logger.LogDebug(LOGTAG,"Touching directory %s",dir)
	if err != nil && !os.IsExist(err) {
		return
	} else {
		err = nil
		return
	}
}

func TouchFile(filepath string) (file *os.File, err error) {
	file, err = os.Create(filepath)
	logger.LogDebug(LOGTAG,"Touching file %s",filepath)
	if err != nil && !os.IsExist(err) {
		return
	} else {
		err = nil
		return
	}
}

func PathExists(fullPath string) bool {
	// checks if a file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func ListFilesInDirectory(directory string) (list []string, err error) {
	directoryFileDescriptor, err := os.Open(directory)
	defer directoryFileDescriptor.Close()
	if err != nil {	return }

	files, err := directoryFileDescriptor.Readdir(0)
	list = make([]string, len(files))
	listIter := 0
	for _, file := range files {
		if !file.IsDir() {
			list[listIter] = file.Name()
			listIter++
		}
	}
	return
}

func CopyDirectory(srcDirectory, dstDirectory string) (err error){
	dirFiles, err := ListFilesInDirectory(srcDirectory)
	if err != nil {
		return
	}
	
	for _, file := range dirFiles {
		srcFile := path.Join(srcDirectory, file)
		fr, err := os.Open(srcFile)
		if err != nil {
			return err
		}
		defer fr.Close()
		dstFile := path.Join(dstDirectory, file)
		fw, err := TouchFile(dstFile)
		if err != nil {
			return err
		}
		defer fw.Close()
		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}
	}
	return 
}
