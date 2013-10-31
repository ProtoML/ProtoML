package main

import (
	"fmt"
	//"io/ioutil"
	//"net/http"
	"os"
	"github.com/ProtoML/ProtoML/api"
)

func main() {
	var args = os.Args

	if len(args) < 2 {
		fmt.Printf("Error: Must define a request\n")
		os.Exit(0)
	}

	if args[1] == api.APIDatatype {
		fmt.Printf("Attempting GET on all datatypes\n")
	} else if args[1] == api.APITransform {
		fmt.Printf("Attempting GET on all transforms\n")
	} else if args[1] == api.APITransformAddChild {
		if len(args) < 3 {
			fmt.Printf("Must specify a JSON file as input when calling /transform/add/child\n")
			os.Exit(0)
		}

		fmt.Printf("Attempting POST on /transform/add/child with JSON from file %s\n", args[2])
	} else {
		if len(args[1]) >= api.APIDatatypePrefixLen && args[1][:api.APIDatatypePrefixLen] == api.APIDatatypePrefix {
			fmt.Printf("Attempting GET on datatype with id %s\n", args[1][api.APIDatatypePrefixLen:])
		} else if len(args[1]) >= api.APITransformDeletePrefixLen && args[1][:api.APITransformDeletePrefixLen] == api.APITransformDeletePrefix {
			fmt.Printf("Attempting DELETE on transform with id %s\n", args[1][api.APITransformDeletePrefixLen:])
		} else if len(args[1]) >= api.APITransformUpdatePrefixLen && args[1][:api.APITransformUpdatePrefixLen] == api.APITransformUpdatePrefix {
			fmt.Printf("Attempting UPDATE on transform with id %s\n", args[1][api.APITransformUpdatePrefixLen:])
		} else if len(args[1]) >= api.APITransformPrefixLen && args[1][:api.APITransformPrefixLen] == api.APITransformPrefix {
			fmt.Printf("Attempting GET on transform with id %s\n", args[1][api.APITransformPrefixLen:])
		} else {
			fmt.Printf("Command currently unsupported\n")
		}
	}

	/*resp, err := http.Get("http://127.0.0.1:8080/transform")

	if err != nil {
		fmt.Printf("Error!\n")
	} else {
		fmt.Printf("No error!\n")
	}

	output, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("%s\n", output)*/
}
