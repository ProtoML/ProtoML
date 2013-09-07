package parsers

/*
   Parser for the json defining a transform.

   Naming convention is transform_name + ".json".
*/

import (
	"github.com/ProtoML/ProtoML-core/interfaces"
	"github.com/ProtoML/ProtoML-core/types"
	"github.com/ProtoML/ProtoML-core/utils"
	"github.com/ProtoML/ProtoML-persist/persist"
	"regexp"
	"strconv"
)

type transformParser func(interfaces.TransformDefinition, *interfaces.Transform)

var (
	TRANSFORMPATH      string
	PARSERS            []transformParser
	NORMAL_DIST_REGEX  = regexp.MustCompile(`^N\(\d+(\.\d+)?,\d+(\.\d+)?\)$`)
	UNIFORM_DIST_REGEX = regexp.MustCompile(`^[\[\(]\d+(\.\d+)?,\d+(\.\d+)?[\]\)]$`)
	FLOAT_REGEX        = regexp.MustCompile(`\d+(\.\d+)?`)
)

func init() {
	TRANSFORMPATH = persist.StringConfig("TRANSFORMPATH")
	PARSERS = []transformParser{
		templateParser,
		deleteParametersParser,
		parametersParser,
		typeValidator,
		inputParser,
		outputParser,
	}
}

func TransformParser(transformPath string) interfaces.Transform {
	transform := emptyTransform(transformPath)
	if transformPath != "" {
		transform_json := transformJsonReader(transformPath)
		for _, parser := range PARSERS {
			parser(transform_json, &transform)
		}
	}
	return transform
}

func transformJsonReader(transformPath string) interfaces.TransformDefinition {
	var transform_json interfaces.TransformDefinition
	utils.HandleError(persist.JsonDecoder(TRANSFORMPATH, transformPath).Decode(&transform_json))
	return transform_json
}

func emptyTransform(name string) (t interfaces.Transform) {
	t.TransformName = name
	t.ParameterConstraints = make(map[string]interfaces.ParameterConstraint)
	return
}

func sampleParser(td interfaces.TransformDefinition, t *interfaces.Transform) {
}

func templateParser(td interfaces.TransformDefinition, t *interfaces.Transform) {
	template := TransformParser(td.Template)
	t.ParameterConstraints = template.ParameterConstraints
	t.InputConstraints = template.InputConstraints
	t.OutputDefinition = template.OutputDefinition
}

func deleteParametersParser(td interfaces.TransformDefinition, t *interfaces.Transform) {
	if td.DeleteParameters {
		t.ParameterConstraints = make(map[string]interfaces.ParameterConstraint)
	}
}

func stringParameterConstraintParser(constraint string) interfaces.ParameterConstraint {
	var parameterConstraints []interfaces.ParameterConstraint

	if constraint[0] == 'q' {
		constraint = constraint[1:]
		parameterConstraints = append(parameterConstraints,
			func(value interface{}) {
				switch value.(type) {
				case int:
				default:
					utils.PrintAndExit("Quantized parameter must be integer.")
				}
			},
		)
	}

	if constraint[:3] == "log" {
		constraint = constraint[3:]
		parameterConstraints = append(parameterConstraints,
			func(value interface{}) {
				utils.Assert(utils.ToFloat64(value) > 0, "Log parameter must be > 0.")
			},
		)
	}

	if NORMAL_DIST_REGEX.MatchString(constraint) {
		// no constraint if normal
	} else if UNIFORM_DIST_REGEX.MatchString(constraint) {
		bounds := FLOAT_REGEX.FindAllString(constraint, 2)
		lower_bound, err_lower_bound := strconv.ParseFloat(bounds[0], 64)
		utils.HandleError(err_lower_bound)
		upper_bound, err_upper_bound := strconv.ParseFloat(bounds[1], 64)
		utils.HandleError(err_upper_bound)

		utils.Assert(lower_bound <= upper_bound, "Lower bound greater than Upper bound.")

		if constraint[0] == '[' {
			parameterConstraints = append(parameterConstraints,
				func(value interface{}) {
					utils.Assert(utils.ToFloat64(value) >= lower_bound, "Parameter below lower bound")
				},
			)
		} else {
			parameterConstraints = append(parameterConstraints,
				func(value interface{}) {
					utils.Assert(utils.ToFloat64(value) > lower_bound, "Parameter below lower bound")
				},
			)
		}

		if constraint[len(constraint)-1] == ']' {
			parameterConstraints = append(parameterConstraints,
				func(value interface{}) {
					utils.Assert(utils.ToFloat64(value) <= upper_bound, "Parameter above upper bound")
				},
			)
		} else {
			parameterConstraints = append(parameterConstraints,
				func(value interface{}) {
					utils.Assert(utils.ToFloat64(value) < upper_bound, "Parameter above upper bound")
				},
			)
		}
	} else {
		utils.PrintAndExit("Invalid Parameter Constraint:", constraint)
	}

	return func(value interface{}) {
		for _, pc := range parameterConstraints {
			pc(value)
		}
	}

}

func listParameterConstraintParser(constraint []string) interfaces.ParameterConstraint {
	return func(value interface{}) {
		string_value := utils.ToString(value)
		for _, enum_value := range constraint {
			if enum_value == string_value {
				return
			}
		}
		utils.PrintAndExit("Value not part of enumeration parameter: " + string_value)
	}
}

func parametersParser(td interfaces.TransformDefinition, t *interfaces.Transform) {
	for key, v := range td.Parameters {
		switch v.(type) {
		case []interface{}:
			new_v := utils.ToStringArray(v.([]interface{}))
			t.ParameterConstraints[key] = listParameterConstraintParser(new_v)
		case string:
			new_v := v.(string)
			t.ParameterConstraints[key] = stringParameterConstraintParser(new_v)
		default:
			utils.PrintAndExit("Incorrect Parameter Constraint:", v)
		}
	}
}

func typeValidator(td interfaces.TransformDefinition, t *interfaces.Transform) {
	for key, value := range td.Types {
		switch value.(type) {
		case bool:
			if _, ok := types.BooleanTransformTypes[key]; !ok {
				utils.PrintAndExit("Invalid Boolean Transform Type: " + key)
			}
		case int:
			if _, ok := types.IntegerTransformTypes[key]; !ok {
				utils.PrintAndExit("Invalid Integer Transform Type: " + key)
			}
		case string:
			if _, ok := types.EnumTransformTypes[key]; !ok {
				utils.PrintAndExit("Invalid Enum Transform Type: " + key)
			}
		default:
			utils.PrintAndExit("Invalid Transform Type: " + key)
		}
	}
}

func inputParser(td interfaces.TransformDefinition, t *interfaces.Transform) {
	if len(td.Input) > 0 {
		t.InputConstraints = parseDataConstraints(td.Input)
	}
}

func outputParser(td interfaces.TransformDefinition, t *interfaces.Transform) {
	if len(td.Output) > 0 {
		t.OutputDefinition = parseDataConstraints(td.Output)
	}
}

func parseDataConstraints(strings_constraints [][]string) []interfaces.DataConstraint {
	dataConstraintsArray := make([]interfaces.DataConstraint, len(strings_constraints))
	for idx, strings_constraint := range strings_constraints {
		dataConstraintsArray[idx] = parseDataConstraint(strings_constraint)
	}
	return dataConstraintsArray
}

func parseDataConstraint(strings_constraint []string) interfaces.DataConstraint {
	var constraints []interfaces.DataConstraint
	for _, string_constraint := range strings_constraint {
		constraints = append(constraints, toDataConstraint(string_constraint))
	}
	return func(d interfaces.Data) {
		for _, constraint := range constraints {
			constraint(d)
		}
	}
}

func toDataConstraint(string_constraint string) interfaces.DataConstraint {
	length := len(string_constraint)
	for i := 1; i < length-1; i++ { // don't check first and last character
		if string_constraint[i] == '=' {
			return toValueDataConstraint(string_constraint[:i], string_constraint[i+1:])
		}
	}
	_, present := types.MainDataTypes[string_constraint]
	if present {
		return toMainDataConstraint(string_constraint)
	} else {
		return toValueDataConstraint(string_constraint, "true")
	}
}

func toValueDataConstraint(key, value string) interfaces.DataConstraint {
	defaultBool, isBool := types.BooleanDataTypes[key]
	if isBool {
		bool_val, err := strconv.ParseBool(value)
		utils.HandleError(err)
		return func(data interfaces.Data) {
			data_val, ok := data.BoolTypes[key]
			if !ok {
				data_val = defaultBool
			}
			utils.Assert(bool_val == data_val, "Data doesn't satisfy boolean constraint", key)
		}
	}
	defaultInt, isInteger := types.IntegerDataTypes[key]
	if isInteger {
		int_val, err := strconv.ParseInt(value, 10, 64)
		utils.HandleError(err)
		return func(data interfaces.Data) {
			data_val, ok := data.IntTypes[key]
			if !ok {
				data_val = defaultInt
			}
			utils.Assert(int_val == data_val, "Data doesn't satisfy integer constraint", key)
		}
	}
	defaultEnumData, isEnum := types.EnumDataTypes[key]
	if isEnum {
		defaultEnum := defaultEnumData.DefaultValue
		return func(data interfaces.Data) {
			data_val, ok := data.EnumTypes[key]
			if !ok {
				data_val = defaultEnum
			}
			utils.Assert(value == data_val, "Data doesn't satisfy enum constraint", key)
		}
	}
	utils.PrintAndExit("Data constraint name not found:", key)
	return func(data interfaces.Data) {}
}

func isMainSubtype(to_check, constraint string) bool {
	if to_check == constraint {
		return true
	}
	for _, parent := range types.MainDataTypes[to_check] {
		if isMainSubtype(parent, constraint) {
			return true
		}
	}
	return false
}

func toMainDataConstraint(string_constraint string) interfaces.DataConstraint {
	return func(data interfaces.Data) {
		utils.Assert(isMainSubtype(data.Main, string_constraint), "Data main type is not subtype of", string_constraint)
	}
}
