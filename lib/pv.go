package lib

import "math"

func PV(fv float64, d float64, n int) float64 {
	dn := math.Pow((1.0 + d), float64(n))
	return fv / dn
}
