package main

import (
	"fmt"
	"math"
)


type node struct{
	value int
	parent int
	left int
	right int
	size int
}


type nodes struct {
	n []node
}


type array struct {
	a []int
}



func main() {
	st := nodes{}
	// arr := []int{3, 1, 5, 4, 2, 0, 6, -1, 1, 1, 2, 3}
	arr := []int{10, 7, 12, 3, 1, 11, 5, 2, 13, 4, 14, 15, 6, -2, 16, -1, 14, 15, -3, -4, 16, 17, 18}

	// Inserting the element to a search tree
	for _, v := range arr {
		st.insert(v)
	}
	st.show()

	// Searching the element in BST
	indxs, ok := st.search(1, 0)
	if ok {
		fmt.Println("Element found at indexes =", indxs)
	}

	// Minimum value of the BST
	minimum := st.min(0)
	fmt.Println("Minimum value of BST", minimum)

	// Maximum value of the BST
	maximum := st.max(0)
	fmt.Println("Maximum value of BST", maximum)

	// Predecessor of the given element
	pred, ok := st.predecessor(3, 0)
	if ok {
		fmt.Println("Predecessor of the passed element is", pred)
	} else {
		fmt.Println("No predecessor for the passed element")
	}

	// Successor of the given element
	suc, ok := st.successor(6, 6)
	if ok {
		fmt.Println("Successor of the passed element is", suc)
	} else {
		fmt.Println("No Successor for the passed element")
	}

	// Print the elements of BST in sorted order 
	sortArr := st.outputSorted(0)
	fmt.Println("BST as Sorted Array =", sortArr)

	// Select: Find the index of the key passed
	minIdx, ok := st.selectIndx(1)
	if ok{
		fmt.Println("Given key found at index", minIdx)
	} else {
		fmt.Println("Passed key not found in BST")
	}

	// add size to all the nodes
	nodeSize := st.addSize(0)
	fmt.Println("Node size at given index =", nodeSize)
	fmt.Println("After adding Size at each node of subtree of given node index, BST =", st.n)

	// Given i, Find the ith smallest element in BST 
	ithVal := st.selectBookV(6, 1)
	fmt.Println("ith smallest element in subtree with given nodeIdx =", ithVal)


// Case 1: Deleting node with no left or right childs
	st = nodes{}
	fmt.Println("-------------Inserting elements of arr-------------")
	// Inserting the element to a search tree
	for _, v := range arr {
		st.insert(v)
	}
	st.show()
	

	fmt.Println("No left or right childs for ele=4, at index 9")
	// No left or right childs for ele=4, at index 9
	st.delete(4)
	st.show()

// Case 2: Deleting node with only Left child and no right child
	st = nodes{}
	fmt.Println("-------------Inserting elements of arr-------------")
	// Inserting the element to a search tree
	for _, v := range arr {
		st.insert(v)
	}
	st.show()

	fmt.Println("Only Left child and no right child; for ele=7 at index 1")
	// Only Left child and no right child; for ele=7 at index 1	
	st.delete(7)
	st.show()

// Case 3: Deleting node with only right child and no left child
	st = nodes{}
	fmt.Println("-------------Inserting elements of arr-------------")
	// Inserting the element to a search tree
	for _, v := range arr {
		st.insert(v)
	}
	st.show()
	fmt.Println("Only right child and no left child; for ele=13 at index 8")
	// Only right child and no left child; for ele=13 at index 8
	st.delete(13)
	st.show()

// Case 4: Deleting node with both left and right child are available
	st = nodes{}
	fmt.Println("-------------Inserting elements of arr-------------")
	// Inserting the element to a search tree
	for _, v := range arr {
		st.insert(v)
	}
	st.show()
	fmt.Println("Both left and right child available; for ele=3 at index 3")
	// Both left and right child available; for ele=3 at index 3
	st.delete(3)
	st.show()


}


/* 
This method inserts an element into Search Tree as you enter the key.
The first entered element is considered as root node and 
all the next elements are placed appropriately to maintain the property of Search Tree
*/
func (v *nodes) insert(ele int) {
	if len(v.n) == 0 {
		v.n = append(v.n, node{ele, -1, -1, -1, 0})
		return
	}
	i := 0
	for true {
		if ele <= v.n[i].value {
			if v.n[i].left < 0 {
				v.n = append(v.n, node{ele, i, -1, -1, 0})
				v.n[i].left = len(v.n)-1
				break
			} else {
				i = v.n[i].left
				continue
			}
		} else {
			if v.n[i].right < 0 {
				v.n = append(v.n, node{ele, i, -1, -1, 0})
				v.n[i].right = len(v.n)-1
				break
			} else {
				i = v.n[i].right
				continue
			}
		}
	}

}

// This method prints the search tree slice
func (v *nodes) show(){
	fmt.Println(v.n)
}

/*
This method searches for the instances of the given key and returns index of the key if found
If multiple instances are available, it returns all the indices of these muliple instances

To search the whole tree we need to pass key which is the element to search and nodeIdx as 0 and
then it searches the key from root node (whose nodeIdx = 0). If you give a different nodeIdx, then
this function only searches for the passed key in subtree from that particular nodeIdx
*/
func (v *nodes) search(key int, nodeIdx int) ([]int, bool) {
	if nodeIdx < 0{
		return nil, false
	}
	i := nodeIdx
	indxs := []int{}
	for i >= 0 {
		if v.n[i].value == key {
			indxs = append(indxs, i)
			i = v.n[i].left
		} else if key < v.n[i].value {
			i = v.n[i].left
		} else if key > v.n[i].value {
			i = v.n[i].right
		}
	}
	return indxs, true
}

/* 
* If nodeIdx = 0, it returns the minimum element of the whole Search Tree
* If nodeIdx = 5, it returns the minimum element of the subtree from that nodeIdx = 5
*/
func (v *nodes) min(nodeIdx int) int {
	var minimum int
	i := nodeIdx
	for i >= 0{
		if v.n[i].left < 0 {
			minimum = v.n[i].value
		}
		i = v.n[i].left
	}
	return minimum
}


/* 
* If nodeIdx = 0, it returns the maximum element of the whole Search Tree
* If nodeIdx = 5, it returns the maximum element of the subtree from that nodeIdx = 5
*/
func (v *nodes) max(nodeIdx int) int {
	var maximum int
	i := nodeIdx
	for i >= 0{
		if v.n[i].right < 0 {
			maximum = v.n[i].value
		}
		i = v.n[i].right
	}
	return maximum
}

/*
* Returns the predecessor of the given key 'ele'

You need to pass "ele" as key and if nodeIdx = 0, then it searches for the "ele" in search tree using
search() function (this function returns indexes of multiple instances if available). 
Thus when nodeIdx = 0, we will find the predecessor for 
each instance of "ele" (if there are multiple instances) and 
we will return the minimum of these predecessors.

For ex, If ele = 3, nodeIdx = 0, then we will search for ele=3 in search tree.
if 3 is present at index 2, 11. At index 2, the predecessor of 3 is 1 and 
at index 11, the predecessor of 3 is 3. Then we will return minimum of these two predecessors 1 and 3
which is 1. So we will return the predecessor of ele=3, nodeIdx=0 as 1.

What if we give ele = 3, nodeIdx = 11, then we directly find the predecessor based on nodeIdx=11,
without even considering the ele=3, As the predecessor for nodeIdx=11 is 3 (from the above example case)
we will return predecessor of ele=3, nodeIdx=11 as 3
*/
func (v *nodes) predecessor(ele, nodeIdx int) (int, bool) {
	var predVal int
	var ok bool
	var indexes []int
	var test bool
	if nodeIdx == 0{
		indexes, test = v.search(ele, nodeIdx)
		if test {
			var rslt []int
			var temp int
			for _, idx := range indexes{
				if idx == 0 {
					temp = v.max(v.n[idx].left)
					rslt = append(rslt, temp)
					continue
				}
				temp, test = v.predecessor(ele, idx)
				if test {
					rslt = append(rslt, temp)
				}
			}
			predVal = minimum(rslt)
			if predVal != math.MaxInt64 {
				ok = true
			}
			return predVal, ok
		}
	}
	i := nodeIdx
	if v.n[i].left >= 0 {
		predVal = v.max(v.n[i].left)
		ok = true
	} else {
		parent := v.n[i].parent
		if parent >= 0 {
			parentsParent := v.n[parent].parent
			if parentsParent >= 0 && v.n[parent].left == i && parent == v.n[parentsParent].right{
				predVal = v.n[parentsParent].value
				ok = true
			} else if v.n[parent].right == i {
				predVal = v.n[parent].value
				ok = true
			}
		}
	}
	return predVal, ok
}

// This function is dependency for predecessor() method and returns the minimum value of the given slice 
// if the given slice is empty it will return math.MaxInt64 as minimum value. So we have to deal with
// this where we used this function to return the minimum value
func minimum(arr []int) int {
	minVal := math.MaxInt64
	for _, v := range arr{
		if v < minVal{
			minVal = v
		}
	}
	return minVal
}

/*
* Returns the successor of the given key 'ele'

You need to pass "ele" as key and if nodeIdx = 0, then it searches for the "ele" in search tree using
search() function (this function returns indexes of multiple instances if available). 
Thus when nodeIdx = 0, we will find the successor for 
each instance of "ele" (if there are multiple instances) and 
we will return the maximum of these successors.

For ex, If ele = 4, nodeIdx = 0, then we will search for ele=4 in search tree.
if 4 is present at index 5, 12. At index 5, the successor of 4 is 6 and 
at index 12, the successor of 4 is 4. Then we will return maximum of these two successors 6 and 4
which is 6. So we will return the successor of ele=4, nodeIdx=0 as 6.

What if we give ele = 4, nodeIdx = 12, then we directly find the successor based on nodeIdx=12,
without even considering the ele=4, As the successor for nodeIdx=12 is 4 (from the above example case)
we will return predecessor of ele=4, nodeIdx=12 as 4
*/
func (v *nodes) successor(ele, nodeIdx int) (int, bool) {
	var sucVal int
	var ok bool
	var indexes []int
	var test bool
	if nodeIdx == 0 {
		indexes, test = v.search(ele, 0)
		if test{
			var rslt []int
			var temp int
			for _, idx := range indexes {
				if idx == 0{
					temp = v.min(v.n[idx].right)
					rslt = append(rslt, temp)
					continue
				}
				temp, test = v.successor(ele, idx)
				if test{
					rslt = append(rslt, temp)
				}
				sucVal = maximum(rslt)
				if sucVal != math.MinInt64 {
					ok = true
				}
				return sucVal, ok
			}
		}
	}

	i := nodeIdx
	if v.n[i].right >= 0 {
		sucVal = v.min(v.n[i].right)
		ok = true
	} else {
		parent := v.n[i].parent
		if parent >= 0 {
			parentsParent := v.n[parent].parent
			if parentsParent >= 0 && v.n[parent].right == i && v.n[parentsParent].left == parent{
				sucVal = v.n[parentsParent].value
				ok = true
			} else if v.n[parent].left == i{
				sucVal = v.n[parent].value
				ok = true
			}
		}
	}
	return sucVal, ok
}


// This function is dependency for successor() method and returns the maximum value of the given slice 
// if the given slice is empty it will return math.MinInt64 as maximum value. So we have to deal with
// this where we used this function to return the maximum value
func maximum(arr []int) int {
	maxVal := math.MinInt64
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}


/*
If nodeIdx = 0, returns all the elements of Search Tree in a sorted order as a slice
If nodeIdx = 5, returns all the elements of subtree from node 5 in a sorted order as a slice
*/
var sortedArray1 []int
func (v *nodes) outputSorted(nodeIdx int) []int{
	if nodeIdx < 0 {
		return []int{}
	}
	i := nodeIdx
	_ = v.outputSorted(v.n[i].left)
	nodeVal := v.n[i].value
	sortedArray1 = append(sortedArray1, nodeVal)
	_ = v.outputSorted(v.n[i].right)
	return sortedArray1
}


// Select: Given a key, Returns index of that key. If multiple instances found return smallest index
// Along with index it also returns boolean true or false, false indicating no such key is found
// If true, means we found the key and the returned index is of that key.
func (v *nodes) selectIndx(key int) (int, bool) {
	indexes, ok := v.search(key, 0)
	var test bool
	var minIdx int
	if ok {
		minIdx = minimum(indexes)
	} else {
		return minIdx, false
	}
	if minIdx != math.MaxInt64 {
		test = true
	}
	return minIdx, test
}

/*
This method appends size of each node to their respective subarray.
size indicates the size of the node meaning how many elements are there in the subtree 
starting from node value 1 (inclduing itself)

For ex, if node 3 has 1 left child and no right child; size = 1 left child + 1 (node 3) + 0 right child
*/
func (v *nodes) addSize(nodeIdx int) int {
	if nodeIdx < 0 {
		return 0
	}
	i := nodeIdx
	left := v.addSize(v.n[i].left)
	nodeItem := 1
	right := v.addSize(v.n[i].right)
	v.n[i].size = left+nodeItem+right
	return v.n[i].size
}


/* Select: Given i for ith Smallest element (i.e. order), Returns value of the instance
If nodeIdx = 0 the ith Smallest is from the whole search tree;
If nodeIdx = some index (for ex, = 4), we will find the ith smallest of the subtree from nodeIdx = 4
*/
func (v *nodes) selectBookV(order, nodeIdx int) (int) {
	if order >= len(v.n) {
		panic("Order not with in the range of the subtree with given nodeIdx")
	}
	if nodeIdx < 0{
		panic("Order not with in the range of the subtree with given nodeIdx")
	}
	i := nodeIdx
	var j int
	leftIdx := v.n[i].left
	if leftIdx >= 0{
		j = v.n[leftIdx].size
	} else {
		j = 0
	}
	if order == j+1 {
		return v.n[i].value
	} else if order < j+1 {
		return v.selectBookV(order, leftIdx)
	} else if order > j+1 {
		return v.selectBookV(order-j-1, v.n[i].right)
	}
	return 0

}

/*
* This method is useful for modifying the elements and their parent, left and right child indices
  so that we can delete that particular element which wants to be deleted without any future dependencies 
  on this element
  Thus, for removing these dependencies on deleting element we need to perform modifications
*/
func (v *nodes) modificationsForDeletion(index, nodeIdx int) {
	i :=  nodeIdx
	if i < 0 {
		return
	}
	v.modificationsForDeletion(index, v.n[i].left)
	v.modificationsForDeletion(index, v.n[i].right)
	if i > index {
		if v.n[i].parent > index {
			v.n[i].parent -= 1
		}
		if v.n[i].left >= 0 {
			v.n[i].left -= 1
		}
		if v.n[i].right >= 0 {
			v.n[i].right -= 1
		}
	} else if i < index {
		if v.n[i].left >= 0 && v.n[i].left > index{
			v.n[i].left -= 1
		}
		if v.n[i].right >= 0 && v.n[i].right > index{
			v.n[i].right -= 1
		}
	}
}


/*
* Here, this method accepts variadic parameters (vals ...int). 
* If len(vals) = 0 i.e. no parameter is passed, we just return
* If len(vals) == 1, i.e. only 1 parameter is passed we consider this as key to be deleted. 
	Deletes the element vals[0]; If multiple instances of vals[0] is present we will delete the element
	whose index is maximum.
* If len(vals) == 2, i.e. 2 parameters are passed, then we consider vals[1] as node Index and to be 
  assumed that we want to delete that particular node. So we neglect the key which is given as vals[0].
  Thus vals[0] is of no use if we pass node index already and that particular node would be deleted 
  irrespective of key passed.
	Even if we just want to delete any particular node by giving node index. For code readability purposes
	Generally, it is better to pass value of that node as vals[0] but not any random element with
	the thought of as if we are going to neglect it anyways and will consider only node index  
*/
func (v *nodes) delete(vals ...int) {
	if len(vals) == 0 {
		return
	}
	var idx int
	if len(vals) == 1 {
		indexes, ok := v.search(vals[0], 0)
		if ok {
			maxVal := maximum(indexes)
			if maxVal != math.MinInt64 {
				idx = maxVal
			}
		}
	} else if len(vals) == 2{
		idx = vals[1]
	} else {
		panic("More than required no.of parameters has been passed")
	}
	parent := v.n[idx].parent
	left := v.n[idx].left
	right := v.n[idx].right
	if left < 0 && right < 0 {
		if v.n[idx].value > v.n[parent].value { // If it's right node
			v.n[parent].right = -1
		} else { // if it's left node
			v.n[parent].left = -1
		}
		v.modificationsForDeletion(idx, 0)
		v.n = append(v.n[:idx], v.n[idx+1:]...)
	} else if left >= 0 && right < 0 {
		v.n[left].parent = v.n[idx].parent
		if v.n[parent].left == idx {
			v.n[parent].left = v.n[idx].left
		} else if v.n[parent].right == idx {
			v.n[parent].right = v.n[idx].left
		}
		v.modificationsForDeletion(idx, 0)
		v.n = append(v.n[:idx], v.n[idx+1:]...)
	} else if left < 0 && right >= 0 {
		v.n[right].parent = v.n[idx].parent
		if v.n[parent].left == idx {
			v.n[parent].left = v.n[idx].right
		} else if v.n[parent].right == idx {
			v.n[parent].right = v.n[idx].right
		}
		v.modificationsForDeletion(idx, 0)
		v.n = append(v.n[:idx], v.n[idx+1:]...)
	} else { // both left node and right node are present i.e. >= 0
		maximumIdx := v.maxIdx(left)
		v.n[idx].value, v.n[maximumIdx].value = v.n[maximumIdx].value, v.n[idx].value
		v.delete(v.n[maximumIdx].value, maximumIdx)
	}
}

/* 
* This method is a dependency for delete() method.
* If nodeIdx = 0, it returns the index of the maximum element of the whole Search Tree
* If nodeIdx = 5, it returns the index of the maximum element of the subtree from that nodeIdx = 5
*/
func (v *nodes) maxIdx(nodeIdx int) int {
	i := nodeIdx
	if v.n[i].right < 0 {
		return i
	}
	return v.maxIdx(v.n[i].right)
}