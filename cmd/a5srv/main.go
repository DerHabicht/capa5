package main

import (
	"fmt"

	"github.com/derhabicht/capa5/internal/cmd/a5srv"
)

// BaseVersion is the SemVer-formatted string that defines the current version of plancal.
// Build information will be added at compile-time.
const BaseVersion = "0.1.0"

// BuildTime is a timestamp of when the build is run. This variable is set at compile-time.
var BuildTime string

// GitRevision is the current Git commit ID. If the tree is dirty at compile-time, an "x-" is prepended to the hash.
// This variable is set at compile-time.
var GitRevision string

// GitBranch is the name of the active Git branch at compile-time. This variable is set at compile-time.
var GitBranch string

func main() {
	var version string
	// This happens if the Makefile isn't used to build the server, e.g. a run configuration in the IDE.
	if BuildTime == "" || GitRevision == "" || GitBranch == "" {
		version = BaseVersion
	} else {
		version = fmt.Sprintf(
			"%s+%s.%s.%s",
			BaseVersion,
			GitBranch,
			GitRevision,
			BuildTime,
		)
	}
	a5srv.Execute(version)
}
