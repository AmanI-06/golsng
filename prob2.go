package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removedups() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		words := strings.Fields(line)
		for _, word := range words {
			if !seen[word] {
				seen[word] = true
			}
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stdout, "Remove Duplicates %v", err)
	}
	fmt.Println((seen))

}
func main() {
	removedups()
}
