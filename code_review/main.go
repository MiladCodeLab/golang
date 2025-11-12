package main

import "fmt"

/*
Input:

[1, 2, 3]
[4, 5, 6] (map: key(2,3), val = 21) O(1)
[7, 8, 9]

n^2 + o(1)= n^2
Output:

[1, 3, 6] indexX = 2, indexY = 3
[5, 12, 21]
[12, 27, 45]
*/

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var result [3][3]int
	fmt.Println(len(arr))
	//result := make([][]int, len(arr))
	for indexX := 0; indexX < 3; indexX++ {
		for indexY := 0; indexY < 3; indexY++ {
			//fmt.Println(indexX, indexY)
			result[indexX][indexY] = cal(arr, indexX, indexY)
		}
	}

	fmt.Println(result)
}
func cal(arr [][]int, indexX, indexY int) int {
	number := 0
	for x := 0; x <= indexX; x++ {
		for y := 0; y <= indexY; y++ {
			//fmt.Println(x, "  ", y, "  ", arr[x][y])
			number += arr[x][y]
		}
	}
	return number
}
