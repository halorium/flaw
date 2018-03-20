package flaw_test

import (
	"encoding/json"
	"testing"

	"github.com/halorium/flaw"
)

func TestFlaw(t *testing.T) {
	expected := "some error"
	flawError := flaw.New(expected)

	// implements error interace
	equals(
		t,
		expected,
		flawError.Error(),
	)

	// implements Wrap function
	flawError = flawError.Wrap("second message")
	// implements fmt.Stringer interface
	expected = "message trace\n-----------\nsecond message (flaw_test.go:22)\nsome error (flaw_test.go:12)\n\nstack trace\n-----------"
	equals(
		t,
		expected,
		flawError.String(),
	)

	// implements json.Marshaler interface
	expected = "{\"message-trace\":[{\"message\":\"second message\",\"pathname\":\"flaw_test.go\",\"line\":22},{\"message\":\"some error\",\"pathname\":\"flaw_test.go\",\"line\":12}],\"stack-trace\":[]}"

	sliceBytes, err := json.Marshal(flawError)

	ok(t, err)

	equals(
		t,
		expected,
		string(sliceBytes),
	)
}
