package problems

import "testing"

func TestNumDupDigitsAtMostN(t *testing.T) {
	t.Log(numDupDigitsAtMostN(20) == 1)
	t.Log(numDupDigitsAtMostN(100) == 10)
	t.Log(numDupDigitsAtMostN(1000) == 262)
}
