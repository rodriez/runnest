package runnest

import (
	"testing"
)

type TestCase struct {
	Name  string
	Given func() interface{}
	When  func(req interface{}) (interface{}, error)
	Then  func(t *testing.T, resp interface{}, e error)
}

type Runest struct {
	t *testing.T
}

func NewRunest(t *testing.T) *Runest {
	return &Runest{t}
}

func (run *Runest) Run(cases []TestCase) {
	for _, tc := range cases {
		run.t.Run(tc.Name, func(t *testing.T) {
			req := tc.Given()
			resp, e := tc.When(req)

			tc.Then(t, resp, e)
		})
	}
}
