package double_point

import (
	"reflect"
	"testing"
)

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//示例 1:
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]

// 示例 2:
// 输入: nums = [0]
// 输出: [0]
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

//示例 1:
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]

// 示例 2:
// 输入: nums = [0]
// 输出: [0]
func TestMoveZeroes(t *testing.T) {
	arr1 := []int{0, 1, 0, 3, 12}
	arr2 := []int{0}
	moveZeroes(arr1)
	moveZeroes(arr2)
	t.Log("arr1 =", arr1)
	t.Log("arr2 =", arr2)
	reflect.DeepEqual(arr1, []int{1, 3, 12, 0, 0})
	reflect.DeepEqual(arr2, []int{0})
}
