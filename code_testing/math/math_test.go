package math_test

import (
	"code_testing/math"
	"testing"
)

func TestAdd(t *testing.T) {
	total := math.Add(2, 2)
	if total != 4 {
		t.Errorf("Sum was incorrect, Actual: %d, Expected: %d", total, 4)
	}
	t.Log("running TestAdd")
}
