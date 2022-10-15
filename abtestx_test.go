package abtestx

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	t.Error("")

	abtest := New([]Test{
		{
			ID:     "A",
			Weight: 0.9,
		},
		{
			ID:     "B",
			Weight: 0.1,
		},
		{
			ID:     "C",
			Weight: 0.0,
		},
	})

	counter := map[string]int{}
	for i := 0; i < 20; i++ {
		test := abtest.Pick()
		counter[test.ID]++
	}

	for k, v := range counter {
		fmt.Printf("%s: %d\n", k, v)
	}
}
