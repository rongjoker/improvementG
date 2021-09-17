package main

import "fmt"

func main()  {
	fmt.Println(findKthLargest([]int{3,2,1,5,6,4},2))
	fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6},4))

}

func findKthLargest(nums []int, k int) int {
	size:=len(nums)
	return quickSort215(nums,k,0,size-1)

}

func quickSort215(nums []int, k, left, right int) int {

	l, r := left, right
	pivot := nums[l]
	for l < r {
		for l < r && nums[r] <= pivot {
			r--
		}
		if l < r {
			nums[l] = nums[r]
			l++
		}

		for l < r && nums[l] >= pivot {
			l++
		}

		if l < r {
			nums[r] = nums[l]
			r--
		}

	}
	nums[r] = pivot

	if l == k-1 {
		return nums[l]
	} else if l < k-1 {
		return quickSort215(nums, k, l+1, right)
	} else {
		return quickSort215(nums, k, left, l-1)
	}
}

