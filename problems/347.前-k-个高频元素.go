package problems

import (
	"container/heap"
)

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func TopKFrequent(nums []int, k int) (ans []int) {
	tmpMap := map[int]*ele{}
	for _, val := range nums {
		if v, ok := tmpMap[val]; ok {
			v.pri++
		} else {
			tmpMap[val] = &ele{
				val: val,
				pri: 1,
			}
		}
	}

	var arr []*ele
	for _, val := range tmpMap {
		arr = append(arr, val)
	}
	hp := hpP347(arr)
	heap.Init(&hp)
	for i := 0; i < k; i++ {
		e := hp.Pop()
		ans = append(ans, e.(ele).val)
		heap.Fix(&hp, 0)
	}
	return
}

type ele struct {
	val int
	pri int
}

type hpP347 []*ele

func (h hpP347) Less(i int, j int) bool {
	return h[i].pri > h[j].pri
}

func (h hpP347) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hpP347) Len() int {
	return len(h)
}

func (h *hpP347) Pop() interface{} {
	if len(*h) < 1 {
		return nil
	}
	old := *h
	ele := old[0]
	*h = old[1:]
	return *ele
}

func (h hpP347) Push(x interface{}) {

}

// @lc code=end
