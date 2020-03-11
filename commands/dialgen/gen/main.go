package main

import (
	"fmt"
	"os"

	"github.com/team-rocos/gomavlib/commands/dialgen/libgen"
)

func main() {
	//kingpin.CommandLine.Help = "Convert Mavlink dialects from XML format into Go format."

	//preamble := kingpin.Flag("preamble", "preamble comment").String()
	//mainDefAddr := kingpin.Arg("xml", "a path or url pointing to a XML Mavlink dialect").Required().String()

	// If common.xml is in a different file location to the main xml file specified, and it's included in the main xml file,
	// then the location to common.xml should be specified as a command line argument. If common.xml is included, and is in
	// the same location as the main xml file specified, then this additional argument is optional.
	//includes := kingpin.Arg("common.xml", "a path or url pointing to the common XML Mavlink dialect - common.xml").Required().String()

	//kingpin.Parse()

	// Specifying include directories as command line arguments is only necessary if the main xml file includes other xml files.
	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough input command line arguments showing xml file locations")
		os.Exit(1)
	}

	var includeDirectories []string

	// Note: First value in os.Args (os.Args[0]) shows the mode of running the code (e.g. debug) and is not used as an argument for this code's functionality
	mainDefAddr := os.Args[1]
	if len(os.Args) >= 3 {
		includeDirectories = os.Args[2:]
	}

	preamble := ""
	err := libgen.GenerateGoCode(preamble, mainDefAddr, includeDirectories)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %s\n", err)
		os.Exit(1)
	}
}
