package main

func main() {

}

/**
 * @title计算两数之和
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
