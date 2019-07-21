package caterr_test

import (
	"github.com/mvndaai/caterr"
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {
	type s string
	categories := []interface{}{1, "category", s("category"), true, false, nil}

	for _, category := range categories {
		err := errors.New("errors")
		if caterr.IsCategory(err, category) {
			t.Error("this should not have matched the category")
		}

		err = caterr.Wrap(err, category, "caterr")
		if !caterr.IsCategory(err, category) {
			t.Error("did not match the category")
		}
	}
}

func TestNew(t *testing.T) {
	category := 0
	err := caterr.New(category, "caterr")
	if !caterr.IsCategory(err, category) {
		t.Error("could not find category")
	}
}

type wrapper struct {
	message string
	wrapped error
}

func (w *wrapper) Unwrap() error { return w.wrapped }
func (w *wrapper) Error() string { return w.message + " : " + w.wrapped.Error() }

func TestExternalWrap(t *testing.T) {
	category := 0
	err := caterr.New(category, "caterr")
	err = &wrapper{message: "wrapper", wrapped: err}
	if !caterr.IsCategory(err, category) {
		t.Error("could not find category")
	}
}

func TestNils(t *testing.T) {
	category := 0
	err := caterr.Wrap(nil, category, "caterr")
	if err != nil {
		t.Error("Wrapping nil should have resulted in nil")
	}
	if caterr.IsCategory(err, category) {
		t.Error("nils should not match a category")
	}
}

func TestWrapMultipleCategories(t *testing.T) {
	category := 0
	err := caterr.New(category, "caterr")
	err = caterr.Wrap(err, "different category", "foo")
	if !caterr.IsCategory(err, category) {
		t.Error("it should have matched the bottom cateogry")
	}
}
