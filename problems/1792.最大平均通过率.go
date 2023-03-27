package problems

import "container/heap"

/*
 * @lc app=leetcode.cn id=1792 lang=golang
 *
 * [1792] 最大平均通过率
 */

// @lc code=start
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	obj := hpa(classes)
	heap.Init(&obj)
	for ; extraStudents > 0; extraStudents-- {
		obj[0][0]++
		obj[0][1]++
		heap.Fix(&obj, 0)

	}
	total := float64(0)
	for _, v := range obj {
		total += float64(v[0]) / float64(v[1])
	}
	return total / float64(obj.Len())
}

type hpa [][]int

func (h hpa) Len() int {
	return len(h)
}

func (h hpa) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hpa) Less(i int, j int) bool {
	a, b := h[i], h[j]
	return ((float64(a[0]+1) / float64(a[1]+1)) - (float64(a[0]) / float64(a[1]))) > (float64(b[0]+1)/float64(b[1]+1) - float64(b[0])/float64(b[1]))
	//return (a[1]-a[0])*b[1]*(b[1]+1) > (b[1]-b[0])*a[1]*(a[1]+1)
}

func (h hpa) Pop() (_ interface{}) {
	return
}

func (h hpa) Push(interface{}) {

}

// @lc code=end
