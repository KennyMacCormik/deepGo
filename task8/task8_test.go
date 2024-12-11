package main

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errList []error
}

func (e *MultiError) Error() string {
	str := strconv.Itoa(len(e.errList)) + " errors occured:\n"
	for _, err := range e.errList {
		str += "\t* " + err.Error()
	}
	return str + "\n"
}

func NewMultiError() *MultiError {
	return &MultiError{errList: make([]error, 0)}
}

func (e *MultiError) append(err error) {
	e.errList = append(e.errList, err)
}

func Append(err error, errs ...error) *MultiError {
	me, ok := err.(*MultiError)
	if !ok {
		me = NewMultiError()
		if err != nil {
			me.append(err)
		}
	}

	for _, val := range errs {
		me.append(val)
	}
	return me
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
