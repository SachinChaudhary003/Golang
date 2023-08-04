package main

import (
	"fmt"
	"sort"
)

func main() {

	//Arrays
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)

	fmt.Println("len:", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//Slices

	var s []int
	fmt.Println(s)
	s = append(s, 1, 2, 3)
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	score := make([]int, 4)
	score[0] = 2
	score[1] = 4
	score[2] = 6
	score[3] = 8
	fmt.Println(score)

	score = append(score, 10)
	fmt.Println(score)

	sort.Ints(score)
	fmt.Println(score)
}
