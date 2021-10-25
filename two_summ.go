package main

func twoSum(nums []int, target int) []int {
	if len(nums) == 1 {
		return []int{}
	}
	if len(nums) == 2 {
		return []int{0, 1}
	}
	for i := 0; i < len(nums); i++ {
		for g := i + 1; g < len(nums); g++ {
			if nums[i]+nums[g] == target {
				return []int{i, g}
			}
		}
	}
	return []int{}
}
