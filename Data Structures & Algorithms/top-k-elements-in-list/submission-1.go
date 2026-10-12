func topKFrequent(nums []int, k int) []int {
res := make([]int, 0)
	preRes := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		preRes[nums[i]] += 1
	}
	for i := range preRes {
		res = append(res, i)
	}
	sort.SliceStable(res, func(i, j int) bool {
		return preRes[res[i]] > preRes[res[j]]
	})

	return res[:k]
}