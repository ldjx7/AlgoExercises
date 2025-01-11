package Interview_coding

import (
	"reflect"
	"sort"
	"testing"
)

// =================================测试题目1=================================
func TestSumArrElements(t *testing.T) {
	SumArrElements(8)
}

// =================================测试题目2=================================
func TestChunk(t *testing.T) {
	// 输出: [[1 2] [3 4] [5]]
	t.Logf("%v", Chunk([]int{1, 2, 3, 4, 5}, 2))

	// 输出: [[1 2 3] [4 5 6] [7 8]]
	t.Logf("%v", Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3))
}

func TestChunkN(t *testing.T) {
	// 输出: [[1 2 3] [4 5]]
	t.Logf("%v", ChunkN([]int{1, 2, 3, 4, 5}, 2))

	// 输出: [[1 2] [3 4] [5]]
	t.Logf("%v", ChunkN([]int{1, 2, 3, 4, 5}, 3))

	// 输出: [[1] [2] [3] [4] [5] [] []]
	t.Logf("%v", ChunkN([]int{1, 2, 3, 4, 5}, 7))
}

// =================================测试题目3=================================
func TestCollectAllDescendantIDs(t *testing.T) {

	// 假设我们有如下树结构：
	//        1
	//      / | \
	//     2  3  4
	//    /|   |
	//   5 6   7

	root := &TreeNode{id: 1, children: []*TreeNode{
		{
			id: 2,
			children: []*TreeNode{
				{id: 5, children: []*TreeNode{}},
				{id: 6, children: []*TreeNode{}},
			},
		},
		{
			id: 3,
			children: []*TreeNode{
				{id: 7, children: []*TreeNode{}},
			},
		},
		{
			id:       4,
			children: []*TreeNode{},
		},
	}}

	result := CollectAllDescendantIDs(root)

	// 预期的结果，这里我们不包含根节点自身的ID
	expected := []int64{2, 5, 6, 3, 7, 4}

	// 因为CollectAllDescendantIDs的实现可能不会保证特定顺序，所以我们应该检查两个集合是否相等
	// 而不是直接比较切片。这里我们简单地对两个切片排序后进行比较。
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CollectAllDescendantIDs() = %v, want %v", result, expected)
	} else {
		t.Log("Test passed!, CollectAllDescendantIDs result:", result)
	}
}

// 测试CollectAllDescendantIDs函数在大型树上的表现
func TestCollectAllDescendantIDSLargeTree(t *testing.T) {
	root, nodeCount := buildLargeTree()

	result := CollectAllDescendantIDs(root)

	// 预期的结果长度
	expectedLength := nodeCount - 1

	// 检查结果长度是否符合预期
	if len(result) != expectedLength {
		t.Errorf("CollectAllDescendantIDs() returned slice of length %d, want %d", len(result), expectedLength)
	} else {
		t.Log("Test passed! The function correctly collected all descendant IDs.")
	}
}

func buildLargeTree() (*TreeNode, int) {
	nodeId := int64(1)
	root := &TreeNode{id: nodeId}
	for i := 0; i < 100; i++ {
		nodeId++
		child1 := &TreeNode{id: nodeId}
		root.children = append(root.children, child1)

		for j := 0; j < 100; j++ {
			nodeId++
			child2 := &TreeNode{id: nodeId}
			child1.children = append(child1.children, child2)

			for k := 0; k < 100; k++ {
				nodeId++
				child3 := &TreeNode{id: nodeId}
				child2.children = append(child2.children, child3)
			}
		}
	}
	return root, int(nodeId)
}

// =================================测试题目4=================================
func TestUpdMap(t *testing.T) {
	UpdMap(10)
}

func TestShardedMap(t *testing.T) {
	UpdSMap(10)
}
