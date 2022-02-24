package parser

import (
	"errors"
	"testing"
)

func TestCheckFailsThenPanics(t *testing.T) {
	defer func() {
		recover()
	}()

	Check(errors.New("test"))

	t.Errorf("did not panic")
}

func TestCheckSuccess(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Should not panic")
		}
	}()

	Check(nil)
}
