package constraintchecker

import (
	"errors"
	"github.com/ProtoML/ProtoML/types"
	"github.com/ProtoML/ProtoML-persist/persist/elastic"
	"fmt"
)

func CheckParam(ind map[string]types.InducedParameter, primary, function map[string]types.TransformParameter, param types.TransformParameter, val types.InducedParameter) (err error) {
	// Params can't be bad right now, so this does nothing
	return
}

func CheckHyper(ind map[string]types.InducedHyperParameter, primary, function map[string]types.TransformHyperParameter, param types.TransformHyperParameter, val types.InducedHyperParameter) (err error) {
	// Complex
	// Essentially, go through each constraint and run the function
	for name, hparam := range ind {
		indValue := hparam.Value
		param, ok := function[name]
		if !ok {
			param, ok = primary[name]
			if !ok {
				err = errors.New(fmt.Sprintf("Could not find HyperParameter %s",name))
				return
			}
		}
		constraintlist := param.Constraints
		// Now go through the constraint list and check that hparam['Value'] satisfies the function
		valid := false
		err = errors.New(fmt.Sprintf("Failed constraint for hyperparameter %s", name))
		for _, constraint := range constraintlist {
			constrSymbol := constraint[0]
			constrParams := constraint[1:]
			c, ok := ConstrainFuncMap[constrSymbol]
			if !ok {
				err = errors.New(fmt.Sprintf("Not a valid function symbol: %s",constrSymbol))
				return
			}
			valid, err = c(indValue, constrParams...)
			if err != nil {
				err = errors.New(fmt.Sprintf("Invalid constraints for %s: %v", constrSymbol, constrParams))
				return
			}
			if valid {
				err = nil
				break
			}
		}
		if len(constraintlist) == 0 {
			valid = true
			err = nil
		}
	}
	return
}
func CheckFile(ind map[string]types.InducedFileParameter, primary, function map[string]types.FileParameter, param types.FileParameter, val types.InducedFileParameter) (err error) {
	// Get list of strings for Type/Format
	for fname, fparam := range ind {
		fparamt, ok := function[fname]
		if !ok {
			fparamt, ok = primary[fname]
			if !ok {
				err = errors.New(fmt.Sprintf("No file by name of %s in template", fname))
				return
			}
		} else if len(fparamt.Format) == 0 {
			continue
		}
		for i, group := range fparam.Data {
			InducedType := group.Columns.ExclusiveType
			valid := false
			for _, TemplateType := range fparamt.ValidTypes {
				ancestor, _ := elastic.IsDataTypeAncestor(InducedType, TemplateType)
				if ancestor || InducedType == TemplateType {
					valid = true
					break
				}
			}
			if !valid {
				err = errors.New(fmt.Sprintf("Could not find matching type for %s in column group %d",InducedType,i))
				return
			}
		}
		membership, _ := Member(fparam.Format, fparamt.Format...)
		if !membership {
			err = errors.New(fmt.Sprintf("Invalid format %s for %s",fparam.Format,fname))
			return
		}
	}
	return
}
func CheckState(ind map[string]types.InducedStateParameter, primary, function map[string]types.StateParameter, param types.StateParameter, val types.InducedStateParameter) (err error) {
	// Get list of strings for Format
	for sname, sparam := range ind {
		sparamt, ok := function[sname]
		if !ok {
			sparamt, ok = primary[sname]
			if !ok {
				err = errors.New(fmt.Sprintf("No state by name of %s in template", sname))
				return
			}
		} else if len(sparamt.Format) == 0 {
			continue
		}
		membership, _ := Member(sparam.Format, sparamt.Format...)
		if !membership {
			err = errors.New(fmt.Sprintf("Invalid format %s for %s",sparam.Format, sname))
			return
		}
	}
	return
}
