package main

import "fmt"

func main()  {
	nums:=[]int{2,7,9,1,3,5,6,4}
	quickSort(nums,0,7)

	fmt.Println(nums)



}

func quickSort(nums []int, left, right int) {
	if left >= right{
		return
	}

	l,r := left,right
	pivot := nums[l]
	for l <r {
		for l < r&& nums[r] <= pivot{
			r--
		}

		if l < r{
			nums[l] = nums[r]
			l++
		}

		for l < r && nums[l] >= pivot {
			l++
		}

		if l < r{
			nums[r] = nums[l]
			r--
		}

	}
	nums[r] = pivot
	quickSort(nums,left,r-1)
	quickSort(nums,r+1,right)


}

