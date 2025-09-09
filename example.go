package main

import "fmt"

func search(arr []int, target int) (int, int) {
	left := 0
	right := len(arr) - 1
	for {
		if arr[left]+arr[right] == target {
			return left, right
		}
		if arr[left]+arr[right] > target {
			fmt.Println("right", right)
			right--
		}
		if arr[left]+arr[right] < target {
			fmt.Println("left", left)
			left++
		}
	}
}

// optimise: map
func main() {
	example := []int{1, 3, 5, 8, 7, 9}
	a1, a2 := search(example, 11)
	fmt.Println(a1, a2)
}
