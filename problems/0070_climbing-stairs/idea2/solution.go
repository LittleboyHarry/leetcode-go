package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}
	return curr
}
