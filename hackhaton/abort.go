package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums[len(nums)/2]
}
