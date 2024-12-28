package hash

import "testing"

func TestTwoSum(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{3, 2, 4}
	nums3 := []int{3, 3}
	t.Log(twoSum(nums1, 9))
	t.Log(twoSum(nums2, 6))
	t.Log(twoSum(nums3, 6))
}
