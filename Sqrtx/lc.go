package sqrtx

import "math"

func MySqrt(x int) int {
    
	x1 := float64(x)

	a := math.Sqrt(x1)

	return int(math.Floor(a))

}