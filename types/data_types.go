package types

// list of default data types
var DefaultDataTypes = []DataType{
	DataType{	
		TypeName: "string",
		ParentTypes: []DataTypeName{},
		Description: "any string data"},
	DataType{	
		TypeName: "YYYY-MM-DD",
		ParentTypes: []DataTypeName{"string"},
		Description: "date data with appropriate format"},
	DataType{
		TypeName: "word",
		ParentTypes: []DataTypeName{"string"},
		Description: "single word data"},
	DataType{	
		TypeName: "text",
		ParentTypes: []DataTypeName{"string"},
		Description: "long string with words"},
	DataType{
		TypeName: "ordinal",
		ParentTypes: []DataTypeName{},
		Description: "data that has an order"},
	DataType{
		TypeName: "number",
		ParentTypes: []DataTypeName{},
		Description: "numerical value"},
	DataType{	
		TypeName: "categorical",
		ParentTypes: []DataTypeName{"number"},
		Description: "unordered data corresponding to categories"},
	DataType{
		TypeName: "int",
		ParentTypes: []DataTypeName{"number","ordinal"},
		Description: "numbered data with a meaningful ordering"},
	DataType{
		TypeName: "real",
		ParentTypes: []DataTypeName{"number","ordinal"},
		Description: "numbered data with meaningful fractional values"},
	DataType{
		TypeName: "0,1",
		ParentTypes: []DataTypeName{"categorical", "int"},
		Description: "0, 1 binary data"},
	DataType{	
		TypeName: "-1,1",
		ParentTypes: []DataTypeName{"categorical", "int"},
		Description: "row based number data"},
}

