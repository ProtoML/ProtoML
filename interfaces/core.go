package interfaces

type EnumWithDefault struct {
	DefaultValue string
	Values       []string
}

type Data struct {
	ParentId  string            // unique id of parent
	DataIndex uint              // output index of the data
	Main      string            // main data type
	BoolTypes map[string]bool   // boolean data types
	IntTypes  map[string]int64  // integer data types
	EnumTypes map[string]string // enum data types
}

type TransformDefinition struct {
	Template         string                 // template to copy
	DeleteParameters bool                   // whether or not to keep template parameters
	Parameters       map[string]interface{} // transform parameters
	Types            map[string]interface{} // transform types
	Input            [][]string             // input definitions
	Output           [][]string             // output definitions
}

type DataConstraint func(Data)

type ParameterConstraint func(interface{})

type Transform struct {
	TransformName        string                         // filename to call
	ParameterConstraints map[string]ParameterConstraint // parameter constraints
	InputConstraints     []DataConstraint               // constraints for inputs
	OutputDefinition     []DataConstraint               // definition for outputs
}

type Parameters map[string]interface{}

type Transformer interface {
	train(id string, parameters Parameters, data ...Data)
	transform(id string, parameters Parameters, data ...Data)
}
