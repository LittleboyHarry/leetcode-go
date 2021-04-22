package main

// 记忆 + 递归

var record = make(map[int]int)

func climbStairs(n int) int {
	if n <= 3 {
		return n
	} else {
		if r, ok := record[n]; ok {
			return r
		} else {
			sum := climbStairs(n-1) + climbStairs(n-2)
			record[n] = sum
			return sum
		}
	}
}

func main() {
	climbStairs(4)
}
