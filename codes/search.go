package main

import (
	"fmt"
	"sort"
)

func seqSearch(items []int, elem int) (idx int) {
	for idx, item := range items {
		if elem == item {
			return idx
		}
	}
	return -1
}

func binSearch(items []int, elem int) int {
	sort.Ints(items)

	start, end := 0, len(items)-1

	for start <= end {
		mid := start + end/2

		if elem > items[mid] {
			start = mid + 1
		} else if elem < items[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {

	arr := []int{1, 19, 2, 4, 59, 21, 88, 41, 9}
	idx1 := seqSearch(arr, 100)
	fmt.Println("seq search:", idx1)
	idx2 := binSearch(arr, 19)
	fmt.Println("bin search:", idx2)

}
