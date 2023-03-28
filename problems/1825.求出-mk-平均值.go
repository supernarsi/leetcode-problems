package problems

import "sort"

/*
 * @lc app=leetcode.cn id=1825 lang=golang
 *
 * [1825] 求出 MK 平均值
 */

// @lc code=start
type MKAverage struct {
	m                 int
	k                 int
	containerItems    []int
	sortedCItems      []int
	allItems          []int
	allItemsNum       int
	containerItemsNum int
}

func Constructor1825(m int, k int) MKAverage {
	return MKAverage{
		m:              m,
		k:              k,
		allItems:       []int{},
		containerItems: []int{},
		sortedCItems:   []int{},
	}
}

func (this *MKAverage) AddElement(num int) {
	this.allItems = append(this.allItems, num)
	this.containerItems = append(this.containerItems, num)
	conItemsLen := len(this.containerItems)
	if conItemsLen == this.m {
		// 容器元素刚好撑满，初始化 sortedItems
		tmpItems := make([]int, len(this.containerItems))
		copy(tmpItems, this.containerItems)
		sort.Ints(tmpItems)
		this.sortedCItems = tmpItems
	} else if conItemsLen > this.m {
		// 容器元素多于 m，移出元素
		rmEle := this.containerItems[0]
		this.containerItems = this.containerItems[1:]
		if rmEle != num {
			// 插入元素和移出元素不相等，则重新排序

			// 查找移出元素在有序数组中的位置
			rmIdx := findSortIdx(this.sortedCItems, rmEle)
			this.sortedCItems = append(this.sortedCItems[:rmIdx], this.sortedCItems[rmIdx+1:]...)
			// 查找插入元素在有序数组中的位置
			addIdx := findSortIdx(this.sortedCItems, num)
			this.sortedCItems = append(this.sortedCItems[:addIdx], append([]int{num}, this.sortedCItems[addIdx:]...)...)
		}
	}
}

func (this *MKAverage) CalculateMKAverage() int {
	if len(this.allItems) < this.m {
		return -1
	}
	sortedItemsLen := len(this.sortedCItems)
	if sortedItemsLen <= 2*this.k {
		return 0
	}
	leftItems := this.sortedCItems[this.k : sortedItemsLen-this.k]
	sum := 0
	for _, n := range leftItems {
		sum += n
	}
	return sum / len(leftItems)
}

func findSortIdx(arr []int, num int) int {
	return sort.SearchInts(arr, num)
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
// @lc code=end
