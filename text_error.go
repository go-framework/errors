package errors

// Text error defined.
type TextError struct {
	Text string `json:"error"`
}

func (e *TextError) Error() string {
	return e.Text
}

// New returns an error that formats as the given text.
func NewTextError(text string) error {
	if len(text) == 0 {
		return nil
	}
	return &TextError{Text: text}
}
