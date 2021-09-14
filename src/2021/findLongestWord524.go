package main

import "fmt"

//给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
//
//如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
//输出："apple"
//输入：s = "abpcplea", dictionary = ["a","b","c"]
//输出："a"

func main() {

	fmt.Println(findLongestWord("abpcplea",[]string{"ale","apple","monkey","plea"}))
	fmt.Println(findLongestWord("abpcplea",[]string{"b","c","a","p"}))

}

func findLongestWord(s string, dictionary []string) string {
	ans := ""
	for _, key := range dictionary {
		i := 0
		for j := range s {
			if s[j] == key[i] {
				i++
				if i == len(key) {
					if len(key) > len(ans) || (len(key) == len(ans) && key < ans) {
						ans = key
					}
					break
				}
			}
		}

	}

	return ans

}

