//https://leetcode.com/problems/move-zeroes/description/
func moveZeroes(nums []int) {
outer:
	for i := 0; i < len(nums); i++ {

		for nums[i] != 0 {
			i++
			if i == len(nums) {
				break outer
			}
		}
		j := i
		for nums[j] == 0 {
			j++
			if j == len(nums) {
				break outer
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}