/*
 * @Date: 2020-12-08 14:16:09
 * @LastEditTime: 2020-12-09 14:04:40
 * @LastEditors: mmx
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("mississippi", "issip"))
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
