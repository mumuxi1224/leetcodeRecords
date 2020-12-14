/*
 * @Date: 2020-12-08 14:16:09
 * @LastEditTime: 2020-12-14 11:43:09
 * @LastEditors: mmx
 */
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 5, 6}
	fmt.Println(searchInsert(nums, 7))
}

/**
 * @Date: 2020-12-08 14:19:07
 * @LastEditors: mmx
 * @param {string} haystack
 * @param {string} needle
 * 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/implement-strstr
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func strStr(haystack string, needle string) int {
	hlen := len(haystack)
	nlen := len(needle)
	if nlen == 0 || haystack == needle {
		return 0
	}
	if hlen < nlen {
		return -1
	}
	h, n, index := 0, 0, -1 //定义指针
	for h < hlen {
		if haystack[h] == needle[n] {
			if index == -1 { //记录第一次相等的下标
				index = h
			}
			if n == nlen-1 { //比较完成 返回下标
				return index
			}
			n++
		} else {
			if index != -1 {
				h = index
				n = 0
				index = -1
			}
		}
		h++
	}
	return -1
}

/**
 * @Date: 2020-12-09 14:07:08
 * @LastEditors: mmx
 * @param {int} dividend
 * @param {int} divisor
 *给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
	你可以假设数组中无重复元素。
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func searchInsert(nums []int, target int) int {
	// nLen := len(nums)
	// for k, v := range nums {
	// 	if v == target {
	// 		return k
	// 	}
	// 	if target > nums[nLen-1] {
	// 		return nLen
	// 	}
	// 	if target < nums[0] {
	// 		return 0
	// 	}
	// 	if target > v && target < nums[k+1] {
	// 		return k + 1
	// 	}
	// }
	// return 0

	//二分法
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		// mid = (left + right) / 2
		mid = left + (right-left)/2 //防止整数溢出
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
