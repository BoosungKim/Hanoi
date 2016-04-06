package main

import "fmt"

// n개의 원판을 start 에서 end 로 옮기는데, 중간 거치 막대로 mid 를 사용함.
func hanoi(n int, start int, mid int, end int) {
	if n == 1 {
		fmt.Println(start, "->", end)
		return
	}
	hanoi(n-1, start, end, mid)
	hanoi(1, start, mid, end)
	hanoi(n-1, mid, start, end)
}
func main() {
	hanoi(3, 1, 3, 2)
}

