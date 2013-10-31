package delimiteradaptor
 
import (
	"errors"
    "os"
    "encoding/csv"
	"io"
	"github.com/ProtoML/ProtoML/utils/osutils"
	"github.com/ProtoML/ProtoML/dependency"
	"github.com/ProtoML/ProtoML/logger"
)

type DelimiterAdaptor struct {
	delimiter rune
}
 
func New(delimiter rune) *DelimiterAdaptor {
	logger.LogInfo("DEMI", "Making delimiter adaptor with %s", delimiter)
	return &DelimiterAdaptor{delimiter}
}

func (adaptor *DelimiterAdaptor) Dependencies() (dependencies []dependency.Dependency) {
    dependencies = []dependency.Dependency{}
    return
}
 
func (adaptor *DelimiterAdaptor) Split(srcPath string, dstPaths []string) (err error) {
	ncols := len(dstPaths)
	if ncols <= 0 {
		err = errors.New("Output split files are over a non-positive number of columns")
		return
	}

	// setup files and readers
	srcFile, err := os.Open(srcPath)
	if err != nil { return err }
	defer srcFile.Close()
 
	reader := csv.NewReader(srcFile)
	reader.Comma = adaptor.delimiter
	reader.FieldsPerRecord = ncols
	
	writers := make([]*csv.Writer,ncols)
	writerQueue := make([]chan string,ncols)
	writerErrors := make([]chan error,ncols)
	for i, dstPath := range dstPaths {
		dstFile, err := osutils.TouchFile(dstPath)
		if err != nil { return err }
		defer dstFile.Close()
		writers[i] = csv.NewWriter(dstFile)
		writerQueue[i] = make(chan string)
		defer close(writerQueue[i])
		writerErrors[i] = make(chan error)
	}
	
	// setup column writer threads
	for i, _ := range writers {
		go func(writer *csv.Writer, writeQueue <-chan string, errChan chan<- error) {
			for {
				select {
				case val, ok := <- writeQueue:
					if !ok {
						writer.Flush()
						return
					}
					// write values
					err := writer.Write([]string{val})
					if err != nil {
						errChan <- err
						return
					}
					writer.Flush()
					errChan <- err
				}
			}
		}(writers[i], writerQueue[i], writerErrors[i])
	}

	// read/concurrent write loop
	err = nil
	var row []string
	for err == nil {
		row, err = reader.Read()
		if err != nil {
			break
		}
		if len(row) != ncols {
			return errors.New("Non-uniform columns in source file in split")
		}
		for i, val := range row {
			writerQueue[i] <- val
		}
		// wait for writers to complete
		writerCompletes := ncols
		for {
			for _, errChan := range writerErrors {
				select {
				case err, _ := <- errChan:
					if err != nil {
						return err
					}
					writerCompletes -= 1
				}
			}
			if writerCompletes == 0 {
				break
			}
		}
	}
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
 
type rowVal struct {
	val string
	index int
	err error
}

func (adaptor *DelimiterAdaptor) Join(srcPaths []string, dstPath string) (err error) {
	ncols := len(srcPaths)
	if ncols == 0 {
		return errors.New("Need at least one column")
	}

	dstFile, err := osutils.TouchFile(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	// joined writer
	writer := csv.NewWriter(dstFile)
	writer.Comma = adaptor.delimiter
	
	// readers
	readers := make([]*csv.Reader,ncols)
	// reader read queue
	readerQueue := make(chan rowVal)
	defer close(readerQueue)
	// signal to continue reading
	readerSignal := make([]chan int,ncols)
	for i, srcPath := range srcPaths {
		srcFile, err := os.Open(srcPath)
		if err != nil { return err }
		defer srcFile.Close()
		readers[i] = csv.NewReader(srcFile)
		readers[i].Comma = adaptor.delimiter
		readers[i].FieldsPerRecord = 1
		readerSignal[i] = make(chan int)
		defer close(readerSignal[i])
	}

	// start readers
	for i, _ := range readers {
		go func(ind int, reader *csv.Reader, readQueue chan<- rowVal, readSignal chan int) {
			rowv := rowVal{index:ind}
			for {
				_, ok := <- readSignal // wait for next read
				rowv.err = nil
				if !ok {
					return
				}
				// read values
				readVal, err := reader.Read()
				if err != nil {
					rowv.err = err
				}
				if len(readVal) != 1 {
					rowv.err = errors.New("More than one column in join file")
				}
				rowv.val = readVal[0]
				readQueue <- rowv // ship off read
			}
		}(i, readers[i], readerQueue, readerSignal[i])
	}
	
	// listen on reads
	row := make([]string, ncols)
	readRowSum := ncols*(ncols-1)/2
	rowIndex := 0
	eofIndex := -1
	totalEof := 0
	for {
		select {
		case rowv, _ := <- readerQueue:
			if rowv.err != nil {
				if rowv.err == io.EOF {
					// non uniform rows
					if(eofIndex >= 0 && eofIndex != rowIndex) {
						err = errors.New("Non-uniform rows on input files")
						return
					} else if (eofIndex < 0) { // eofIndex not ste
						eofIndex = rowIndex
					} else { // true eof
						if totalEof < ncols {
							break
						} else { // increment eof completion
							totalEof += 1
						}
					}
				} else { // reader error
					err = rowv.err
					return err
				}
			} else { // accumlate read value
				readRowSum -= rowv.index
				row[rowv.index] = rowv.val
			}
		default:
			// finshed reading row
			if readRowSum == 0 {
				readRowSum = ncols*(ncols-1)/2
				err = writer.Write(row)
				if err != nil {
					return
				}
				writer.Flush()
				for _, readSignal := range readerSignal {
					readSignal <- 1 // queue next read
				}
			}
		}
	}
	
	return nil
}
 
func (adaptor *DelimiterAdaptor) Shape(path string) (ncols, nrows int, err error) {
	// setup files and readers
	srcFile, err := os.Open(path)
	if err != nil { return }
	defer srcFile.Close()

	reader := csv.NewReader(srcFile)
	reader.Comma = adaptor.delimiter
	
	// stream through file
	nrows = 0  
	_, err = reader.Read()
	for err == nil {
		nrows += 1
		_, err = reader.Read()
	}	
	if  err != nil && err != io.EOF {
		return ncols, nrows, err
	}
	
	err = nil
	ncols = reader.FieldsPerRecord
	return
}

func translate(ncols int, srcPath, dstPath string, srcDelimiter, dstDelimiter rune) (err error) {
    // setup files
	srcFile, err := os.Open(srcPath)
	if err != nil { return err }
	defer srcFile.Close()
	dstFile, err := os.Open(dstPath)
	if err != nil { return err }
	defer dstFile.Close()
 
	// setup read writer
	reader := csv.NewReader(srcFile)
	reader.Comma = srcDelimiter
	reader.FieldsPerRecord = ncols
	writer := csv.NewWriter(dstFile)
	writer.Comma = dstDelimiter
	
	// transfer
	row, err :=  reader.Read()
	for err != nil {
		err := writer.Write(row)
		if err != nil {
			return err
		}
		writer.Flush()
		err = nil
	}
	
	if err != io.EOF {
		return
	}
	err = nil
	return 
}
 
func (adaptor *DelimiterAdaptor) ToRaw(srcPath, dstPath string) (error) {
	return translate(1, srcPath, dstPath, adaptor.delimiter, '\n')
}
 
func (adaptor *DelimiterAdaptor) FromRaw(srcPath, dstPath string) (err error) {
	return translate(1, srcPath, dstPath, '\n', adaptor.delimiter)
}
