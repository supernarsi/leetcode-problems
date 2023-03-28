package problems

import "testing"

func TestCountAndSay(t *testing.T) {
	t.Log(countAndSay(1) == "1")
	t.Log(countAndSay(2) == "11")
	t.Log(countAndSay(3) == "21")
	t.Log(countAndSay(4) == "1211")
	t.Log(countAndSay(5) == "111221")
	t.Log(countAndSay(6) == "312211")
	t.Log(countAndSay(7) == "13112221")
	t.Log(countAndSay(8) == "1113213211")
	t.Log(countAndSay(9) == "31131211131221")
	t.Log(countAndSay(10) == "13211311123113112211")
}
