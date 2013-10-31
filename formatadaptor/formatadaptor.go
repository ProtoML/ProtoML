package formatadaptor

import (
	"errors"
	"fmt"
	"github.com/ProtoML/ProtoML/types"
	"github.com/ProtoML/ProtoML/dependency"
	"github.com/ProtoML/ProtoML/logger"
	"path"
)

const LOGTAG = "FormatAdaptor"
 
type FileFormatAdaptor interface {
	Dependencies() (dependencies []dependency.Dependency)
	Split(srcFile string, dstFiles []string) (err error)
	Join(srcFiles []string, dstFile string) (err error)
	ToRaw(srcFile string, dstFile string) (err error)
	FromRaw(srcFile string, dstFile string) (err error)
	Shape(path string) (nrows, ncols int, err error)
}

type FileFormatCollection struct {
	adaptors map[string]FileFormatAdaptor
}

func (fc *FileFormatCollection) RegisterAdaptor(formattype string, adaptor FileFormatAdaptor) (err error) {
	/* Allow overwriting of adaptors
	if _, ok := fc.adaptors[formattype]; ok {
		err = errors.New(fmt.Sprintf("Format type %s already registered", formattype))
		return
	} */
	fc.adaptors[formattype] = adaptor
	return
}

func (fc *FileFormatCollection) UnregisterAdaptor(formattype string) (err error) {
	if _, ok := fc.adaptors[formattype]; !ok {
		err = errors.New(fmt.Sprintf("Format type %s not registered", formattype))
		return
	}
	delete(fc.adaptors, formattype)
	return
}
 
func (fc *FileFormatCollection) GetAdaptor(formattype string) (adaptor FileFormatAdaptor, err error) {
	adaptor, ok := fc.adaptors[formattype]
	if !ok {
		err = errors.New(fmt.Sprintf("Format type %s not registered", formattype))
		adaptor = nil
	}
	return adaptor, err
}

func (fc *FileFormatCollection) ListAdaptors() (adaptors []string) {
	for adaptor, _ := range fc.adaptors {
		adaptors = append(adaptors, adaptor)
	}
	return
}

// segment columns into data groups and a map of group index to original column index
func SplitColumns(cols types.DatasetColumns, ncols int) (groups []types.DataGroupColumns, groupColsIndex [][]int) {                         
	// tag: [index] -> index: [tag]
	indexTags := make(map[int][]string)
	for tag, indexes := range cols.Tags {
		for _, index := range indexes {
			indexTags[index] = append(indexTags[index], tag)
		}
	}
	groups = make([]types.DataGroupColumns, len(cols.ExclusiveTypes))
	groupColsIndex = make([][]int, len(cols.ExclusiveTypes)) //group index -> [raw col index]
	i := 0
	for typ, indexes := range cols.ExclusiveTypes {
		groups[i].ExclusiveType = typ
		groups[i].Tags = make([][]string, len(indexes))
		for groupIndex, index := range indexes {
			// fill out tags for column indexes
			if tags, ok := indexTags[index]; ok {
				groups[i].Tags[groupIndex] = append(groups[i].Tags[groupIndex], tags...)
			}
			// group to col index map
			groupColsIndex[i] = append(groupColsIndex[i],index)
		}
		i++
	}
	return
}

func (fc *FileFormatCollection) Split(dataset types.DatasetFile, dstDir string) (dataGroups []types.DataGroup, splitFiles []string, groupColsIndex [][]int, err error) {
	if _, ok := fc.adaptors[dataset.FileFormat]; !ok {
		err = errors.New(fmt.Sprintf("Format type %s not registered", dataset.FileFormat))
		return
	}
	adaptor, _ := fc.adaptors[dataset.FileFormat]

	// get dataset shape
	ncols, nrows, err := adaptor.Shape(dataset.Path)
	if err != nil {
		return
	}
	if nrows != dataset.NRows {
		err = errors.New(fmt.Sprintf("Dataset number of rows(%d) and the number of rows the adaptor found(%d) are different", dataset.NRows, nrows))
	}
	if err != nil {
		return
	}
	if ncols != dataset.NCols {
		err = errors.New(fmt.Sprintf("Dataset number of columns(%d) and the number of columns the adaptor found(%d) are different", dataset.NCols, ncols))
	}
	if err != nil {
		return
	}

	// construct DataGroups and index map to groups
	// groupColsIndex maps groupIndex to raw column index
	groupCols, groupColsIndex := SplitColumns(dataset.Columns, ncols)
	logger.LogDebug(LOGTAG, "DataGroup Columns for %s", path.Base(dataset.Path))
	for _, col := range groupCols {
		logger.LogDebug(LOGTAG, "\t%v", col)
	}
	

	// split dataset into columns
	splitFiles = make([]string,dataset.NCols)
	for i, _ := range splitFiles {
		splitFiles[i] = path.Join(dstDir,fmt.Sprintf("%010d.%s",i,dataset.FileFormat))
	}
	err = adaptor.Split(dataset.Path, splitFiles)
	if err != nil {
		return
	}
	if ncols != len(splitFiles) {
		err = errors.New(fmt.Sprintf("Format adaptor for %s did not split into single columns, found %d column files while needing %d columns", dataset.FileFormat, len(splitFiles), ncols))
		return
	}
	var dncols, dnrows int
	for _, sfile := range splitFiles {
		dncols, dnrows, err = adaptor.Shape(sfile)
		if err != nil {
			return
		}
		if dncols != 1 {
			err = errors.New(fmt.Sprintf("Format adaptor for %s did not split into single columns, found %d column files while needing 1 column", dataset.FileFormat, dncols))
			return
		}
		if dnrows != nrows {
			err = errors.New(fmt.Sprintf("Format adaptor for %s did not conserve number rows, found %d rows while needing %d rows", dataset.FileFormat, dnrows, nrows))
			return
		}
	}

	// construct DataGroups
	dataGroups = make([]types.DataGroup, len(groupCols))
	for i, _ := range groupCols {
		group := groupCols[i]
		dataGroup := types.DataGroup{}
		dataGroup.FileFormat = dataset.FileFormat
		dataGroup.NCols = len(group.Tags)
		dataGroup.NRows = nrows
		dataGroup.Columns = group
		dataGroup.Source = path.Base(dataset.Path)
		dataGroups[i] = dataGroup
	}
	return
}
