package types

/*
   Types that apply to transforms, along with their default values.

   These are not necessary for normal use, but may be useful for data collection, higher level learning, more intelligent scheduling, and more.

   Removing a data type should be as simple as commenting it out in this file.
*/

import (
	"github.com/ProtoML/ProtoML-core/interfaces"
)

var BooleanTransformTypes = map[string]bool{
	"predictor":                false, // transform that predicts a target variable
	"boosted":                  false, // transform that already has been boosted
	"linear":                   false, // transform based on linear models
	"tree":                     false, // transform based on trees
	"distance":                 false, // transform based on distance
	"rbf":                      false, // transform based on radial basis function
	"cluster":                  false, // transform that clusters
	"feature selection":        false, // transform that eliminates features
	"feature expansion":        false, // transform that creates more features
	"monotonic":                false, // transform that doesn't change the ordering of the data (e.g. useless for tree transforms)
	"constant ram":             false, // transform that takes a constant amount of memory
	"stateless":                false, // transform that doesn't require validation to determine generalization performance
	"probabilistic":            false, // transform based on probabilistic methods
	"missing handler":          false, // transform that can handle missing data
	"boostable":                false, // transform that can benefit from boosting
	"baggable":                 false, // transform that can benefit from bagging
	"repeatable":               false, // transform that can benefit from repeated application
	"dimensionality reduction": false, // transform to reduct number of columns
	"frequency":                false, // frequency based transform
	"cache":                    true,  // whether or not output should be cached
	"depth increase":           true,  // transform increases "depth" data type
}

var IntegerTransformTypes = map[string]int64{
	"cores": 1, // number of cores the transform requires
	"gpus":  0, // number of gpus the transform requires
}

var EnumTransformTypes = map[string]interfaces.EnumWithDefault{}
