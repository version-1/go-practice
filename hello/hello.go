package main

import "fmt"

func main() {
	Hello()
}

func Hello() {
	fmt.Println("Hello")
}

func Sum(a int, b int) int {
	return a + b
}

func IsUnderage(a int) bool {
	if a >= 20 {
		return false
	}
	return true
}

func PrintTen() bool {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	return true
}

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return Fib(n - 1) + Fib(n-2)
}

func GetSeqNumArray(n int) []int {
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, i)
	}
	return a
}

func IsDuplicate (a []int) bool {
	m := make(map[int]bool)
	for _, n := range a {
		if m[n] {
			return true
		}
		m[n] = true
	}

	return false
}
