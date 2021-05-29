package runnest_test

import (
	"testing"

	"github.com/rodriez/runnest"
)

type fakeService struct {
	active bool
}

func (s *fakeService) ping() string {
	if s.active {
		return "pong"
	}

	return ""
}

func TestRun(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given ping When while service is active Then return pong",
			Given: func() interface{} {
				return true
			},
			When: func(req interface{}) (interface{}, error) {
				active := req.(bool)
				serv := &fakeService{active}

				return serv.ping(), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				if e != nil {
					t.Errorf("Error: %s", e.Error())
				}

				if pong := resp.(string); pong != "pong" {
					t.Errorf("Error: expected pong Received %s", pong)
				}
			},
		},
		{
			Name: "Given ping When while service is inactive Then return empty string",
			Given: func() interface{} {
				return false
			},
			When: func(req interface{}) (interface{}, error) {
				active := req.(bool)
				serv := &fakeService{active}

				return serv.ping(), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				if e != nil {
					t.Errorf("Error: %s", e.Error())
				}

				if pong := resp.(string); pong != "" {
					t.Errorf("Error: expected pong Received %s", pong)
				}
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}
