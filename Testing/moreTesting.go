package main

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}
