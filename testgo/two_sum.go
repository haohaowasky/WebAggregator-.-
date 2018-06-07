package main

import("fmt"
        )


func twoSum(nums []int, target int) []int {
    check:= make(map[int]int)
    for i,j := range(nums) {
        middleman := target - j
        if _, ok := check[middleman]; ok{
            return []int {check[middleman],i}
        }
        check[j] = i
    }
    return nil
}

func main(){
  r := [...]int {1,2,3,4,5}
  j:= twoSum(r[:],5)
  fmt.Println(j)
}
