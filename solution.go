package square

import (
	"math"
)

type NumberShapeSides int64

func CalcSquare(sideLen float64, sidesNum NumberShapeSides) float64 {
	const pi = 3.14
	var result float64

	if sidesNum == 3 {
		result = math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	} else if sidesNum == 4 {
		result = math.Pow(sideLen, 2)
	} else {
		result = pi * math.Pow(sideLen, 2)
	}
	return result
}
