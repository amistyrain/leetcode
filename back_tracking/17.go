package main

import "fmt"

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return res
	}

	backtrack(digits, ``, 0)
	fmt.Println(res)
	return res
}

var res []string

var numTobyte = map[byte]string{
	'2': `abc`,
	'3': `def`,
	'4': `ghi`,
	'5': `jkl`,
	'6': `mno`,
	'7': `pqrs`,
	'8': `tuv`,
	'9': `wxyz`,
}

func backtrack(digits, used string, start int) {
	if len(used) == len(digits) {
		res = append(res, used)
		return
	}
	str := numTobyte[digits[start]]

	for i := 0; i < len(str); i++ {
		used += string(str[i])
		backtrack(digits, used, start+1)
		used = used[:len(used)-1]

	}
}

func main() {
	letterCombinations(`235`)
}
