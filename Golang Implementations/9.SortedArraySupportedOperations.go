package main

import (
	"fmt"
	"log"
	"errors"
	"math/rand"
	"time"
)

type array struct {
	a []int
}




func main() {
	arr := array{[]int{-8, -5, -2, 1, 3, 4, 6, 8}}
	idx, ok := binSearch(&arr, 0, len(arr.a)-1, -2)
	if ok {
		fmt.Println("Element found at index =", idx)
	} else {
		fmt.Println("Element Not Found")
	}

	arr = array{[]int{38, 39, 42, 57, 59, 59, 77, 83, 87, 91, 91, 98, 99, 99}}
	idxs, ok := binSearchAllInstances(&arr, 0, len(arr.a)-1, 99)
	if ok {
		fmt.Println("Element found at indexes =", idxs)
	} else {
		fmt.Println("Element Not Found")
	}

	val, ok := predecessor(&arr, 99)
	if ok {
		fmt.Println("Predecessor =", val)
	} else {
		fmt.Println("No Predecessor for this element")
	}

	val, ok = successor(&arr, 87)
	if ok {
		fmt.Println("Successor =", val)
	} else {
		fmt.Println("No Successor for this element")
	}

	idx = selectIndx(&arr, 99)
	fmt.Println("Given element is at index", idx)

	valOrder := selectIndxBook(&arr, 13)
	fmt.Println("The value of given ith smallest element =", valOrder)

	rankVal := rank(&arr, 99)
	fmt.Println("Rank of given element is", rankVal)
}


//Binary Search Implementation: For distinct Elements and gives any one instance if found
func binSearch(v *array, p, r, key int) (int, bool) {
	if p > r{
		return 0, false
	}
	q := (p+r)/2
	if v.a[q] == key{
		return q, true
	} else if v.a[q] > key {
		return binSearch(v, p, q-1, key)
	}else if v.a[q] < key {
		return binSearch(v, q+1, r, key)
	}
	return 0, false
}

//Binary Search Implementation: For distinct or repeated Elements and 
// if found, gives the instance whose index is minimum
func binSearchAllInstances(v *array, p, r, key int) ([]int, bool) {
	if p > r{
		return nil, false
	}
	q := (p+r)/2
	if v.a[q] == key{
		indexes := []int{q}
		for i:=q+1; i<len(v.a); i++ {
			if v.a[i] == key{
				indexes = append(indexes, i)
			} else {
				break
			}
		}
		for i:=q-1; i>=0; i-- {
			if v.a[i] == key{
				indexes = append(indexes, i)
			} else {
				break
			}
		}
		return indexes, true
	} else if v.a[q] > key {
		return binSearchAllInstances(v, p, q-1, key)
	}else if v.a[q] < key {
		return binSearchAllInstances(v, q+1, r, key)
	}
	return nil, false
}

// Return Minimum element
func findMin(v *array) int {
	return v.a[0]
}

// Return Maximum element
func findMax(v *array) int {
	return v.a[len(v.a)-1]
}


// Predrecessor of an element
func predecessor(v *array, key int) (int, bool) {
	var value int
	var truthVal bool
	indx, ok := binSearchAllInstances(v, 0, len(v.a)-1, key)
	if !ok {
		log.Fatal(errors.New("Element Not Found"))
	}
	var minIdx int
	if len(indx) == 1 {
		 minIdx = indx[0]
	}else if len(indx) > 1 {
		minIdx = rSelect(&array{indx}, 0, len(indx)-1, 1)
	}
	if minIdx-1 >= 0{
		value = v.a[minIdx-1]
		truthVal = true
	} else {
		truthVal = false
	}
	return value, truthVal
}


// Successor of an element
func successor(v *array, key int) (int, bool) {
	var value int
	var truthVal bool
	indx, ok := binSearchAllInstances(v, 0, len(v.a)-1, key)
	if !ok {
		log.Fatal(errors.New("Element Not Found"))
	}
	var maxIdx int
	if len(indx) == 1 {
		 maxIdx = indx[0]
	}else if len(indx) > 1 {
		maxIdx = rSelect(&array{indx}, 0, len(indx)-1, len(indx))
	}
	if maxIdx+1 < len(v.a){
		value = v.a[maxIdx+1]
		truthVal = true
	} else {
		truthVal = false
	}
	return value, truthVal
}


// selectIndx: Given a key, Returns index of that key. 
// If multiple instances are found it returns instance whose index is minimum
func selectIndx(v *array, key int) int {
	indxs, ok := binSearchAllInstances(v, 0, len(v.a)-1, key)
	if !ok {
		log.Fatal(errors.New("Element Not Found"))
	}
	var minIndx int
	if len(indxs) == 1{
		minIndx = indxs[0]
	} else {
		minIndx = rSelect(&array{indxs}, 0, len(indxs)-1, 1)
	}
	return minIndx
}

// Given i for ith smallest element (i.e. order), Returns value of the instance.
func selectIndxBook(v *array, order int) int{
	return v.a[order-1]
}

// Rank Implementation using Binary Search: Returns total no.of elements <= search key (ele)
// (Works for both distinct or repeated elements)
func rank(v *array, key int) int {
	var rankVal int
	indxs, ok := binSearchAllInstances(v, 0, len(v.a)-1, key)
	if !ok {
		log.Fatal(errors.New("Element Not Found"))
	} 
	if len(indxs) == 1 {
		rankVal = indxs[0]+1
	} else {
		rankVal = rSelect(&array{indxs}, 0, len(indxs)-1, len(indxs))+1
	}
	return rankVal
}


// This function finds the ith smallest element of the passed array
// by taking pivot as random index
// This function is a dependency for predecessor(), successor(), selectIndx(), rank() functions
func rSelect(v *array, p, r int, order int) int{
	target := order-1
	if p >= r{
		return v.a[r]
	}
	rand.Seed(time.Now().UnixNano())
	// q is the pivot element
	q := rand.Intn(r-p) + p
	v.a[p], v.a[q] = v.a[q], v.a[p]
	i, j := p+1, p+1
	for j <= r{
		if v.a[j] <= v.a[p]{
			v.a[i], v.a[j] = v.a[j], v.a[i]
			i++
		}
		j++
	}
	v.a[p], v.a[i-1] = v.a[i-1], v.a[p]
	q = i-1
	var targetEle int
	if q == target{
		targetEle = v.a[q]
	} else if target < q {
		targetEle = rSelect(v, p, q-1, order)
	} else { // target > q
		targetEle = rSelect(v, q+1, r, order)
	}
	return targetEle
}
