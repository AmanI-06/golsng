package main

import "fmt"

func stack1() {

	var stack []int
	fmt.Println("\n Pushing elemnts")
	stack = append(stack, 0)
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)

	for len(stack) > 0 {
		n := len(stack) - 1
		fmt.Println(stack[n])

		stack = stack[:n]

	}
}

type stack []int

func (i *stack) Isnull() bool {
	return len(*i) == 0
}

func (i *stack) Push(m int) {
	*i = append(*i, m)

}

func (i *stack) Pop() (int, bool) {
	if i.Isnull() {
		return -1, false
	} else {
		index := len(*i) - 1
		value := (*i)[index]
		(*i) = (*i)[:index]
		return value, true
	}
}
func stackpointers() {
	var eg stack
	eg.Push(1)
	eg.Push(2)
	eg.Push(3)

	for len(eg) > 0 {
		x, y := eg.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
func main(){
	// stack1()
	stackpointers()
}