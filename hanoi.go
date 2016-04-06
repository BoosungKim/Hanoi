package main

import "fmt"

// n���� ������ start ���� end �� �ű�µ�, �߰� ��ġ ����� mid �� �����.
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

