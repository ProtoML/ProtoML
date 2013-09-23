package main

import (
	"fmt"
	//"io/ioutil"
	//"net/http"
	"os"
)

const datatype                        = "/datatype"
const transform                       = "/transform"
const transformAddRoot                = "/transform/add/root"
const transformAddChild               = "/transform/add/child"

const datatypePrefix                  = "/datatype/"
const transformPrefix                 = "/transform/"
const transformDeletePrefix           = "/transform/delete/"
const transformUpdatePrefix           = "/transform/update/"

const datatypePrefixLen               = len(datatypePrefix)
const transformPrefixLen              = len(transformPrefix)
const transformDeletePrefixLen        = len(transformDeletePrefix)
const transformUpdatePrefixLen        = len(transformUpdatePrefix)

func main() {
	var args = os.Args

	if len(args) < 2 {
		fmt.Printf("Error: Must define a request\n")
		os.Exit(0)
	}

	if args[1] == datatype {
		fmt.Printf("Attempting GET on all datatypes\n")
	} else if args[1] == transform {
		fmt.Printf("Attempting GET on all transforms\n")
	} else if args[1] == transformAddRoot {
		if len(args) < 3 {
			fmt.Printf("Must specify a JSON file as input when calling /transform/add/root\n")
			os.Exit(0)
		}

		fmt.Printf("Attempting POST on /transform/add/root with JSON from file %s\n", args[2])
	} else if args[1] == transformAddChild {
		if len(args) < 3 {
			fmt.Printf("Must specify a JSON file as input when calling /transform/add/child\n")
			os.Exit(0)
		}

		fmt.Printf("Attempting POST on /transform/add/child with JSON from file %s\n", args[2])
	} else {
		if len(args[1]) >= datatypePrefixLen && args[1][:datatypePrefixLen] == datatypePrefix {
			fmt.Printf("Attempting GET on datatype with id %s\n", args[1][datatypePrefixLen:])
		} else if len(args[1]) >= transformDeletePrefixLen && args[1][:transformDeletePrefixLen] == transformDeletePrefix {
			fmt.Printf("Attempting DELETE on transform with id %s\n", args[1][transformDeletePrefixLen:])
		} else if len(args[1]) >= transformUpdatePrefixLen && args[1][:transformUpdatePrefixLen] == transformUpdatePrefix {
			fmt.Printf("Attempting UPDATE on transform with id %s\n", args[1][transformUpdatePrefixLen:])
		} else if len(args[1]) >= transformPrefixLen && args[1][:transformPrefixLen] == transformPrefix {
			fmt.Printf("Attempting GET on transform with id %s\n", args[1][transformPrefixLen:])
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
