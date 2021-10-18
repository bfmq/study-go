package main

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 0})
	moveZeroes([]int{})
}
