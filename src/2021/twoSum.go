package main

import (
	"fmt"
)

func main(){
	fmt.Println("hello world !" + " I am joker")
	var a int
	fmt.Println("xxxx:",a)
	var xxx string = "Runnable"
	fmt.Println("abc:",xxx)
	var c = true
	fmt.Println("kkkk:",c)

	const name,t = "joker",1
	fmt.Println("name:",name)
	fmt.Println("t:",t)

	tt := 10
	if tt >20 {
		fmt.Println("tt 大于 20")
	}else {
		fmt.Println("tt 小于等于 20")
	}

	sum:=0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("最大值是: ",max(100,200))
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i := 0;i<len(balance2);i++{
		fmt.Println(balance2[i])
	}
	//nums := []int{2,7,11,15}
	nums := []int{3,2,4}
	fmt.Println("ans is ",twoSum(nums, 6))




}

func max(num1, num2 int)  int{

	var ret int

	if num1 >= num2 {
		ret = num1
	}else {
		ret = num2
	}

	return ret

}

func twoSum(nums []int, target int) []int {
	var hash = map[int]int{}/*创建集合 */
	var ans []int

	for i,v :=range nums{
		if k,ok := hash[target - v];ok{
			return []int{k,i}

		}else{
			hash[v] = i
		}

	}

	return ans


}
