package runnest

import (
	"testing"
)

type TestCase struct {
	Name   string
	Skip   bool
	Before func()
	Given  func() interface{}
	When   func(req interface{}) (interface{}, error)
	Then   func(t *testing.T, resp interface{}, e error)
}

func (test TestCase) Valid() bool {
	return test.When != nil && test.Then != nil
}

func (test TestCase) Run(t *testing.T) {
	t.Run(test.Name, func(t *testing.T) {
		if test.Before != nil {
			test.Before()
		}

		var req interface{}
		if test.Given != nil {
			req = test.Given()
		}

		resp, e := test.When(req)
		test.Then(t, resp, e)
	})
}
