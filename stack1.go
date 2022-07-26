package main

import "fmt"

func push(s *[]int, i int) {
	*s = append(*s, i)

}
func pop(s *[]int) {
	leng := len(*s) - 1
	if leng <= 0 {
		fmt.Println("Stack Underflow")
	} else {
		*s = (*s)[:leng]
	}
}
func main() {
	s := make([]int, 4)
	push(&s, 0)
	push(&s, 1)
	push(&s, 2)
	push(&s, 3)
	fmt.Printf("%d",s)
	pop(&s)
	pop(&s)
	fmt.Printf("%d",s)
}
