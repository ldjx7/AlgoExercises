package hash

import "testing"

// 示例 1：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
// 示例 2：
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
// 示例 3：
// 输入：nums = [3,3], target = 6
// 输出：[0,1]
func TestTwoSum(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{3, 2, 4}
	nums3 := []int{3, 3}
	t.Log(twoSum(nums1, 9))
	t.Log(twoSum(nums2, 6))
	t.Log(twoSum(nums3, 6))
}
