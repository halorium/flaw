package flaw_test

import (
	"fmt"
	"testing"

	"github.com/halorium/flaw"
)

func TestNew(t *testing.T) {
	errorMessage := "error message"

	err := flaw.New(errorMessage)

	equals(
		t,
		fmt.Sprintf("%T", err),
		"*flaw.flawError",
	)

	equals(
		t,
		errorMessage,
		err.Error(),
	)
}
