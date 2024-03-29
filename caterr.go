/*
Package caterr creates errors with a category to help understand how to handle them.
*/
package caterr

import "errors"

// Interface can be used for type checking especially in As
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

// HasCategory checks if an error has a specified category
func HasCategory(err error, category interface{}) bool {
	return errors.Is(err, &impl{category: category})
}

func (e *impl) Error() string {
	if e.wrapped != nil {
		return e.message + " : " + e.wrapped.Error()
	}
	return e.message
}

func (e *impl) As(err interface{}) bool {
	if caterr, ok := err.(Interface); ok {
		return caterr.Category() == e.Category()

	}
	return false
}

func (e *impl) Is(err error) bool     { return e.As(err) }
func (e *impl) Unwrap() error         { return e.wrapped }
func (e *impl) Category() interface{} { return e.category }
