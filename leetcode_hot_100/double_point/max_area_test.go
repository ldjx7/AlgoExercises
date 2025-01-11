package double_point

import "testing"

// 11. 盛最多水的容器
// 中等
// 相关标签
// 相关企业
// 提示
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例 2：
// 输入：height = [1,1]
// 输出：1
func maxArea(height []int) int {
	//return maxAreaSolutionV1(height)
	return maxAreaSolutionV2(height)
}

// 暴力破解
func maxAreaSolutionV1(height []int) int {
	toMaxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > toMaxArea {
				toMaxArea = area
			}
		}
	}
	return toMaxArea
}

// 双指针
// 双指针的核心点在于,面积始终是受最短的边影响,从左右两边同时移动指针
// 当计算出面积后,移动较短的那一边来判断面积是否增大,以此类推
func maxAreaSolutionV2(height []int) int {
	left, right := 0, len(height)-1
	toMaxArea := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > toMaxArea {
			toMaxArea = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return toMaxArea
}

func TestMaxArea(t *testing.T) {
	example1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	example2 := []int{1, 1}
	t.Log("示例1结果:", maxArea(example1))
	t.Log("示例2结果:", maxArea(example2))
}
