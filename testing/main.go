package main

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
