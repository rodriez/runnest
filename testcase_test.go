package runnest_test

import (
	"testing"

	"github.com/rodriez/runnest"
)

func TestTestCaseRun(t *testing.T) {
	var str string

	runnest.TestCase{
		Name: "Given valid testcase When Run is called Then execute the testcase succesfully",
		Before: func() {
			str = ""
		},
		Given: func() interface{} {
			return str
		},
		When: func(req interface{}) (interface{}, error) {
			str = "First"

			return str, nil
		},
		Then: func(t *testing.T, resp interface{}, e error) {
			if e != nil {
				t.Errorf("Error: %s", e.Error())
			}

			str := resp.(string)

			if str != "" && str != "First" {
				t.Errorf("Error: First pong Received %s", str)
			}
		},
	}.Run(t)
}

func TestTestCaseValidate(t *testing.T) {
	valid := runnest.TestCase{
		Name: "Given invalid test When while service is active Then return pong",
	}.Valid()

	if valid {
		t.Errorf("Validate test failed, expected false but got true")
	}
}
