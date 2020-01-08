package main

import (
	"fmt"
)

//所谓的回文字符串，是指具有左右对称特点的字符串，例如 "abcba" 就是一个回文字符串。
//
//使用双指针可以很容易判断一个字符串是否是回文字符串：令一个指针从左到右遍历，一个指针从右到左遍历，这两个指针同时移动一个位置，
//每次都判断两个指针指向的字符是否相同，如果都相同，字符串才是具有左右对称性质的回文字符串。

func valid(s string, left, rigth, delTimes int) bool {
	if s == "" {
		return false
	}
	for left <= rigth {
		if s[left-1:left] == s[rigth-1:rigth] {
			left += 1
			rigth -= 1
		} else {
			delTimes += 1
			if delTimes == 2 {
				return false
			}
			if s[left: left + 1] == s[rigth-1:rigth] {
				if valid(s, left + 1, rigth, delTimes) {
					return true
				}
			}
			if s[left-1:left] == s[rigth - 2: rigth - 1] {
				if valid(s, left, rigth -1, delTimes) {
					return true
				}
			}
			return false
		}
	}

	return true
}

func main()  {
	s := "ffffffaaaaaaaafffffff"
	fmt.Println(valid(s, 1, len(s), 0))
}
