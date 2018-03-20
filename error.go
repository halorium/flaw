package flaw

import (
	"encoding/json"
	"fmt"
)

// Error interface implements the builtin error interface as well as the
// fmt.Stringer interface, json.Marshaler interface and the Wrap function
type Error interface {
	error
	fmt.Stringer
	json.Marshaler
	Wrap(message string) Error
}
