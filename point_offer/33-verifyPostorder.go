package main

func verifyPostorder(postorder []int) bool {
	return dp(postorder, 0, len(postorder)-1)
}

func dp(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}
	x := left
	for postorder[x] < postorder[right] {
		x++
	}
	y := x
	for postorder[y] > postorder[right] {
		x++
	}

	return x == right && dp(postorder, left, y-1) && dp(postorder, y, right-1)
}
