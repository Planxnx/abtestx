package abtestx

import (
	"sync/atomic"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// RoundRobinTest is a test data for round-robin strategy.
type RoundRobinTest struct {
	ID       string       // Required
	Callback func() error // Optional
}

// RoundRobin is a ab-test with round-robin strategy.
type RoundRobin struct {
	tests []*RoundRobinTest
	ptr   uint64
}

// NewRoundRobin creates a new ab-test with round-robin strategy instance.
func NewRoundRobin(tests []RoundRobinTest) Client {
	if len(tests) == 0 {
		panic(errors.New("tests is empty"))
	}

	t := lo.ToSlicePtr(tests)

	for i, test := range t {
		// validating
		if test.ID == "" {
			panic(errors.Errorf("test.ID[%v] is empty", i))
		}
	}

	return &RoundRobin{
		tests: t,
		ptr:   0,
	}
}

// Run using round-robin to choose test and execute the callback of the test.
func (c *RoundRobin) Run() error {
	_, callback := c.Pick()
	if callback != nil {
		return callback()
	}
	return nil
}

// Pick  using round-robin to choose test returns a test.
func (c *RoundRobin) Pick() (id string, callback func() error) {
	if test := c.tests[atomic.AddUint64(&c.ptr, 1)%uint64(len(c.tests))]; test != nil {
		return test.ID, test.Callback
	}
	return "", nil
}
