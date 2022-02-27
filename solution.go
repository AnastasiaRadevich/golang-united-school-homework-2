package square

import (
	"math"
)

type SidesNumber int

const (
	SidesCircle   SidesNumber = 0
	SidesTriangle SidesNumber = 3
	SidesSquare   SidesNumber = 4
)

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {
	var result float64

	if sidesNum == SidesCircle {
		result = math.Pi * math.Pow(sideLen, 2)
	} else if sidesNum == SidesTriangle {
		result = math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	} else if sidesNum == SidesSquare {
		result = math.Pow(sideLen, 2)
	} else {
		result = 0
	}
	return result
}
