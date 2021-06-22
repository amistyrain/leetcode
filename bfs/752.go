package main

import "fmt"

func openLock(deadends []string, target string) int {
	deads := make(map[string]struct{}, len(deadends))
	for _, dead := range deadends {
		deads[dead] = struct{}{}
	}
	visited := make(map[string]struct{})
	initD := `0000`
	visited[initD] = struct{}{}

	list := []string{}
	list = append(list, initD)
	step := 0

	for len(list) > 0 {
		size := len(list)
		for i := 0; i < size; i++ {
			if list[i] == target {
				return step
			}
			if _, ok := deads[list[i]]; ok {
				continue
			}
			for j := 0; j < 4; j++ {
				downStr := down(list[i], j)
				if _, ok := visited[downStr]; !ok {
					visited[downStr] = struct{}{}
					list = append(list, downStr)
				}
				upStr := up(list[i], j)
				if _, ok := visited[upStr]; !ok {
					visited[upStr] = struct{}{}
					list = append(list, upStr)
				}
			}
		}
		list = list[size:]
		step++
	}

	return -1
}

func down(str string, index int) string {
	bt := []byte(str)
	if bt[index] == '0' {
		bt[index] = '9'
	} else {
		bt[index] = byte(int(bt[index]) - 1)
	}

	return string(bt)
}

func up(str string, index int) string {
	bt := []byte(str)
	if bt[index] == '9' {
		bt[index] = '0'
	} else {
		bt[index] = byte(int(bt[index]) + 1)
	}

	return string(bt)
}

func main() {
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, `8888`))
}
