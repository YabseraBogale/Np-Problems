package algorthims

import (
	"fmt"
	"math"
)

func HammingDistance(string_1, string_2 string) (int, error) {
	count := 0
	if len(string_1) != len(string_2) {
		return 0, fmt.Errorf("Strings must be of equal length.")
	}
	for i, _ := range string_1 {
		if string_1[i] == string_2[i] {
			count++
		}
	}
	return count, nil
}

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
