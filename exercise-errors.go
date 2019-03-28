package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

// name ...
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

// Sqrt ...
// error という型を返すと、Error() という interface が発動する。
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var z float64 = 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x)/(2*z)
	}
	// var t float64
	// var cnt int
	// t= math.Pow(10, -15)
	// for math.Abs((z*z - x)/(2*z)) > t{
	// 	z -= (z*z - x)/(2*z)
	// 	cnt += 1
	// }
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))


}
