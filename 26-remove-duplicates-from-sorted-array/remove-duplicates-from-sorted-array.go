func removeDuplicates(nums []int) int {
	numExist := map[int]bool{}
	k := 0

	for _, num := range nums {
		if !numExist[num] {
			nums[k] = num
			k++
			numExist[num] = true
		}
	}
	return k
}
