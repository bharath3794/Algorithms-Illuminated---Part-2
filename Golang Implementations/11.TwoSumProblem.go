package main

import (
	"fmt"
	"math"
	"sort"
)

type pair struct {
	p []int
	cnt int
}

func main() {
	arr1 := []int{6, 5, 6, -7, -8, -8, -2, 8, -8, -8}
	arr2 := []int{-8, -5, -2, 1, -1, -1, -1, -1, 6, 8, -8, 3, 4, -3, -10, 6}

	// Method 1 : Using Hash Map
	pairs := twoSumHashMap1(arr1, -2)
	fmt.Println("Pairs that add upto given target =", pairs)

	pairs = twoSumHashMap1(arr2, -2)
	fmt.Println("Pairs that add upto given target =", pairs)

	// Method 2 : Using Hash Map
	newPairs := twoSumHashMap2(arr1, -2)
	fmt.Println("Pairs that add upto given target = ")
	for _, val := range newPairs{
		fmt.Printf("pair = %v, count = %v\n", val.p, val.cnt)
	}
	
	// Method 3 : Recursive Solution (Divide & Conquer)
	newPairs = twoSumHashMap2(arr2, -2)
	fmt.Println("Pairs that add upto given target = ")
	for _, val := range newPairs{
		fmt.Printf("pair = %v, count = %v\n", val.p, val.cnt)
	}

	pairs2 := twoSumRec(arr1, -2, true)
	fmt.Println("Pairs that add upto given target =", pairs2)

	pairs2 = twoSumRec(arr2, -2, true)
	fmt.Println("Pairs that add upto given target =", pairs2)

}


// Method 1: Using Hash Map
// Two-sum problem using hash map
// Returns only distinct pairs i.e. if (x, y) is appeared twice or more returned just once
func twoSumHashMap1(arr []int, target int) [][]int {
	dict := map[int][]int{}
	for _, v := range arr {
		dict[v] = []int{target-v, 0}
	}
	pairs := [][]int{}
	for key, val := range dict{
		if _, ok := dict[val[0]]; ok && val[1] == 0 {
			pairs = append(pairs, []int{key, val[0]})
			dict[key][1] = 1
			dict[val[0]][1] = 1
		}
	}
	return pairs
}



// Method 2: Extension to previous Method 2 (Along with pairs it also gives count)
// (i.e. how many of same (x, y) pair are present as there are repeated elements)
func twoSumHashMap2(arr []int, target int) []pair{
	dict := map[int][]int{}
	for _, v := range arr {
		if _, ok := dict[v]; !ok{
			dict[v] = []int{target-v, 1, 0}
		} else {
			dict[v][1] += 1
		}
	}
	pairs := []pair{}
	for key, val := range dict {
		if _, ok := dict[val[0]]; ok && val[2] == 0 {
			var cnt int
			if key != val[0] {
				cnt = int(math.Min(float64(dict[key][1]), float64(dict[val[0]][1])))
			} else {
				cnt = val[1]/2
			}
			pairs = append(pairs, pair{[]int{key, val[0]}, cnt})
			dict[key][2] = 1
			dict[val[0]][2] = 1
		}
	}
	return pairs
}

// Global Variable which is a dependency for below twoSumRec() function.
var pairs = [][]int{}
var added = map[int]int{}
// Method 3: Recursive Solution (Divide & Conquer)
// Returns only distinct pairs i.e. if (x, y) is appeared twice or more returned just once
func twoSumRec(arr []int, target int, initialSort bool) [][]int{
	if len(arr) == 1{
		return [][]int{arr}
	}
	if initialSort == true {
		sort.Ints(arr)
		pairs = [][]int{}
		added = map[int]int{}
	}
	mid := (len(arr)-1)/2
	left := twoSumRec(arr[:mid+1], target, false)
	right := twoSumRec(arr[mid+1:], target, false)
	rightMod := []int{}
	for _, v := range right {
		rightMod = append(rightMod, v[0])
	}
	for _, ele := range left {
		if _, ok := added[ele[0]]; !ok{
			val, ok := binSearch(rightMod, 0, len(rightMod)-1, target-ele[0])
			if ok {
				pairs = append(pairs, []int{ele[0], val})
				added[ele[0]] = 1
				added[val] = 1
			}
		}
	}
	if initialSort == false {
		return append(left, right...)
	} 
	return pairs
}


//This function is dependency for twoSumRec() function
//Binary Search Implementation: For distinct Elements and gives any one instance if found
func binSearch(v []int, p, r, key int) (int, bool) {
	if p > r{
		return 0, false
	}
	q := (p+r)/2
	if v[q] == key{
		return v[q], true
	} else if v[q] > key {
		return binSearch(v, p, q-1, key)
	}else if v[q] < key {
		return binSearch(v, q+1, r, key)
	}
	return 0, false
}