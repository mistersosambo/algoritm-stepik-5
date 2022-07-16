package main

import "fmt"

func main() {
	var n, z int
	var m []int
	count := 0
	fmt.Scan(&n)
	for j := 0; j < n; j++ {
		fmt.Scan(&z)
		m = append(m, z)
	}
	for _, m := range m {
		if m > 0 {
			count += 1
		} else {
			continue
		}
	}
	fmt.Println(count)
}
