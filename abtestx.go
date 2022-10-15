package abtestx

import (
	"sort"

	"github.com/Planxnx/abtestx/pkg/rand"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

const (
	DefaultTotalWeight float64 = 1.0
)

type Test struct {
	ID            string       // Required
	Weight        float64      // Optional
	Callback      func() error // Optional
	runningWeight float64
}

type Client struct {
	tests       []*Test
	totalWeight float64
}

func New(tests []Test) *Client {
	if len(tests) == 0 {
		panic(errors.New("tests is empty"))
	}

	t := fillWeight(DefaultTotalWeight, lo.ToSlicePtr(tests))

	sort.Slice(t, func(i, j int) bool {
		return t[i].Weight < t[j].Weight
	})

	totalWeight := 0.0
	for i, test := range t {
		// validating
		if test.ID == "" {
			panic(errors.Errorf("test.ID[%v] is empty", i))
		}

		totalWeight += test.Weight
		test.runningWeight = totalWeight
	}

	if totalWeight > 1.0 {
		panic(errors.Errorf("invalid total weight, must be 1.0: %v", totalWeight))
	}

	return &Client{
		tests:       t,
		totalWeight: totalWeight,
	}
}

// Run using weighted random to choose test and execute the callback of the test.
func (c *Client) Run() error {
	test := c.Pick()
	if test.Callback != nil {
		return test.Callback()
	}
	return nil
}

// Pick  using weighted random to choose test returns a test.
func (c *Client) Pick() Test {
	w := rand.Floats64n(c.totalWeight)
	for _, test := range c.tests {
		if test.runningWeight > w {
			return *test
		}
	}
	return *c.tests[len(c.tests)-1]
}

func fillWeight(totalWeight float64, tests []*Test) []*Test {
	emptyWeightTests := make([]*Test, 0)
	accumulatedWeight := 0.0

	for _, test := range tests {
		if test.Weight <= 0 {
			emptyWeightTests = append(emptyWeightTests, test)
		}

		accumulatedWeight += test.Weight
	}

	emptyWeightLen := len(emptyWeightTests)
	if emptyWeightLen > 0 {
		weight := (totalWeight - accumulatedWeight) / float64(emptyWeightLen)
		for _, test := range emptyWeightTests {
			test.Weight = weight
		}
	}

	return tests
}
