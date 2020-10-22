/*
 * @Date: 2020-10-21 16:08:53
 * @LastEditTime: 2020-10-22 18:42:36
 * @LastEditors: mmx
 */
package main

func main() {

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
