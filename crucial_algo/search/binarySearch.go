package main

import (
	"fmt"
	"strconv"
)

func search(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right - left) /2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid -1
		} else {
			left = mid + 1
		}

	}
	return -1
	
}

func main() {
	arr := []int{1,2,3,4,5,6,7}
	target := 4
	result := search(arr, target)
	if result != -1 {
		fmt.Println("target found at index " + strconv.Itoa(result))
	} else {
		fmt.Println("target not found")
	}
}
