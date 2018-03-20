package flaw

import (
	"path"
	"runtime"
)

var REPO_ROOT string

func init() {
	_, filename, _, ok := runtime.Caller(0)

	if ok {
		REPO_ROOT = path.Dir(filename)
	}
}
