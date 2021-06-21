package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(diffWaysToCompute(`2*3-4*5`))
}

func diffWaysToCompute(expression string) []int {
	res := []int{}

	for i := 0; i < len(expression); i++ {
		if expression[i] == '-' || expression[i] == '+' || expression[i] == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			for _, vl := range left {
				for _, vr := range right {
					if expression[i] == '-' {
						res = append(res, vl-vr)
					} else if expression[i] == '+' {
						res = append(res, vl+vr)
					} else if expression[i] == '*' {
						res = append(res, vl*vr)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		i, _ := strconv.Atoi(expression)
		res = append(res, i)
	}

	return res
}
