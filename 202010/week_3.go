/*
 * @Date: 2020-10-21 16:08:53
 * @LastEditTime: 2020-10-23 13:14:31
 * @LastEditors: mmx
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	reverse(-132)
}

/**
 * @Date: 2020-10-21 17:42:00
 * @LastEditors: mmx
 * @title 计算两数之和
 * @rule 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
		 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 * @url https://leetcode-cn.com/problems/two-sum
 * @param {[]int} nums [3,3]
 * @param {int} target 6
 * @return {*}
 * @tips 除了暴力破解之外，可以创建一个nums的value作key的map
		 遍历时判断target-value的结果是否在map中
		 每次遍历完成将nums的key和value放入map中 map[value] = key  防止使用相同的元素
*/
func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for k, v := range nums {
		if rv := target - v; temp[rv] != k {
			if rk, ok := temp[rv]; ok {
				return []int{k, rk}
			}
		}
		temp[v] = k
	}
	return []int{-1, -1}
}

/**
 * @Date: 2020-10-21 18:03:42
 * @LastEditors: mmx
 * @title 无重复字符的最长子串
 * @rule 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 * @url https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 * @param {string} s
 * @tips 构造一个滑块模型。不重复的元素位置做左边界，遍历的key做右边界
		 当有重复元素时，将左边界移动到重复元素位置 key-左边界+1的值即为当前边界长度
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	start := 0
	temp := make(map[rune]int)
	result := 0
	for k, v := range s {
		index, ok := temp[v]
		if ok && index >= start {
			//滑动到重复元素后一位
			start = index + 1
		}
		temp[v] = k
		if result < k-start {
			result = k - start
		}
	}
	return result + 1
}

/**
 * @Date: 2020-10-22 21:32:09
 * @LastEditors: mmx
 * @param {string} date
 * @title 转变日期格式
 * @rule 给你一个字符串 date ，它的格式为 Day Month Year ，其中：
			Day 是集合 {"1st", "2nd", "3rd", "4th", ..., "30th", "31st"} 中的一个元素。
			Month 是集合 {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"} 中的一个元素。
			Year 的范围在 ​[1900, 2100] 之间。
			请你将字符串转变为 YYYY-MM-DD 的格式，其中：

			YYYY 表示 4 位的年份。
			MM 表示 2 位的月份。
			DD 表示 2 位的天数。  输入：date = "20th Oct 2052"  输出："2052-10-20"
 * @url https://leetcode-cn.com/problems/reformat-date/
 * @param {string} s
*/
func reformatDate(date string) string {
	dateMap := map[string]string{
		"Jan": "1", "Feb": "2", "Mar": "3", "Apr": "4", "May": "5", "Jun": "6", "Jul": "7", "Aug": "8", "Sep": "9", "Oct": "10", "Nov": "11", "Dec": "12",
	}
	newDate := strings.Split(date, " ")
	day, _ := strconv.Atoi(newDate[0][:len(newDate[0])-2])
	return fmt.Sprintf("%s-%02s-%02d", newDate[2], dateMap[newDate[1]], day)
}

/**
 * @Date: 2020-10-23 11:45:37
 * @LastEditors: mmx
 * @title 整数反转
 * @rule 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 * @url https://leetcode-cn.com/problems/reverse-integer/
 * @param {int} x
 */
func reverse(x int) int {
	max := 1<<31 - 1
	min := -1 << 31
	y := 0
	for x != 0 {
		y = y*10 + (x % 10)
		x /= 10
	}
	if y > max || y < min {
		return 0
	}
	return y
}
