package main

import (
	"fmt"
	"math"
)

func main() {
	list := [][]int32{
		{-10, 3, 0, 5, -4},
		{2, -1, 0, 2, -8},
		{9, -2, -5, 6, 0},
		{9, -7, 4, 8, -2},
		{3, 7, 8, -5, 0},
	}

	result := math.Abs(float64(calculateDiagonal(0, 0, 0, list)))
	fmt.Println(result)
}

type MapKey struct {
	x, y int32
}

var cache = make(map[MapKey]int32)

func calculateDiagonal(primary, second int32, row int, list [][]int32) int32 {
	if v, ok := cache[MapKey{primary, second}]; ok {
		return v
	}

	if row == len(list) {
		return int32(primary - second)
	}

	output := calculateDiagonal(primary+list[row][row],
		second+list[row][len(list)-1-row], // 5-1-1 = 3
		row+1,                             // 1
		list)
	cache[MapKey{primary, second}] = output

	return output
}
