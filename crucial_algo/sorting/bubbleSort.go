package main

import "fmt"


func b_sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main(){
	arr := []int{5,2,1,1,7,6,6}
	fmt.Println(b_sort(arr))

}