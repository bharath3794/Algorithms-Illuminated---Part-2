package main

import (
	"fmt"
)

type array struct{
	a []int
}

func main() {
	// Peforming Push, Pop, Delete operations on heap
	fmt.Println("-----Performing Heap Operations-----")
	arr := []int{1, 6, 2, 7, 8, 3}
	heap := array{}
	for _, val := range arr {
		heap.heapPush(val)
	}
	fmt.Println("After Pushing, heap =", heap)
	fmt.Println("Deleting value 7 from heap =", heap)
	heap.heapDelete(7)
	fmt.Println("After Deleting value 7, heap =", heap)
	fmt.Println("Performing Pop(), Popped Element =", heap.heapPop())
	fmt.Println("After Pop, heap =", heap)

	fmt.Println("-----Performing Heap Sort and Median Maintaineance Applications using Heap-----")
	// Heap Sort and Median Maintaineance Problem
	arr = []int{59, 87, 39, 77, 59, 91, 42, 99, 83, 38, 98, 91}
	fmt.Println("arr =", arr)
	sortedArr := heapSort(arr)
	fmt.Println("sortedArr =", sortedArr)
	mediansMaintainArr := medianMaintain(arr)
	fmt.Println("mediansMaintainArr =", mediansMaintainArr)
	
}


/*
* heapSort sorts the given array of elements by first inserting them to heap using
  heapPush() method and pop out minimum ele each time using heapPop() method
* To insert array elements into heap Complexity = 𝑂(𝑛); 
* To pop one element = 𝑂(𝑙𝑜𝑔(𝑛)); For n elements = 𝑂(𝑛𝑙𝑜𝑔(𝑛))
* Total Complexity = 𝑂(𝑛) + 𝑂(𝑛𝑙𝑜𝑔(𝑛)); We can neglect 𝑂(𝑛) as 𝑂(𝑛) < 𝑂(𝑛𝑙𝑜𝑔(𝑛));
* So, Overall Algorithm Complexity: 𝑂(𝑛𝑙𝑜𝑔(𝑛))
*/
func heapSort(v []int) array{
	heap := array{}
	for _, val := range v{
		heap.heapPush(val)
	}
	heapLngth := len(heap.a)
	sortedArr := array{}
	for i:=0; i<heapLngth; i++ {
		sortedArr.a = append(sortedArr.a, heap.heapPop())
	}
	return sortedArr
}


/*
* heapPush() puts the given ele into heap by maintaining the heap's property
* Algorithm Complexity: 𝑂(𝑙𝑜𝑔(𝑛))
*/
func (v *array) heapPush(ele int) {
	v.a = append(v.a, ele)
	i := len(v.a)-1
	for i >= 1 && v.a[i] < v.a[(i-1)/2]{
		v.a[i], v.a[(i-1)/2] = v.a[(i-1)/2], v.a[i]
		i = (i-1)/2
	}
}

/*
* heapPop() pops out the minimum ele from heap and also maintain the heap's property while popping out.
* Algorithm Complexity: 𝑂(𝑙𝑜𝑔(𝑛))
*/
func (v *array) heapPop() int{
	lastIdx := len(v.a)-1
	v.a[0], v.a[lastIdx] = v.a[lastIdx], v.a[0]
	ele := v.a[lastIdx]
	v.a = v.a[:lastIdx]
	i := 0
	for true {
		var j int
		if (2*i+2) < len(v.a){
			if v.a[2*i+1] <= v.a[2*i+2]{
				j = 2*i+1
			} else {
				j = 2*i+2
			}
		} else if (2*i+1) < len(v.a) {
			j = 2*i+1
		} else {
			break
		}
		if v.a[i] > v.a[j]{
			v.a[i], v.a[j] = v.a[j], v.a[i]
			i = j
		} else {
			break
		}
	}
	return ele
}


func (v *array) heapDelete(key int) {
	var i int
	for k, ele := range v.a {
		if ele == key {
			i = k
			break
		}
	}
	lastIdx := len(v.a)-1
	v.a[i], v.a[lastIdx] = v.a[lastIdx], v.a[i]
	v.a = v.a[:lastIdx]
	if (i-1)/2 > 0 && v.a[i] < v.a[(i-1)/2] {
		for (i-1)/2 > 0 {
			if v.a[i] < v.a[(i-1)/2] {
				v.a[i], v.a[(i-1)/2] = v.a[(i-1)/2], v.a[i]
				i = (i-1)/2
			} else {
				break
			}
		}
	} else {
		for true {
			var j int
			if (2*i+2) < len(v.a){
				if v.a[2*i+1] <= v.a[2*i+2]{
					j = 2*i+1
				} else {
					j = 2*i+2
				}
			} else if (2*i+1) < len(v.a) {
				j = 2*i+1
			} else {
				break
			}
			if v.a[i] > v.a[j]{
				v.a[i], v.a[j] = v.a[j], v.a[i]
				i = j
			} else {
				break
			}
		}
	}

}


/*
* Application: Median Maintenance using Heap [ 𝑂(𝑙𝑜𝑔(𝑛) ] for each element in the array
Here we considered median as: If there are 2k-1 elements (Odd length) median = k^{th} element 
and if there are 2k elements (even length) median = k^{th} element 
i.e. lower order term of both k^{th} element and (k+1)^{th} element
*/
func medianMaintain(v []int) []int {
	leftHeap := array{}
	var median int
	rightHeap := array{}
	median = v[0]
	mediansArr := []int{median}
	for i:=1; i<len(v); i++ {
		if v[i] <= median {
			leftHeap.heapPush(-v[i])
		} else {
			rightHeap.heapPush(v[i])
		}
		if len(leftHeap.a)+1 < len(rightHeap.a){
			leftHeap.heapPush(-median)
			median = rightHeap.heapPop()
		} else if len(leftHeap.a) > len(rightHeap.a){
			rightHeap.heapPush(median)
			median = -leftHeap.heapPop()
		}
		mediansArr = append(mediansArr, median)
	}
	return mediansArr
}