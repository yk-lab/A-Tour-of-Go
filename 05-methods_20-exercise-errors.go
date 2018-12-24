package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	s := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number:  %v", s)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
