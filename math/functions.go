package math
import "math"
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func IsEqual(f1, f2, p float64) bool {
	return math.Abs(f1 - f2) < math.Abs(p)
}