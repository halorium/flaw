package flaw

// From accepts a builtin error and returns a flaw Error
func From(err error) Error {
	if err == nil {
		return nil
	}

	return create(err.Error())
}
