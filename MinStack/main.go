package main

import "fmt"

type MinStack struct {
	Slice []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{
		Slice:[]int{},
	}
	return stack
}


func (this *MinStack) Push(x int)  {
	this.Slice = append(this.Slice, x)
}


func (this *MinStack) Pop()  {
	this.Slice = this.Slice[:len(this.Slice) - 1]
}


func (this *MinStack) Top() int {
	return this.Slice[len(this.Slice) - 1]
}


func (this *MinStack) GetMin() int {
	min := this.Slice[0]
	for i := 1; i < len(this.Slice); i++ {
		if this.Slice[i] < min {
			 min = this.Slice[i]
		}
	}
	return min
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
}

