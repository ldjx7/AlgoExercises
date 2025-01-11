package double_point

import (
	"sort"
	"testing"
)

//15. 三数之和

//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//注意：答案中不可以包含重复的三元组。

//示例 1：
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//示例 2：
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//示例 3：
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
//提示：
//3 <= nums.length <= 3000
//-105 <= nums[i] <= 105

func threeSum(nums []int) [][]int {
	// return threeSumSolutionV1(nums) // 卡在第308个用例 执行超时
	return threeSumSolutionV2(nums)
}

func threeSumSolutionV1(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	duplicate := make(map[[3]int]bool)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmp := [3]int{nums[i], nums[j], nums[k]}
					// 去重. 数组和切片共享底层,切片的排序会使数组顺序也改变
					sort.Ints(tmp[:])
					if duplicate[tmp] {
						continue
					} else {
						result = append(result, tmp[:])
						duplicate[tmp] = true
					}
				}
			}
		}
	}
	return result
}

// 首先对数组进行排序，排序后固定一个数 nums[i]，再使用左右指针指向 nums[i]后面的两端，数字分别为 nums[L] 和 nums[R]，计算三个数的和 sum 判断是否满足为 0，满足则添加进结果集
// 如果 nums[i]大于 0，则三数之和必然无法等于 0，结束循环
// 如果 nums[i] == nums[i−1]，则说明该数字重复，会导致结果重复，所以应该跳过
// 当 sum == 0 时，nums[L] == nums[L+1] 则会导致结果重复，应该跳过，L++
// 当 sum == 0 时，nums[R] == nums[R−1] 则会导致结果重复，应该跳过，R−−
func threeSumSolutionV2(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		// 如果 nums[i]大于 0，则三数之和必然无法等于 0
		if nums[i] > 0 {
			break
		}
		// 如果 nums[i] == nums[i−1]，则说明该数字重复，会导致结果重复，所以应该跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}

	}
	return result
}

// hash解法
func threeSumSolutionV3(nums []int) [][]int {
	result := make([][]int, 0)

	//

	return result
}

func TestThreeSum(t *testing.T) {
	example1 := []int{-1, 0, 1, 2, -1, -4}
	example2 := []int{0, 1, 1}
	example3 := []int{0, 0, 0}

	t.Log("示例1: ", threeSum(example1))
	t.Log("示例2: ", threeSum(example2))
	t.Log("示例3: ", threeSum(example3))
}
