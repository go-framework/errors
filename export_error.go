package errors

// Export error defined.
type ExportError struct {
	Text string `json:"error"`
}

func (e *ExportError) Error() string {
	return e.Text
}

// New returns an error that formats as the given err.
func NewExportError(err interface{}) error {
	if err == nil {
		return nil
	}
	return &ExportError{Text: GetDetail(err)}
}
