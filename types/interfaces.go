package types

type DataTypeName string

type DataType struct {
	TypeName    DataTypeName
	ParentTypes []DataTypeName
	Description string
}

type TransformParameter []string

type TransformHyperParameter struct {
	Default string // the default value of parameter
	Type        []DataTypeName   // will check if it is of one of the types
	Constraints []ConstraintSexp // list of constraints
	//Description string // description of the parameter
}


type ConstraintSexp []string
/*
type DataConstraint struct {
	ExclusiveType DataTypeName
	Restriction   []ConstraintSexp // String restriction as specified, to be interpreted by type checker
}
*/

type FileParameter struct {
	Path				string
	Format			[]string // as long as the format of the file is in here it's good
	Description string
}

type StateParameter struct {
	Path string
	Format []string
	Description string
}

type TransformFunction struct {
	// function name, description
	Description string
	// transform parameters
	Parameters      map[string]TransformParameter
	HyperParameters map[string]TransformHyperParameter
	// script to run
	Exec string
	// input definitions
	Inputs map[string]FileParameter
	// output definitions
	Outputs map[string]FileParameter
	InputStates map[string]StateParameter
	OutputStates map[string]StateParameter
}

type Transform struct {
	PrimaryParameters      map[string]TransformParameter
	PrimaryHyperParameters map[string]TransformHyperParameter
	// a transform to copy
	Template string
	// script to run
	PrimaryExec string
	// help text
	Documentation string
	// functions
	Functions map[string]TransformFunction
	// state formate created and accepted
	PrimaryInputs map[string]FileParameter
	PrimaryOutputs map[string]FileParameter
	PrimaryInputStates map[string]StateParameter
	PrimaryOutputStates map[string]StateParameter
}

type InducedParameter string

type InducedFileParameter struct {
	Data DataGroup
	Path string
	Format string
}

type InducedStateParameter struct {
	Format string
	Path string
}

type InducedHyperParameter struct {
	PrimitiveType string
	Value string
}

type InducedTransform struct {
	Template        string
	Exec						string
	Function				string
	Parameters      map[string]InducedParameter // inserted valid parameters. Parameters are unchecked strings
	HyperParameters map[string]InducedHyperParameter // inserted valid hyperparameters
	Inputs          map[string]InducedFileParameter // input definitions
	Outputs         map[string]InducedFileParameter // output definitions
	InputStates			map[string]InducedStateParameter
	OutputStates		map[string]InducedStateParameter
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
	NRows      int
	NCols      int
	Columns    DatasetColumns
}

type DataGroupColumns struct {
	ExclusiveType DataTypeName
	Tags          [][]string
}

type DataGroup struct {
	FileFormat string
	NRows      int
	NCols      int
	Columns    DataGroupColumns
	Source     string
}
