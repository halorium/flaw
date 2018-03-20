package flaw

// frame is a single unit in a collection of message trace or stack trace
type frame struct {
	Message  string `json:"message,omitempty"`
	Pathname string `json:"pathname,omitempty"`
	Line     int    `json:"line,omitempty"`
}
