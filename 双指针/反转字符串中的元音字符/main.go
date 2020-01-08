package main

import (
	"fmt"
	"strings"
)

//使用双指针，一个指针从头向尾遍历，一个指针从尾到头遍历，当两个指针都遍历到元音字符时，交换这两个元音字符。
//
//为了快速判断一个字符是不是元音字符，我们将全部元音字符添加到集合 HashSet 中，从而以 O(1) 的时间复杂度进行该操作。
//
//时间复杂度为 O(N)：只需要遍历所有元素一次
//空间复杂度 O(1)：只需要使用两个额外变量


var keyMap = map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "A": 1, "E": 1, "I": 1, "O": 1, "U": 1}

func reverseVowels(s string) (res string, err error) {
	if s == "" {
		return res, fmt.Errorf("参数错误")
	}
	left := 1
	rigth := len(s)
	arr := make([]string, len(s))
	for left < rigth {
		c := 0
		arr[left-1], arr[rigth-1] = s[left-1:left], s[rigth-1:rigth]

		if _, ok := keyMap[arr[left-1]]; ok {
			c += 1
		} else {
			left += 1
		}

		if _, ok := keyMap[arr[rigth-1]]; ok {
			c += 1
		} else {
			rigth -= 1
		}

		if c == 2 {
			arr[left-1], arr[rigth-1] = arr[rigth-1],arr[left-1]
			left += 1
			rigth -= 1
		}
	}

	return strings.Join(arr, ""), err
}
func main()  {
	s := "leetcode"
	ss , err := reverseVowels(s)
	fmt.Println(ss, err)
}
