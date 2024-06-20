package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}
func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	fmt.Printf("Running a computation funciton with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 = ", compute(2, 2, add))
	fmt.Println("2 * 2 = ", compute(2, 2, func(lhs, rhs int) int {
		return lhs * rhs
	}))

	div := func(lhs, rhs int) int {
		fmt.Printf("Dividing %v with %v = ", lhs, rhs)
		return lhs / rhs
	}

	fmt.Println(compute(12, 4, div))
}
