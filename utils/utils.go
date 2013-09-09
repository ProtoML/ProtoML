package utils

import (
	"fmt"
	"os"
)

func PanicOnError(err error, logMsg string) {
	if err != nil {
		panic(fmt.Sprintf("%v\nerror: %v", logMsg, err))
	}
}

func ErrorPanic(err error) {
	if err != nil {
		panic(fmt.Sprintf("\nerror: %v", err))
	}
}

func PrintAndExit(stuff ...interface{}) {
	fmt.Println(stuff...)
	os.Exit(1)
}

func HandleError(err error) {
	if err != nil {
		PrintAndExit(err)
	}
}

func Assert(ok bool, reason ...interface{}) {
	if !ok {
		PrintAndExit(reason...)
	}
}

func ToFloat64(i interface{}) float64 {
	o, ok := i.(float64)
	Assert(ok, "Cannot convert", i, "to float64.")
	return o
}

func ToString(i interface{}) string {
	o, ok := i.(string)
	Assert(ok, "Cannot convert", i, "to string.")
	return o
}

func ToStringArray(i []interface{}) []string {
	string_array := make([]string, len(i))
	for idx, value := range i {
		string_array[idx] = ToString(value)
	}
	return string_array
}
