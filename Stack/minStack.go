package Stack

type MinStack struct {
	data    []int
	minData []int
}

func Constructor() MinStack {
	return MinStack{data: make([]int, 0, 5), minData: make([]int, 0, 5)}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.minData) != 0 {
		val = min(val, this.minData[len(this.minData)-1])
	}
	this.minData = append(this.minData, val)
}

func (this *MinStack) Pop() {
	l := len(this.data)
	this.data = this.data[:l-1]
	this.minData = this.minData[:l-1]
}

func (this *MinStack) Top() int {
	l := len(this.data)
	return this.data[l-1]
}

func (this *MinStack) GetMin() int {
	l := len(this.minData)
	return this.minData[l-1]
}
