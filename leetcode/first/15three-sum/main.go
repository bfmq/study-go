package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) (result [][]int) {
	length := len(nums)
	if length < 3 {
		return
	}

	sort.Ints(nums)
	for target := 0; target < length-2; target++ {
		if nums[target] > 0 {
			return
		}
		if target > 0 && nums[target] == nums[target-1] {
			continue
		}

		i := target + 1
		j := length - 1
		for i < j {
			if nums[target]+nums[i]+nums[j] < 0 {
				i++
			} else if nums[target]+nums[i]+nums[j] > 0 {
				j--
			} else {
				result = append(result, []int{nums[target], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			}
		}
	}
	return
}

func main() {
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result)
	result = threeSum([]int{})
	fmt.Println(result)
	result = threeSum([]int{0})
	fmt.Println(result)
	result = threeSum([]int{0, 0, 0, 0})
	fmt.Println(result)
}
