package main

import "fmt"

func array(n int) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])

	}
	fmt.Println("The array is :", a)
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if a[i] == a[j] {
				count += 1
			}
		}
		if count%2 != 0 {
			fmt.Println("The odd occurring element is :", a[i])

		}

	}

}
func main() {

	var n int
	fmt.Scan(&n)
	array(n)
}
