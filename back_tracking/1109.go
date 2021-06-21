package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	res := make([]int, n)

	for i := 0; i < len(bookings); i++ {
		start := bookings[i][0] - 1
		end := bookings[i][1] - 1

		diff[start] += bookings[i][2]
		if end+1 < n {
			diff[end+1] -= bookings[i][2]
		}
	}

	res[0] = diff[0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + diff[i]
	}

	return res
}

func main() {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}
