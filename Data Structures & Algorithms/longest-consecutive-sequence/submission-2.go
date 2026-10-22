func longestConsecutive(nums []int) int {
	if len(nums)==0{
		return 0
	}
		quickSort(nums, 0, len(nums)-1)
	k := 1
	res := 1
	for i := 1; i < len(nums); i++ {
		if res < k {
			res = k
		}
		if nums[i]-nums[i-1] == 1 {
			k++
		} else if nums[i] != nums[i-1] {
			k = 1
		}
	}
	if res < k {
		res = k
	}
	return res
}
func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}