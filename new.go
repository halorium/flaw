package flaw

// New accepts a string and returns a Flaw
func New(message string) Flaw {
	return create(message)
}
