package main

import ("fmt"
        )

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func main(){
  array1 := [...]int {1,2,3,4,5}
  target := 5
  result := twoSum(array1[:], target)
  fmt.Println(result)
}
