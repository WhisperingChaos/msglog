package msg

/*
Interface to output messages.

  P  - Expects message content is already completely rendered.

*/

type I interface {
	P(msg string, fields ...map[string]interface{})
}

func NewDiscard() I {
	d := new(discard)
	return d
}

// private --------------------------------------------------------------------

type discard struct {
}

func (discard) P(msg string, fields ...map[string]interface{}) {}
