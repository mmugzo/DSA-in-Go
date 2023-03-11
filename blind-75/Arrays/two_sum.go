package main

import "fmt"

func twoSum(nums []int, target int) []int {
	
	checked := make(map[int]int)
	var ans []int

	for idx, num := range nums {
		diff := target - num

		if c, ok := checked[diff]; ok {
			ans = []int{c, idx}
			break
		}
		checked[num] = idx
	}
	return ans

}
 

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums,target))
}



