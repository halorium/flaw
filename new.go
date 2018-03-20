package flaw

// New accepts a string and returns a Flaw
func New(message string) Error {
	return create(message)
}
