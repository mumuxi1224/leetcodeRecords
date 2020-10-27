/*
 * @Date: 2020-10-26 17:34:58
 * @LastEditTime: 2020-10-27 18:57:01
 * @LastEditors: mmx
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(20))
}

/**
* @Date: 2020-10-26 17:36:12
* @LastEditors: mmx
* @title 回文数
* @rule 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
* @url https://leetcode-cn.com/problems/palindrome-number/
* @param {int} x
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tempX := x
	var y int = 0
	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}
	fmt.Println(tempX, y)
	return tempX == y
}

/**
 * @Date: 2020-10-26 17:50:27
 * @LastEditors: mmx
 * @title 整数转罗马数字
 * @rule 罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
		字符          数值
		I             1
		V             5
		X             10
		L             50
		C             100
		D             500
		M             1000
		例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

		通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

		I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
		X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
		C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
		给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。
 * @url https://leetcode-cn.com/problems/integer-to-roman/
 * @param {int} num
*/
func intToRoman(num int) string {
	var result string
	var aims = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var char = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	//不用每次调整char
	tempK := 0
	for num != 0 {
		for k, v := range aims {
			if num >= v {
				tempK += k
				num -= v
				result += char[tempK]
				//从上次开始的位置比较
				aims = aims[k:]
				break
			}
		}
	}

	return result
}
