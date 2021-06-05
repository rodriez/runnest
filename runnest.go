package runnest

import (
	"testing"
)

type Runest struct {
	t *testing.T
}

func NewRunest(t *testing.T) *Runest {
	return &Runest{t}
}

func (run *Runest) Run(cases []TestCase) {
	for _, tc := range cases {
		if !tc.Skip && tc.Valid() {
			tc.Run(run.t)
		}
	}
}
