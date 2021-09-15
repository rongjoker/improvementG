package main

import (
	"fmt"
)

//峰值元素是指其值严格大于左右相邻值的元素。
//
//给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//
//你可以假设 nums[-1] = nums[n] = -∞ 。
//
//你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-peak-element
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//输入：nums = [1,2,3,1]
//输出：2
//解释：3 是峰值元素，你的函数应该返回其索引 2。
//输入：nums = [1,2,1,3,5,6,4]
//输出：1 或 5
//解释：你的函数可以返回索引 1，其峰值元素为 2；
//     或者返回索引 5， 其峰值元素为 6。

func main() {

	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{1}))
	fmt.Println(findPeakElement([]int{1, 2}))
	fmt.Println(findPeakElement([]int{6, 5, 4, 3, 2, 3, 2}))

}

func findPeakElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[length-1] > nums[length-2] {
		return length - 1
	}
	left := 1
	right := len(nums) - 2
	for left <= right {
		middle := left + ((right - left) >> 1)
		if (nums[middle] > nums[middle+1]) && (nums[middle] > nums[middle-1]) {
			return middle
		} else if nums[middle] < nums[middle+1] {
			left = middle + 1
		} else {
			right = middle - 1
		}

	}

	return -1

}
