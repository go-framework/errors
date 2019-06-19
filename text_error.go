package errors

// Text error defined.
type TextError struct {
	Text string `json:"text"`
}

func (e *TextError) Error() string {
	return e.Text
}

// New returns an error that formats as the given text.
func NewTextError(text string) error {
	return &TextError{Text: text}
}
