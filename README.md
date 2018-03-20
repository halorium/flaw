## flaw is a Go(lang) custom error package

### Features
This was created because sometimes you need a little more information than what
the builtin error interface can provide.

A flaw.Error allows you to "Wrap" your errors as they bubble up adding a message trace
as well as a stack trace to your errors.

The flaw.Error interface implements the following interfaces:
* builtin error interface
* fmt.Stringer interface
* json.Marshaler interface

This allows for detailed error message output in JSON formatted logs like
Loggly and many others.

### Usage
```
go get github.com/halorium/flaw
```

```
package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/halorium/flaw"
)

func main() {
	flawError := MyFunc1()

	if flawError != nil {
		fmt.Println(flawError)
		// original error from external call 2

		fmt.Println(flawError.String())
		// message trace
		// -----------
		// cannot perform MyFunc1 (main.go:34)
		// cannot perform MyFunc2 (main.go:46)
		// original error from external call 2 (main.go:46)
		//
		// stack trace
		// -----------
		// main.go:46
		// main.go:30
		// main.go:11

		data, _ := json.Marshal(flawError)
		fmt.Println(string(data))
		// {
		//   "message-trace": [
		//     {
		//       "message": "cannot perform MyFunc1",
		//       "pathname": "main.go",
		//       "line": 50
		//     },
		//     {
		//       "message": "cannot perform MyFunc2",
		//       "pathname": "main.go",
		//       "line": 62
		//     },
		//     {
		//       "message": "original error from external call 2",
		//       "pathname": "main.go",
		//       "line": 62
		//     }
		//   ],
		//   "stack-trace": [
		//     {
		//       "pathname": "main.go",
		//       "line": 62
		//     },
		//     {
		//       "pathname": "main.go",
		//       "line": 46
		//     },
		//     {
		//       "pathname": "main.go",
		//       "line": 12
		//     }
		//   ]
		// }
	}
}

func MyFunc1() flaw.Error {
	err := ExternalCall1()

	if err != nil {
		// create a flaw Error from a standard error
		// wrap flaw with additional information
		return flaw.From(err).Wrap("cannot perform ExternalCall1")
	}

	flawError := MyFunc2()

	if flawError != nil {
		// wrap flaw with additional information
		return flawError.Wrap("cannot perform MyFunc1")
	}

	return nil
}

func MyFunc2() flaw.Error {
	err := ExternalCall2()

	if err != nil {
		// create a flaw Error from a standard error
		// wrap flaw with additional information
		return flaw.From(err).Wrap("cannot perform MyFunc2")
	}

	return nil
}

func ExternalCall1() error {
	return nil
}

func ExternalCall2() error {
	return errors.New("original error from external call 2")
}
```

#### Credits
This is a refactor from code designed and written by Tom Mornini and Loren Hale
