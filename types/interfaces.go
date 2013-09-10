package types

type DataTypeName string

type DataType struct {
	TypeName    DataTypeName
	ParentTypes []DataTypeName
	Description string
	Validator   func(Data) bool // function to test if data matches data type
}

type FileFormat struct {
	FormatName  string
	Description string
	Validator   func(Data) bool
}

type TransformParameter struct {
	ParameterName string   // name of the parameter
	Values        []string // only filed in for enum parameters
	Distribution  string   // only filled in for numerical parameters
	NoConstraint  bool     // only filled in for arbitrary string parameters
	Description   string   // description of the parameter
}

type DataConstraint struct {
	ExclusiveType DataTypeName
	Sparse        bool
	Missing       bool
	NCols         uint
	FileFormat    FileFormat
}

type Transform struct {
	// a transform to copy
	Template string
	// help text
	Documentation string
	// whether or not to keep the template's parameters
	OverwriteParameters bool
	// transform parameters
	Parameters []TransformParameter
	// input definitions
	Input []DataConstraint
	// output definitions
	Output []DataConstraint
}

type Data struct {
	DataId        string
	ExclusiveType DataTypeName
	Sparse        bool
	Missing       bool
	NRows         uint
	NCols         uint
	FileFormat    FileFormat
}

type RunRequest struct {
	DataNamespace  string // namespace of the data
	TransformName  string // name of transform
	JsonParameters string // filename of json containing parameters
	Data           []Data // input data
}
