package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 7, 11, 15}
	target := 18

	res := twoSum(arr, target)
	fmt.Println(res)
}

// twoSum returns the indexes of the first two values that add up to the target in a list
// nums = [2, 7, 11, 15], target = 9, solution: [0,1]
// we can make this very efficient by storing a map of complements. A complement is the amount you must add to something to make it whole.
// so for each item in the list, we substract it from the target, e.g.
/*
	for this case [2, 7, 11, 15], target: 18, first iteration, 2-18 = 16 => 16 is the value we must add to '2' to get to 18
	so we store this in a map {16:0}, meaning that if we encounter 16 at some point in the list, we know it's complement its at index 0
	so we can return the current index, and the stored index in the map.
	For this example it would go like so:
	check if the current number's complement exists in teh map, is so return indexes, otherwise, calculate the complement and store it in map
	2 - 18 = 16 => map{16:0}
	7 - 18 = 11 => map{ 16:0, 11:1 }
	11 exists in map, return
*/
// with this, we can see if the current's numbers complement exists in the map, if so, return current index, plus index in map
// else, we add this unseen complement to the map and continue.
func twoSum(nums []int, target int) []int {
	compMap := make(map[int]int)

	for i, num := range nums {
		idx, ok := compMap[num]
		if ok {
			return []int{idx, i}
		}

		complement := target - num
		compMap[complement] = i
		fmt.Println("complement map:", compMap)
	}

	return []int{}
}

// TwoSum given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Given nums = [2, 7, 11, 15], target = 9,
// iterate through the list of numbers, find its complement, by substracting the num from the target,
// store the complement's index location (num required to add to get the target) it in a k/v store
// for each value we iterate through, check if it's complement is in the map, if so, return both indexes
func TwoSum(nums []int, target int) []int {
	mapStore := make(map[int]int)
	for i, num := range nums {
		if idx, ok := mapStore[num]; ok {
			return []int{idx, i}
		}
		mapStore[target-num] = i
	}

	return nil
}
