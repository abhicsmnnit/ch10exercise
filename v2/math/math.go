// Package math contains types and functions for doing exactly what you expect it to do... math!
package math

import (
	"golang.org/x/exp/constraints"
)

// Number is a constraint that permits only integers and float types.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds numbers
//
// It accepts two Number parameters, and returns the sum of the two Numbers.
//
// For more details on addition, check out [MathIsFun].
//
// [MathIsFun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](num1, num2 T) T {
	return num1 + num2
}
