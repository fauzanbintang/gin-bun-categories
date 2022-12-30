package errs

type defaultError struct {
	code int
	Msg  string `json:"message"`
}

func (e *defaultError) Error() string {
	return e.Msg
}

func DefaultForm(code int, err error) *defaultError {
	return &defaultError{
		code: code,
		Msg:  err.Error(),
	}
}
