// Exercise1_q1.go file
package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	r   int
	c   int
	mat [][]int
}

func (arr Matrix) getrowcount() int {
	return arr.r
}

func (arr Matrix) getcolcount() int {
	return arr.c
}

func (arr *Matrix) setelement(i, j, val int) {
	arr.mat[i][j] = val
}

func (arr Matrix) addmatrix(arr1 Matrix) Matrix {
	var ans Matrix
	ans.r = arr.r
	ans.c = arr.c
	ans.mat = arr.mat
	for i := 0; i < arr.r; i++ {
		for j := 0; j < arr.c; j++ {
			ans.mat[i][j] = arr.mat[i][j] + arr1.mat[i][j]
		}
	}
	return ans
}

func (arr Matrix) printjson() {
	val, err := json.MarshalIndent(arr.mat, "", "  ")

	if err != nil {
		fmt.Println("Something is Wrong ")
	} else {
		fmt.Println("Everything is Fine", string(val))
	}
}

func main() {

	arr := Matrix{2, 3, [][]int{{1, 2, 3}, {4, 5, 6}}}
	fmt.Println(arr.getrowcount())
	fmt.Println(arr.getcolcount())
	fmt.Println(arr)
	arr.setelement(0, 0, 50)
	fmt.Println(arr)
	arr1 := arr
	fmt.Println(arr.addmatrix(arr1))
	arr.printjson()

}
