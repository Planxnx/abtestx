package abtestx_test

import (
	"errors"
	"testing"

	"github.com/Planxnx/abtestx"
	testifyassert "github.com/stretchr/testify/assert"
)

func TestWeightedRandom(t *testing.T) {
	t.Run("success_case", func(t *testing.T) {
		assert := testifyassert.New(t)

		counter := map[string]int{}
		testData := []abtestx.WeightedRandomTest{
			{
				ID:     "0",
				Weight: 0.7,
				Callback: func() error {
					counter["0"]++
					return nil
				},
			},
			{
				ID:     "1",
				Weight: 0.3,
				Callback: func() error {
					counter["1"]++
					return nil
				},
			},
			{
				ID:     "2",
				Weight: 0.0,
				Callback: func() error {
					counter["2"]++
					return nil
				},
			},
		}

		assert.NotPanics(func() {
			abtest := abtestx.NewWeightedRandom(testData)
			for i := 0; i < 100; i++ {
				err := abtest.Run()
				assert.NoError(err)
			}
		})

		for _, test := range testData {
			if test.Weight > 0 {
				assert.NotZerof(counter[test.ID], "id:%v isn't execute", test.ID)
			} else {
				assert.Zerof(counter[test.ID], "id:%v shouldn't execute", test.ID)
			}
		}
	})

	t.Run("fail_case", func(t *testing.T) {
		assert := testifyassert.New(t)

		assert.Panics(func() {
			_ = abtestx.NewWeightedRandom(nil)
		})

		assert.Panics(func() {
			_ = abtestx.NewWeightedRandom([]abtestx.WeightedRandomTest{
				{}, {}, {},
			})
		})

		assert.Panics(func() {
			_ = abtestx.NewWeightedRandom([]abtestx.WeightedRandomTest{
				{ID: "0", Weight: 0.5},
				{ID: "1", Weight: 0.5},
				{ID: "2", Weight: 0.5},
			})
		})

		assert.NotPanics(func() {
			abtest := abtestx.NewWeightedRandom([]abtestx.WeightedRandomTest{
				{ID: "0", Weight: 0.6},
				{ID: "1", Callback: func() error {
					return errors.New("abtestx error")
				}},
			})
			var isError, isNotError bool
			for i := 0; i < 100; i++ {
				err := abtest.Run()
				if err != nil {
					isError = true
				} else {
					isNotError = true
				}
			}
			assert.True(isError && isNotError, "expected tests is should have error and not error")
		})
	})
}
