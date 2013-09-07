package main

import (
	"fmt"
	// "github.com/ProtoML/ProtoML-core/parsers"
	// "github.com/ProtoML/ProtoML-core/interfaces"
	"github.com/ProtoML/ProtoML-core/parsers"
	// "github.com/ProtoML/ProtoML-core/types"
)

var Z = 4

func main() {
	// fmt.Println(types.MainDataTypes)
	// fmt.Println(types.MainDataTypes["binary"])
	// fmt.Println(types.EnumDataTypes)
	// fmt.Println(types.EnumDataTypes["file format"])
	fmt.Println(parsers.TransformParser("test"))
	// panic("doo")
}
