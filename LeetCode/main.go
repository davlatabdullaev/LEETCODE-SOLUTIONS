package main

import (
	"fmt"
	monotonicarray "leetcode/MonotonicArray"
)

func main() {
	// nums := []int{1, 3, 5, 6}
	// target := 7
	// searchinsertposition.SearchInsert(nums, target)
	// 	digits := []int{1,2,3}
	// 	plusone.PlusOne(digits)

	//  a := "1"
	//  b := "3"
	//  c := addbinary.AddBinary(a,b)
	//  fmt.Println(c)

	// son := sqrtx.MySqrt(8)
	// fmt.Println(son)

	// a:=issubsequence.IsSubsequence("abc", "ahjdj")

	// fmt.Println(a)

	// a := [][]int{
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 0},
	// 	{1, 0, 0, 0, 0},
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 1},
	// }

	// b := thekweakestrowsinamatrix.KWeakestRows(a, 3)
	// fmt.Println(b)

	// A := findthedifference.FindTheDifference("h", "hj")

	// fmt.Println(A)

	nums := []int{1,2,3,4}
	m := monotonicarray.IsMonotonic(nums)
	fmt.Println(m)

}
