package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var value_idx_map map[int]int
	value_idx_map = make(map[int]int)
	for idx, value := range nums {
		if v, ok := value_idx_map[value]; ok {
			if value == target/2 {
				return []int{v, idx}
			}
		} else {
			value_idx_map[value] = idx
		}

	}

	for k, v := range value_idx_map {
		other_key := target - k
		if other_key == k {
			continue
		}
		if other_idx, ok := value_idx_map[other_key]; ok {
			return []int{v, other_idx}
		}
	}
	return make([]int,2)
}

func main() {
	nums := []int{3,3}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}
