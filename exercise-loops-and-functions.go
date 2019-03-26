package main
import (
	"fmt"
	"math"
)

// Sqrt ...
func Sqrt(x float64) int {
	var z float64 = 1.0
	// for i := 0; i < 10; i++ {
	// 	z -= (z*z - x)/(2*z)
	// }
	var t float64
	var cnt int
	t= math.Pow(10, -15)
	for math.Abs((z*z - x)/(2*z)) > t{
		z -= (z*z - x)/(2*z)
		cnt += 1
	}
	return cnt
}

func main() {
	var s float64 = 11
	fmt.Println(Sqrt(s))
	fmt.Println(math.Sqrt(s))


}

