package problems

import "testing"

func TestBaseNeg2(t *testing.T) {
	t.Log(baseNeg2(0) == "0")
	t.Log(baseNeg2(2) == "110")
	t.Log(baseNeg2(3) == "111")
	t.Log(baseNeg2(4) == "100")
	t.Log(baseNeg2(5) == "101")
	t.Log(baseNeg2(6) == "110")
	t.Log(baseNeg2(7) == "111")
	t.Log(baseNeg2(8) == "1000")
	t.Log(baseNeg2(9) == "1001")
	t.Log(baseNeg2(10) == "1010")
}
