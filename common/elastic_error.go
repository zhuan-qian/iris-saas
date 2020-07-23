package common

import "fmt"

//elasticSearch错误
type esError struct {
	s string
}

func NewEsError(s string) *esError {
	return &esError{s}
}

func (e *esError) Error() string {
	return e.s
}

func IsEsError(err error) bool {
	_, ok := err.(*esError)
	if ok {
		fmt.Println(ok, err.Error())
	}
	return ok
}
