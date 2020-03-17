package gomavlib

import (
	"testing"

	"github.com/stretchr/testify/require"
	libgen "github.com/team-rocos/gomavlib/commands/dialgen/libgen"
)

// DEFINE PUBLIC TYPES AND STRUCTURES.

// DEFINE PRIVATE TYPES AND STRUCTURES.

// DEFINE PUBLIC STATIC FUNCTIONS.

func TestDialectRT(t *testing.T) {
	// Parse the XML file.
	includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("./mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml", includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	drt, err = NewDialectRT(version, defs)
	require.NoError(t, err)

}

// DEFINE PUBLIC RECEIVER FUNCTIONS.

// DEFINE PRIVATE STATIC FUNCTIONS.

// DEFINE PRIVATE RECEIVER FUNCTIONS.

// ALL DONE.
