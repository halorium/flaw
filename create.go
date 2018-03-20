package flaw

import (
	"path"
	"runtime"
	"strings"
)

// create is used by New and From to return a flaw Error
// set the first message trace and stack trace
func create(message string) *flaw {
	_, pathname, line, ok := runtime.Caller(2)

	repoRoot := path.Dir(pathname)

	if !ok {
		panic("not ok")
	}

	messageTrace := []frame{
		{
			Message:  message,
			Pathname: stripPathname(pathname, repoRoot),
			Line:     line,
		},
	}

	return &flaw{
		messageTrace: messageTrace,
		stackTrace:   getStackFrames(repoRoot),
	}
}

func getStackFrames(repoRoot string) []frame {
	frames := []frame{}

	atTop := true

	for i := 1; ; i++ {
		_, pathname, line, ok := runtime.Caller(i)

		if !ok {
			break
		}

		// do not include the flaw files in the stack trace
		if atTop && strings.Contains(pathname, "/flaw/") {
			continue
		}

		atTop = false

		stackFrame := frame{
			Pathname: stripPathname(pathname, repoRoot),
			Line:     line,
		}

		frames = append(frames, stackFrame)
	}

	// remove go runtime entrypoints
	return frames[0 : len(frames)-2]
}
