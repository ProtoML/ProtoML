package types

type DataTypeName string

type DataType struct {
	TypeName    DataTypeName      
	ParentTypes []DataTypeName
	Description string
}

type TransformParameter struct {
	ParameterName string   // name of the parameter
	Values        []string // only filed in for enum parameters
	DefaultValue  string   // the default value of parameter
	Distribution  string   // only filled in for numerical parameters
	NoConstraint  bool     // only filled in for arbitrary string parameters
	Description   string   // description of the parameter
}

type DataConstraint struct {
	ExclusiveType DataTypeName
	Tags          []string
}

type TransformState struct {
	StateExt string
}

type TransformFunction struct {
	// function name
	Name string
	// transform parameters
	Parameters []TransformParameter
	// input definitions
	Input []DataConstraint
	// output definitions
	Output []DataConstraint
}

type Transform struct {
	// a transform to copy
	Template string
	// help text
	Documentation string
	// functions
	Functions []TransformFunction
	// formats acceptable
	FileFormats []string
	// state formate created and accepted
	State TransformState
}

type InducedTransform struct {
	Template	 string
	Parameters	 map[string]string // partially applied valid parameters
	Input		 map[string]DataGroup  // input definitions
	Output		 map[string]DataGroup  // output definitions
	State        string //optional state file for transform
}

type DataColumnTypeGroup map[DataTypeName][]int
type DataColumnTagGroup map[string][]int

type DatasetColumns struct {
	ExclusiveTypes DataColumnTypeGroup
	Tags           DataColumnTagGroup
}

type DatasetFile struct {
	Path       string
	FileFormat string
	NRows      uint
	NCols      uint
	Columns    DatasetColumns
}

type DataGroupColumns struct {
	ExclusiveType DataTypeName
	Tags          [][]string
}

type DataGroup struct {
	FileFormat string
	NRows      uint
	NCols      uint
	Columns    DataGroupColumns
	Source     string
}
