/*
Package caterr creates errors with a category to help understand how to handle them.
*/
package caterr

import (
	errors "golang.org/x/xerrors"
)

// Interface can be used
type Interface interface {
	error
	Category() interface{}
	Unwrap() error
	As(interface{}) bool
	Is(error) bool
}

type impl struct {
	category interface{}
	message  string
	wrapped  error
}

// New creates an error with a category
func New(category interface{}, message string) error {
	return &impl{
		category: category,
		message:  message,
	}
}

// Wrap wraps an error with a category and message
func Wrap(err error, category interface{}, message string) error {
	if err == nil {
		return nil
	}
	return &impl{
		category: category,
		message:  message,
		wrapped:  err,
	}
}

// IsCategory determines if the error has a category in its stack
func IsCategory(err error, category interface{}) bool {
	if err == nil {
		return false
	}
	var cat Interface = &impl{}
	for {
		if errors.As(err, &cat) {
			if cat.Category() == category {
				return true
			}
		}
		err = errors.Unwrap(err)
		if err == nil {
			return false
		}
	}
}

func (e *impl) Error() string {
	if e.wrapped != nil {
		return e.message + " : " + e.wrapped.Error()
	}
	return e.message
}

func (e *impl) As(err interface{}) bool {
	_, ok := err.(Interface)
	return ok
}

func (e *impl) Is(err error) bool     { return e.As(err) }
func (e *impl) Unwrap() error         { return e.wrapped }
func (e *impl) Category() interface{} { return e.category }
