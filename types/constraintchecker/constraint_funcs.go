package constraintchecker

import (
	"strconv"
	"errors"
	"github.com/ProtoML/ProtoML/types"
)

//type ValidateConstraintDefinition func(params ...string) error
type Constraint func(value string, params ...string) (valid bool, err error)

var (
	ConstrainFuncMap = map[string]Constraint {
		"(": ExclusiveLeftBound,
		")": ExclusiveRightBound,
		"[": InclusiveLeftBound,
		"]": InclusiveRightBound,
		"(]": ExclusiveInclusiveBound,
		"[)": InclusiveExclusiveBound,
		"()": ExclusiveBound,
		"[]": InclusiveBound,
		"=" : Member,
	}
	TypeConvMap = map[types.DataTypeName]func(string)(interface{}, error) {
		"bool":func (r string) (interface{}, error) {return strconv.ParseBool(r)},
		"int":func (r string) (interface{}, error) {return paramToInt(r)},
		"real":func (r string) (interface{}, error) {return paramToFloat(r)},
		"string":func (r string) (v interface{}, e error) {v = r; return}, // identity
	}
)

func paramToInt(param string) (int, error) {
	return strconv.Atoi(param)
}

func paramsToInt(params ...string) (nParams []int, err error) {
	nParams = make([]int,len(params))
	for i, p := range params {
		n, err := paramToInt(p)
		if err != nil {
			return nParams, err
		}
		nParams[i] = n
	}
	return
}

func paramToFloat(param string) (float64, error) {
	return strconv.ParseFloat(param, 64)
}

func paramsToFloat(params ...string) (fParams []float64, err error) {
	fParams = make([]float64,len(params))
	for i, p := range params {
		f, err := paramToFloat(p)
		if err != nil {
			return fParams, err
		}
		fParams[i] = f
	}
	return
}

// )
func ExclusiveRightBound(value string, params ...string) (valid bool, err error) {
	valid = false
	ParsedValue, err := paramToFloat(value)
	if err != nil {
		err := errors.New("Bad Value")
		return valid, err
	}
	against := params[0]
	ParsedAgainst, err := paramToFloat(against)
	if err != nil {
		err := errors.New("Bad Constraint")
		return valid, err
	}
	if ParsedValue < ParsedAgainst {
		err = nil
		valid = true
		return valid, err
	}
	return
}
// (
func ExclusiveLeftBound(value string, params ...string) (valid bool, err error) {
	valid = false
	ParsedValue, err := paramToFloat(value)
	if err != nil {
		err := errors.New("Bad Value")
		return valid, err
	}
	against := params[0]
	ParsedAgainst, err := paramToFloat(against)
	if err != nil {
		err := errors.New("Bad Constraint")
		return valid, err
	}
	if ParsedValue > ParsedAgainst {
		err = nil
		valid = true
		return valid, err
	}
	return
}
// [
func InclusiveLeftBound(value string, params ...string) (valid bool, err error) {
	valid = false
	ParsedValue, err := paramToFloat(value)
	if err != nil {
		err := errors.New("Bad Value")
		return valid, err
	}
	against := params[0]
	ParsedAgainst, err := paramToFloat(against)
	if err != nil {
		err := errors.New("Bad Constraint")
		return valid, err
	}
	if ParsedValue >= ParsedAgainst {
		err = nil
		valid = true
		return valid, err
	}
	return
}
// ]
func InclusiveRightBound(value string, params ...string) (valid bool, err error) {
	valid = false
	ParsedValue, err := paramToFloat(value)
	if err != nil {
		err := errors.New("Bad Value")
		return valid, err
	}
	against := params[0]
	ParsedAgainst, err := paramToFloat(against)
	if err != nil {
		err := errors.New("Bad Constraint")
		return valid, err
	}
	if ParsedValue <= ParsedAgainst {
		err = nil
		valid = true
		return valid, err
	}
	return
}

// ()
func ExclusiveBound(value string, params ...string) (valid bool, err error) {
	valid = false
	ParsedValue, err := paramToFloat(value)
	if err != nil {
		err := errors.New("Bad Value")
		return valid, err
	}
	againsts := params[0:2]
	ParsedAgainsts, err := paramsToFloat(againsts...)
	if err != nil {
		err := errors.New("Bad Constraint")
		return valid, err
	}
	if ParsedAgainsts[0] < ParsedValue && ParsedValue < ParsedAgainsts[1] {
		err = nil
		valid = true
		return valid, err
	}
	return
}

// []
func InclusiveBound(value string, params ...string) (valid bool, err error) {
	valid = false
	ParsedValue, err := paramToFloat(value)
	if err != nil {
		err := errors.New("Bad Value")
		return valid, err
	}
	againsts := params[0:2]
	ParsedAgainsts, err := paramsToFloat(againsts...)
	if err != nil {
		err := errors.New("Bad Constraint")
		return valid, err
	}
	if ParsedAgainsts[0] <= ParsedValue && ParsedValue <= ParsedAgainsts[1] {
		err = nil
		valid = true
		return valid, err
	}
	return
}

// [)
func InclusiveExclusiveBound(value string, params ...string) (valid bool, err error) {
	valid = false
	ParsedValue, err := paramToFloat(value)
	if err != nil {
		err := errors.New("Bad Value")
		return valid, err
	}
	againsts := params[0:2]
	ParsedAgainsts, err := paramsToFloat(againsts...)
	if err != nil {
		err := errors.New("Bad Constraint")
		return valid, err
	}
	if ParsedAgainsts[0] <= ParsedValue && ParsedValue < ParsedAgainsts[1] {
		err = nil
		valid = true
		return valid, err
	}
	return
}

// (]
func ExclusiveInclusiveBound(value string, params ...string) (valid bool, err error) {
	valid = false
	ParsedValue, err := paramToFloat(value)
	if err != nil {
		err := errors.New("Bad Value")
		return valid, err
	}
	againsts := params[0:2]
	ParsedAgainsts, err := paramsToFloat(againsts...)
	if err != nil {
		err := errors.New("Bad Constraint")
		return valid, err
	}
	if ParsedAgainsts[0] < ParsedValue && ParsedValue <= ParsedAgainsts[1] {
		err = nil
		valid = true
		return valid, err
	}
	return
}

func Member(value string, params ...string) (valid bool, err error) {
	err = nil
	valid = false
	for _, against := range params {
		if value == against {
			valid = true
			return valid, err
		}
	}
	return
}
