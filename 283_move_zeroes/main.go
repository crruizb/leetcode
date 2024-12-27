package main

// https://leetcode.com/problems/move-zeroes/?envType=study-plan-v2&envId=leetcode-75
func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	for n := range arr {
		println(arr[n])
	}
}

func moveZeroes(nums []int) {
	i, j := 0, 0
	for n := range nums {
		if nums[n] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
