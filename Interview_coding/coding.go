package Interview_coding

import (
	"fmt"
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
