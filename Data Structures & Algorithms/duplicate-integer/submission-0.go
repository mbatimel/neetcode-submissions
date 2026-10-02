func hasDuplicate(nums []int) bool {
	res := make(map[int]bool,len(nums))
	for _,n:=range nums{
		if ok:= res[n];ok{
			return true
		}
		res[n]=true
	}
    return  false
}