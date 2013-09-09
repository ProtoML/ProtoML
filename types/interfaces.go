package types

type DataTypeName string

type DataType struct {
	TypeName    DataTypeName
	ParentTypes []DataTypeName
	Description string
	Validator   func(Data) bool // function to test if data matches data type
}

type FileFormat struct {
	FormatName string
	Extension  string
	Validator  func(Data) bool
}

type Tag struct {
	UniqueName  string
	Description string
}

type TransformParameter struct {
	ParameterName string   // name of the parameter
	Values        []string // only filed in for enum parameters
	Distribution  string   // only filled in for numerical parameters
	NoConstraint  bool     // only filled in for arbitrary string parameters
	Description   string   // description of the parameter
}

type Transform struct {
	// a transform to copy
	Template string
	// help text
	Documentation string
	// tags to find and categorize the transform
	Tags []Tag
	// whether or not to keep the template's parameters
	OverwriteParameters bool
	// transform parameters
	Parameters []TransformParameter
	// input definitions
	Input [][]string
	// output definitions
	Output [][]string
}

type Data struct {
	ParentId       string
	ExclusiveType  DataTypeName
	Sparse         bool
	Missing        bool
	NRows          uint
	NCols          uint
	FileFormat     FileFormat
	AddtionalTypes []DataTypeName
}

type RunRequest struct {
	DataNamespace  string // namespace of the data
	TransformName  string // name of transform
	JsonParameters string // filename of json containing parameters
	Data           []Data // input data
}
