package common

//请求错误
type requireError struct {
	s string
}

func NewRequireError(s string) *requireError {
	return &requireError{s}
}

func (e *requireError) Error() string {
	return e.s
}

func IsRequireError(err error) bool {
	_, ok := err.(*requireError)
	return ok
}
