package types

type EnumWithDefault struct {
	DefaultValue string
	Values       []string
}

type DataTypeName string

type DataType struct {
	TypeName DataTypeName
	ParentType DataTypeName
	Exclusive bool
	Description string
	Validator func(DataType) error //stub
}

type TransformParameter struct {
	Parameter string
	Default string
	Description string
}

type TransformParameterValue struct {
	ParameterType TransformParameter
	Value interface{}
}

type TransformTemplate string

type Transform struct {
	Template TransformTemplate
	ExclusiveInputTypes []DataTypeName
	AdditionalInputTypes []DataTypeName
	ExclusiveOutputTypes []DataTypeName
	AdditionalOutputTypes []DataTypeName
	Executor string
	ExecutorFlags []TransformParameter
	Parameters map[string]interface{}
}

type Data struct {
	ParentId string
	ExclusiveType DataTypeName
	AddtionalTypes []DataTypeName
}

type PipelineNode struct {
	Template TransformTemplate
	Id string
	ParentId string
}

type Pipeline struct {
	Nodes []PipelineNode
}
