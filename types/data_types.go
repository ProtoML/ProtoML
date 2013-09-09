package types

var DataTypes = map[DataTypeName]DataType{
	"log": {TypeName: "log",
		ParentTypes: []DataTypeName{},
		Description: "for displaying / book keeping",
		Validator:   func(d Data) bool { return true }},
	"scalar": {TypeName: "scalar",
		ParentTypes: []DataTypeName{"log"},
		Description: "single numbers such as metrics",
		Validator:   func(d Data) bool { return true }},
	"observation": {TypeName: "observation",
		ParentTypes: []DataTypeName{},
		Description: "row based data",
		Validator:   func(d Data) bool { return true }},
	"number": {TypeName: "number",
		ParentTypes: []DataTypeName{"observation"},
		Description: "row based number data",
		Validator:   func(d Data) bool { return true }},
}

/*
   "string":      {"observation"},            // any string data
   "text":        {"string"},                 // long strings with words
   "date-string": {"string"},                 // date data
   "word":        {"string"},                 // single words
   "categorical": {"number"},                 // unordered data
   "numerical":   {"number"},                 // ordered data
   "0,1":         {"categorical", "ordinal"}, // 0, 1 binary data
   "-1,1":        {"categorical", "ordinal"}, // -1, +1 binary data
   "ordinal":     {"numerical"},              // ordered data with meaningless fractions
*/
