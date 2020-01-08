package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(target int) (res []int, err error) {
	if target <= 0 {
		return res, fmt.Errorf("target 不能<=0")
	}
	left := 0
	rigth := int(math.Sqrt(float64(target)))

	for left < rigth {
		sum := int(left * left + rigth * rigth)
		if sum == target {
			return []int{int(left), int(rigth)}, err
		} else if sum > target {
			rigth -= 1
		} else {
			left += 1
		}
	}
	return res, err
}

func main()  {
	res, err := judgeSquareSum(11)
	fmt.Println(res, err)
}
