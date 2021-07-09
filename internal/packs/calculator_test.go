package packs

import (
	"reflect"
	"testing"
)

func TestCalculator(t *testing.T) {
	packSizes := []int{250, 500, 1000, 2000, 5000}

	testCases := []struct {
		count    int
		expected Packs
	}{
		{
			count:    1,
			expected: Packs{Pack{250, 1}},
		},
		{
			count:    250,
			expected: Packs{Pack{250, 1}},
		},
		{
			count:    251,
			expected: Packs{Pack{500, 1}},
		},
		{
			count:    501,
			expected: Packs{Pack{250, 1}, Pack{500, 1}},
		},
		{
			count:    12001,
			expected: Packs{Pack{250, 1}, Pack{2000, 1}, Pack{5000, 2}},
		},
	}

	for _, testCase := range testCases {
		actual := Calculate(testCase.count, packSizes)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("incorrect calculation for %d, expected - %v, actual - %v\n", testCase.count, testCase.expected, actual)
		}
	}
}

func TestCalculatorEmpty(t *testing.T) {
	actual := Calculate(251, []int{})
	if len(actual) != 0 {
		t.Fail()
	}
}

func TestCalculator0Count(t *testing.T) {
	actual := Calculate(0, []int{250, 500, 1000, 2000, 5000})
	if len(actual) != 0 {
		t.Fail()
	}
}
