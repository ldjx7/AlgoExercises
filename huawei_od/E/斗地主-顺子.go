package E

import (
	"strconv"
	"strings"
)

// FindSequences 斗地主-顺子
// 题目描述
// 在斗地主扑克牌游戏中，扑克牌由小到大的顺序为:3.4,5.6,7.8,9,10.J,Q.K.A.2，玩家可以出的扑克牌阵型有:单张、对子、顺子、飞机、炸弹等。
// 其中顺子的出牌规则为:由至少5张由小到大连续递增的扑克牌只 组成，且不能包含2。
// 例如:(3.4,5,6,7}、(3.4,5,6,7,8,9,10,J,Q,K,A}都是有效的顺子;而{J,Q,K,A,2}、(2,3,4,5,6}、(3,4,5,6}、(3,4,5,6,8)等
// 都不是顺子给定一个包含13张牌的数组，如果有满足出牌规则的顺子，请输出顺子。
// 如果存在多个顺子，请每行输出一个顺子，且需要按顺子的第一张牌的大小(必须从小到大)依次输出。
// 如果没有满足出牌规则的顺子，请输出NO。
// 输入描述：
// 13张任意顺序的扑克牌，每张扑克牌数字用空格隔开，每张扑克牌的数字都是合法的，并且不包括大小王:2 9 J 2 3 4 K A 7 9 A 5 6
// 不需要考虑输入为异常α字符的情况
// 输出描述：
// 组成的顺子，每张扑克牌数字用空格隔开:34567
// 示例1
// 输入
// 2 9 J 2 3 4 K A 7 9 A 5 6
// 输出
// 3 4 5 6 7
// 说明13张牌中，可以组成的顺子只有1组:3 4 5 6 7.
// 示例2
// 输入
// 2 9 J 10 3 4 K A 7 Q A 5 6
// 输出
// 3 4 5 6 7
// 9 10 J Q K A
// 说明13张牌中，可以组成2组顺子，从小到大分别为:3 4 5 6 7和9 10 J Q K A
// 示例3
// 输入
// 2 9 9 9 3 4 K A 10 Q A 5 6
// 输出
// No
// 说明13张牌中，无法组成顺子。
func FindSequences(cardArr []int) {

}

func transCard(input string) []int {
	inputArr := strings.Split(input, " ")
	cardArr := make([]int, len(inputArr))
	for _, card := range inputArr {
		cardInt, ok := 0, false
		if cardInt, ok = cardMap[card]; !ok {
			cardInt, _ = strconv.Atoi(card)
		}
		cardArr = append(cardArr, cardInt)
	}
	return cardArr
}

var cardMap = map[string]int{
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
	"2": 0,
}

/*
+------------------------------+-------------------------------+-----------------------------------+-------------------------------------------+
| 方法                         | 使用场景                        | 特点                               | 缺点                                      |
+-----------------------------+-------------------------------+------------------------------------+------------------------------------------+
| fmt.Scan 系列               | 简单格式化输入                   | 快速，语法简单，支持直接解析到变量       | 格式严格，输入需匹配变量类型，处理复杂输入繁琐    |
+----------------------------+-------------------------------+------------------------------------+------------------------------------------+
| bufio.Scanner              | 按行读取或需要循环读取的场景       | 灵活，支持逐行处理输入，内存占用少       | 默认缓冲区限制 64 KB，大量数据需手动调整        |
+----------------------------+-------------------------------+-----------------------------------+-------------------------------------------+
| os.Stdin 和 ioutil.ReadAll | 一次性读取所有输入               | 简单快捷，适合处理多行或大块文本        | 内存占用较高，输入结束需按 EOF                 |
+----------------------------+-----------------------------------+------------------------------------------+-------------------------------+
*/
