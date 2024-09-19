/*
Copyright 2016 The Kubernetes Authors.
*/

package version

import (
	"fmt"
	"runtime"
)

var (
	// VERSION, BUILD_DATE, GIT_COMMIT are filled in by the build script
	VERSION     = "<Will be added by go build>"
	COMMIT_SHA1 = "<Will be added by go build>"
	BUILD_DATE  = "<Will be added by go build>"
)

func GetVersion() string {
	return fmt.Sprintf("Version: %s, Commit SHA: %s, Build Date: %s, Go Version: %s", VERSION, COMMIT_SHA1, BUILD_DATE, runtime.Version())
}
