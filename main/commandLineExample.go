package main

import (
	"flag"
	"fmt"
	"os"
)

//This is how you declare a global variable
var someOption bool

//This is how you declare a global constant
const usageMsg string = "goprog [-someoOtion] args\n"

func main() {
	fmt.Println(usageMsg)
	flag.BoolVar(&someOption, "someOption", false, "Run with someOption")
	//Setting Usage will cause usage to be executed if options are provided
	//that were never defined, e.g. "goprog -someOption -foo"
	flag.Usage = usage
	flag.Parse()
	if someOption {
		fmt.Printf("someOption was set\n")
	}
	//If there are other required command line arguments, that are not
	//options, they will now be available to parse manually.  flag does
	//not do this part for you.
	for _, v := range flag.Args() {
		fmt.Printf("%+v\n", v)
	}

	//Calling this program as "./goprog -someOption dog cat goldfish"
	//outputs
	//someOption was set
	//dog
	//cat
	//goldfish
}

func usage() {
	fmt.Printf(usageMsg)
	flag.PrintDefaults()
	os.Exit(1)
}