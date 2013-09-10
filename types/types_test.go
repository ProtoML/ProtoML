package types

import (
	"testing"
)

func TestTypeName(t *testing.T) {
	for typeName, dataType := range DataTypes {
		if typeName != dataType.TypeName {
			t.Errorf("TypeName (%s) not matching entry in DataTypes (%s)", typeName, dataType.TypeName)
		}
	}
}

func TestFileFormat(t *testing.T) {
	for formatName, fileFormat := range FileFormats {
		if formatName != fileFormat.FormatName {
			t.Errorf("formatName (%s) not matching entry in FileFormats (%s)", formatName, fileFormat.FormatName)
		}
	}
}
