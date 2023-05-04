package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (e *MyAppError) Error() string {
	return e.Err.Error()
}

func (e *MyAppError) Unwrap() error {
	return e.Err
}
