package types

type DataTypeName string

type DataType struct {
	TypeName    DataTypeName
	ParentTypes []DataTypeName
	Description string
}

type TransformParameter struct {
	Default string // the default value of parameter
	Description string // description of the parameter
}

type TransformHyperParameter struct {
	Default string // the default value of parameter
	Type        []DataTypeName   // will check if it is of one of the types
	Constraints []ConstraintSexp // list of constraints
	Description string // description of the parameter
}


type ConstraintSexp []string

type FileParameter struct {
	Type      []DataTypeName
	Format			[]string // as long as the format of the file is in here it's good
	Description string
	Optional string
}

type StateParameter struct {
	Format []string
	Description string
	Optional string
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
	Name string //transform name
	Template string //transform original template
	// parameters
	PrimaryParameters      map[string]TransformParameter
	PrimaryHyperParameters map[string]TransformHyperParameter
	// command to run
	PrimaryExec string
	// help text
	Documentation string
	// functions
	Functions map[string]TransformFunction
	// data/state format created and accepted
	PrimaryInputs map[string]FileParameter
	PrimaryOutputs map[string]FileParameter
	PrimaryInputStates map[string]StateParameter
	PrimaryOutputStates map[string]StateParameter
}

type ElasticID string

type State struct {
	Source ElasticID
}

type InducedParameter string

type InducedDataGroup struct {
	Id ElasticID
	SelectedView []int
}

type InducedFileParameter struct {
	Data []DataGroup
	Path string
	Format string
}

type InducedStateParameter struct {
	Format string
	Path string
}

type InducedHyperParameter struct {
	Type DataTypeName
	Value string
}

type InducedTransform struct {
	Name            string
	TemplateID      ElasticID
	Function				string
	Parameters      map[string]InducedParameter // inserted valid parameters. Parameters are unchecked strings
	HyperParameters map[string]InducedHyperParameter // inserted valid hyperparameters
	// input definitions
	InputsIDs       map[string][]InducedDataGroup
	// output definitions
	OutputsIDs      map[string][]ElasticID
	// state definitions
	InputStatesIDs   	map[string]ElasticID
	OutputStatesIDs		map[string]ElasticID

	// denotes valid transform
	Error string

	// runtime members
	Id       ElasticID
	Template        string
	Exec            string
	Inputs          map[string]InducedFileParameter
	Outputs         map[string]InducedFileParameter
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

type ProtoMLVertex [2]string
func NewProtoMLVertex(typ string, id ElasticID) (v ProtoMLVertex) {
	v[0] = typ
	v[1] = string(id)
	return
}
type ProtoMLEdge [2]ProtoMLVertex
func NewProtoMLEdge(typeSource string, idSource ElasticID, typeSink string, idSink ElasticID) (e ProtoMLEdge) {
	e[0] = NewProtoMLVertex(typeSource, idSource)
	e[1] = NewProtoMLVertex(typeSink, idSink)
	return
}
type ProtoMLGraph struct {
	Vertices []ProtoMLVertex
	Edges []ProtoMLEdge
}
