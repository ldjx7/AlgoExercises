package Interview_coding

import (
	"fmt"
	"strconv"
	"sync"
)

// =================================题目一=================================
// 写一个go程序,使用goroutine和channel计算出一个数据中的所有元素之和
func sum(arr []int, ch chan int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		recover()
	}()
	res := 0
	for _, value := range arr {
		res += value
	}
	ch <- res
}

func buildIntArr(length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = i
	}
	return result
}

// SumArrElements 考察点主要是协程和通道的使用
func SumArrElements(length int) {
	arr := buildIntArr(length)
	fmt.Println("构建的数组为: ", arr)
	n := 2 // 将数组分成两部分
	ch := make(chan int, n)
	var wg sync.WaitGroup
	// 划分数组为两部分，并使用 goroutine 计算每部分的和
	wg.Add(n)
	go sum(arr[:len(arr)/n], ch, &wg)
	go sum(arr[len(arr)/n:], ch, &wg)

	// 等待所有 goroutine 完成
	wg.Wait()
	close(ch)

	// 汇总所有部分的和
	totalSum := 0
	for s := range ch {
		totalSum += s
	}
	fmt.Println("数组元素的总和是:", totalSum)
}

// =================================题目二=================================

// Chunk 返回一个数组，其中的元素按size的长度分成若干组。如果数组不能均匀分割，最后的块将是剩余的元素。
func Chunk[T any, Slice ~[]T](collection Slice, size int) []Slice {
	length := len(collection)
	groupNum := length / size

	if length%size != 0 {
		groupNum++
	}
	result := make([]Slice, 0, groupNum)
	for i := 0; i < groupNum; i++ {
		end := (i + 1) * size
		if end > length {
			end = length
		}
		result = append(result, collection[i*size:end])
	}
	return result
}

// ChunkN 是一个泛型函数，接收一个列表 list 和一个正整数 n。
// 目标是将给定的列表均匀地分割成 n 个子列表，每个子列表的元素数量尽可能相等。
// 如果不能完全平均分配，则允许最左边的几个子列表比右边的多一个元素。
//
// 特殊情形处理：
//   - 如果 n 大于列表 list 的长度，那么返回的结果应包含 len(list) 个含有单个元素的子列表，
//     以及若干个空子列表（即长度为0的子列表），以确保总共返回 n 个子列表。
//
// 示例：
// ChunkN([]int{1, 2, 3, 4, 5}, 2) 返回 [][]int{{1, 2, 3}, {4, 5}}
// ChunkN([]int{1, 2, 3, 4, 5}, 3) 返回 [][]int{{1, 2}, {3, 4}, {5}}
// ChunkN([]int{1, 2, 3, 4, 5}, 7) 返回 [][]int{{1}, {2}, {3}, {4}, {5}, {}, {}}
func ChunkN[T any](list []T, n int) [][]T {

	result := make([][]T, n)

	length := len(list)
	if n > length {
		for i := 0; i < n; i++ {
			if i < length {
				result[i] = []T{list[i]}
				continue
			}
			result[i] = []T{}
		}
		return result
	}

	elementNum := length / n
	if length%n != 0 {
		elementNum++
	}
	for i := 0; i < n; i++ {
		end := (i + 1) * elementNum
		if end > length {
			end = length
		}
		result[i] = append(result[i], list[i*elementNum:end]...)
	}

	return result
}

// =================================题目三=================================

type TreeNode struct {
	id       int64
	children []*TreeNode // 不可直接访问，需要通过 GetChildren() 方法访问
}

// GetChildren 计算密集型操作
func (n *TreeNode) GetChildren() []*TreeNode {
	performComputation() // 模拟一些计算操作
	return n.children
}

func performComputation() {
	x := 0
	for i := 0; i < 10000; i++ {
		x += i
	}
}

// CollectAllDescendantIDs 收集所有的子节点
// TODO 实现这个函数
func CollectAllDescendantIDs(root *TreeNode) []int64 {
	result := make([]int64, 0)
	getAllChildID(&result, root)
	return result

}

func getAllChildID(ids *[]int64, node *TreeNode) {
	if node == nil {
		return
	}

	children := node.GetChildren()
	if len(children) == 0 {
		return
	}

	for _, child := range children {
		*ids = append(*ids, child.id)
		getAllChildID(ids, child)
		//ids = append(ids, getAllChildID([]int64{}, child)...)
	}
}

// =================================题目四=================================
// 并发更新一个map，map是string->int的结构，使所有的key对应的value都+1
// 所有的都需要自己实现
// 1.不要写死map数据，先写一个构建map的函数，便于在测试大map时更方便
// 2.并发更新map，使所有key对应的value值都+1
// 题目分析：map是并发不安全的，更新时需要对其 加锁/解锁 操作
// 题目疑问：由于map并发不安全，并发时对其进行加锁，还是相当于同步操作？这样有什么意义呢？
// 面试官回答：这里是模拟操作，value+1只是一个非常简单的场景，实际场景可能非常耗时，计算过程不需要加锁，只在更新map的这一步操作加锁

var rw = sync.RWMutex{}

func buildMap(length int) map[string]int {
	m := map[string]int{}
	for i := 0; i < length; i++ {
		m[strconv.Itoa(i)] = i
	}
	return m
}

func DoSomething(m map[string]int, key string, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.Lock()
	defer rw.Unlock()
	m[key] += 1
}

func UpdMap(mapLen int) {
	m1 := buildMap(mapLen)
	fmt.Println("build map: ", m1)
	wg := sync.WaitGroup{}

	rw.RLock()
	// 由于map并发不安全，不能同时读写，这里先加锁将所有的key读取出来
	keys := make([]string, 0, len(m1))
	for k := range m1 {
		keys = append(keys, k)
	}
	rw.RUnlock() // 解锁，允许写入操作

	for _, k := range keys {
		wg.Add(1)
		go DoSomething(m1, k, &wg)
	}
	wg.Wait()
	fmt.Println("update after map:", m1)
}

// 解法二（GPT答案）：采用分片的方式
const shardCount = 16

type ShardedMap struct {
	shards [shardCount]map[string]int
	locks  [shardCount]sync.Mutex
}

// NewShardedMap 初始化分片
func NewShardedMap(length int) *ShardedMap {
	sm := &ShardedMap{}
	for i := 0; i < shardCount; i++ {
		sm.shards[i] = make(map[string]int, length/shardCount)
	}
	return sm
}

// 获取分片索引
func (sm *ShardedMap) getShard(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])) % shardCount
	}
	return hash
}

// Update 更新键值
func (sm *ShardedMap) Update(key string, delta int) {
	shardIndex := sm.getShard(key)
	sm.locks[shardIndex].Lock()
	defer sm.locks[shardIndex].Unlock()
	sm.shards[shardIndex][key] += delta
}

// Flatten 获取完整的 map
func (sm *ShardedMap) Flatten() map[string]int {
	result := make(map[string]int)
	for i := 0; i < shardCount; i++ {
		sm.locks[i].Lock()
		for k, v := range sm.shards[i] {
			result[k] = v
		}
		sm.locks[i].Unlock()
	}
	return result
}

// 构建测试数据
func buildSMap(length int) *ShardedMap {
	sm := NewShardedMap(length)
	for i := 0; i < length; i++ {
		key := strconv.Itoa(i)
		sm.Update(key, i)
	}
	return sm
}

func UpdSMap(mapLen int) {
	sm := buildSMap(mapLen)
	fmt.Println("build map: ", sm.Flatten())
	wg := sync.WaitGroup{}

	for i := 0; i < mapLen; i++ {
		key := strconv.Itoa(i)
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			sm.Update(k, 1)
		}(key)
	}

	wg.Wait()
	result := sm.Flatten()
	fmt.Println("update after map:", result)
}
