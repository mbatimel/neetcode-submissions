func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i==j{
				continue
			}
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				break
			}

		}
	}
	return res
}