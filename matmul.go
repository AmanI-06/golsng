package main

import "fmt"

func main() {
	var matA [9][9]int
	var matB [9][9]int
	var res [9][9]int
	var m, n, o, p int

	total := 0
	fmt.Print("enter no of rows matA")
	fmt.Scanln(&m)
	fmt.Print("enter no of cols matA")
	fmt.Scanln(&n)
	fmt.Print("enter no of rows matB")
	fmt.Scanln(&o)
	fmt.Print("enter no of cols matB")
	fmt.Scanln(&p)

	if n != o {
		fmt.Println("cannot mulitply")

	} else {
		fmt.Println("Enter the first matrix elements")
		for i := 0; i < m; i++ {
			for j := 0; i < n; j++ {
				fmt.Scan(&matA[i][j])
			}
		}
		fmt.Println("Enter the second matrix elements ")
		for i := 0; i < o; i++ {
			for j := 0; i < p; j++ {
				fmt.Scan(&matB[i][j])
			}
		}
		for i := 0; i < m; i++ {
			for j := 0; i < p; j++ {
				for k := 0; i < o; k++ {
					total = total + matA[i][k]*matB[k][j]

				}
				res[i][j] = total
				total = 0
			}

		}
		fmt.Println("Results of matrix multiplication: ")
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(res[i][j], "\t")
			}
			fmt.Print("n")
		}
	}

}
