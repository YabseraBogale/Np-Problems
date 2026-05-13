package algorthims

import (
	"fmt"
	"math"
)

func NewtonRaphson(f func(float64) float64, fd func(float64) float64, initialGuess float64, tolerance float64, maxIterations int) (float64, error) {
	x := initialGuess
	for i := 0; i < maxIterations; i++ {
		fx := f(x)
		dfx := fd(x)
		if math.Abs(dfx) < 1e-12 {
			return x, fmt.Errorf("derivative too small; possible local extremum at x = %s", x)
		}
		nextX := x - (fx / dfx)
		if math.Abs(nextX-x) < tolerance {
			fmt.Printf("Converged in %d iterations.\n", i+1)
			return nextX, nil
		}
		x = nextX
	}
	return x, nil
}
