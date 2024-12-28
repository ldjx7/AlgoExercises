package hash

// TwoSum 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
// 你可以按任意顺序返回答案
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
func twoSum(nums []int, target int) []int {
	//return twoSumSolutionV1(nums, target)
	return twoSumSolutionV2(nums, target)
}

// 暴力破解 复杂度 O(n^2)
func twoSumSolutionV1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 哈希映射
// 这道题本身如果通过暴力遍历的话也是很容易解决的，时间复杂度在 O(n2)
// 由于哈希查找的时间复杂度为 O(1)，所以可以利用哈希容器 map 降低时间复杂度
// 遍历数组 nums，i 为当前下标，每个值都判断map中是否存在 target-nums[i] 的 key 值
// 如果存在则找到了两个值，如果不存在则将当前的 (nums[i],i) 存入 map 中，继续遍历直到找到为止
func twoSumSolutionV2(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{m[target-num], index}
		}
		m[num] = index
	}
	return nil
}
