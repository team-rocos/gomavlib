package main

import (
	"fmt"
	"os"

	libgen "github.com/team-rocos/gomavlib/commands/dialgen/libgen"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.CommandLine.Help = "Convert Mavlink dialects from XML format into Go format."

	preamble := kingpin.Flag("preamble", "preamble comment").String()
	mainDefAddr := kingpin.Arg("xml", "a path or url pointing to a XML Mavlink dialect").Required().String()

	// If common.xml is in a different file location to the main xml file specified, and it's included in the main xml file,
	// then the location to common.xml should be specified as a command line argument. If common.xml is included, and is in
	// the same location as the main xml file specified, then this additional argument is optional.
	commonAddr := kingpin.Arg("common.xml", "a path or url pointing to the common XML Mavlink dialect - common.xml").String()

	kingpin.Parse()

	err := libgen.GenerateGoCode(*preamble, *mainDefAddr, *commonAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %s\n", err)
		os.Exit(1)
	}
}
