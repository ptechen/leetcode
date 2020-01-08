package main

//给定一个整数数组，返回两个数字的索引，以便它们加起来成为一个特定的目标。
//
//您可以假定每个输入都只有一个解决方案，并且您可能不会两次使用同一元素。
import "fmt"

func twoSum1(nums []int, target int) []int {
	left :=0
	right := len(nums) - 1
	for {
		if nums[left] < target {
			if left == right {
				break
			}
			if nums[left] + nums[right] == target {
				return []int{left, right}
			}else if nums[left] + nums[right] > target{
				right -= 1
			} else {
				left += 1
			}
		}
	}

	return []int{}
}


func main()  {
	arr := []int{0, 2, 2, 4, 7, 9}
	target := 10

	A := twoSum1(arr, target)
	fmt.Println(A)
}
