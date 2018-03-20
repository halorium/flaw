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
		"*flaw.flaw",
		fmt.Sprintf("%T", err),
	)

	equals(
		t,
		errorMessage,
		err.Error(),
	)
}
