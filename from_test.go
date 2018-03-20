package flaw_test

import (
	"errors"
	"testing"

	"github.com/halorium/flaw"
)

func TestFrom(t *testing.T) {
	builtinError := errors.New("error message")

	flawError := flaw.From(builtinError)

	equals(
		t,
		builtinError.Error(),
		flawError.Error(),
	)
}
