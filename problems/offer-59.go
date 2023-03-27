package problems

type MaxQueue struct {
	que, maxQue []int
}

func Constructor() MaxQueue {
	return MaxQueue{que: []int{}, maxQue: []int{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxQue) == 0 {
		return -1
	}
	return this.maxQue[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.que = append(this.que, value)
	for len(this.maxQue) > 0 && this.maxQue[len(this.maxQue)-1] < value {
		this.maxQue = this.maxQue[:len(this.maxQue)]
	}
	this.maxQue = append(this.maxQue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.que) == 0 {
		return -1
	}
	val := this.que[0]
	this.que = this.que[1:]
	if val == this.maxQue[0] {
		this.maxQue = this.maxQue[1:]
	}
	return val
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
