package types

/*
   Types that apply to data.

   Removing a data type should be as simple as commenting it out in this file.

   Type names should not contain an equal sign and should be unique.
*/

import (
	"github.com/ProtoML/ProtoML-core/interfaces"
)

type parents []string

var MainDataTypes = map[string]parents{
	"log":         {},                         // for displaying / book keeping
	"scalar":      {"log"},                    // single numbers such as metrics
	"observation": {},                         // row based data
	"number":      {"observation"},            // row based number data
	"string":      {"observation"},            // any string data
	"text":        {"string"},                 // long strings with words
	"date-string": {"string"},                 // date data
	"word":        {"string"},                 // single words
	"categorical": {"number"},                 // unordered data
	"numerical":   {"number"},                 // ordered data
	"0,1":         {"categorical", "ordinal"}, // 0, 1 binary data
	"-1,1":        {"categorical", "ordinal"}, // -1, +1 binary data
	"ordinal":     {"numerical"},              // ordered data with meaningless fractions
}

var BooleanDataTypes = map[string]bool{
	"sparse":      false, // having relatively few non-zero entries
	"prediction":  false, // output of a predictor
	"missing":     false, // contains missing values
	"positive":    false, // greater than 0
	"nonnegative": false, // greater than or equal to 0
	"scaled":      false, // 0 mean, 1 standard deviation, or close enough (binary)
}

var IntegerDataTypes = map[string]int64{
	"depth":       0, // the number of transforms passed through
	"state depth": 0, // the number of stateful transforms passed through
	"nrows":       1, // the number of observations in the data
	"ncols":       1, // the number of features in the data
}

var EnumDataTypes = map[string]interfaces.EnumWithDefault{
	"file format": {
		"txt", // default is text data
		[]string{
			"txt",    // text files
			"vw",     // vowpal wabbit format
			"csv",    // comma separated values
			"libsvm", // libsvm format
			"npy",    // numpy memmap (memory map)
			"h5",     // hierarchical data format 5
			"tsv",    // tab separated values
			"rda",    // R data
			"dat",    // matlab data
			"arff",   // weka format
		},
	},
	"precision": {
		"float32", // default is 32-bit float
		[]string{
			"float32", // 32-bit float
			"float64", // 64-bit float
			"int32",   // 32-bit integer
			"int64",   // 64-bit integer
			"int8",    // 8-bit integer
			"uint8",   // 6-bit unsigned integer
		},
	},
}
