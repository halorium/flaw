package flaw

import (
	"encoding/json"
	"fmt"
	"strings"
)

// flaw is the custom error type that implements the flaw interface
// as well as the error interface
type flaw struct {
	messageTrace []frame
	stackTrace   []frame
}

// Error returns the originating errors message
// Error implements the builtin error interface on the flaw Error type
func (err *flaw) Error() string {
	return err.messageTrace[len(err.messageTrace)-1].Message
}

// Wrap adds another message on the top of the Message Trace stack
// This can be thought of like wrapping errors with layers as they bubble up
// This is part of the flaw interface.
func (err *flaw) Wrap(message string) Error {
	fe := create(message)

	err.messageTrace = append(fe.messageTrace, err.messageTrace...)

	return err
}

// String adds the stringer interface to Error and is part of the flaw interface
func (err *flaw) String() string {
	return strings.TrimSpace(
		"message trace\n" +
			"-----------\n" +
			err.messageTraceString() +
			"\n" +
			"stack trace\n" +
			"-----------\n" +
			err.stackTraceString(),
	)
}

// MarshalJSON implements the json.Marshaler interface.
func (err *flaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		MessageTrace []frame `json:"message-trace"`
		StackTrace   []frame `json:"stack-trace"`
	}{
		err.messageTrace,
		err.stackTrace,
	})
}

// Private functions

// stackTrace builds the stack trace string for output
func (err *flaw) stackTraceString() string {
	stackTraceString := ""

	for _, stackFrame := range err.stackTrace {
		stackTraceString += fmt.Sprintf(
			"%s:%d\n",
			stackFrame.Pathname,
			stackFrame.Line,
		)
	}

	return stackTraceString
}

// messageTrace builds the message trace string for output
func (err *flaw) messageTraceString() string {
	messageTraceString := ""

	for _, messageFrame := range err.messageTrace {
		messageTraceString += fmt.Sprintf(
			"%s (%s:%d)\n",
			messageFrame.Message,
			messageFrame.Pathname,
			messageFrame.Line,
		)
	}

	return messageTraceString
}
