func maxArea(heights []int) int {
	k := 0
	l := 0
	r := len(heights) - 1
	for l < r {
		sum := (r - l) * int(math.Min(float64(heights[l]), float64(heights[r])))
		if sum > k {
			k = sum
		}
		if heights[l] < heights[r] {
			l += 1
		}else {
			r -= 1
		}

	}
	return k
}