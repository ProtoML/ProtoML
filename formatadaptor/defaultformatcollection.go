package formatadaptor

import (
	"github.com/ProtoML/ProtoML/formatadaptor/csvadaptor"
	"github.com/ProtoML/ProtoML/formatadaptor/tsvadaptor"
)

func DefaultFileFormatCollection() (fc *FileFormatCollection) {
	fc = &FileFormatCollection{}
	fc.adaptors = make(map[string]FileFormatAdaptor)
	fc.RegisterAdaptor("csv",csvadaptor.New())
	fc.RegisterAdaptor("tsv",tsvadaptor.New())
	return
}
