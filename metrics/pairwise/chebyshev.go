package pairwise

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

type Chebyshev struct{}

func NewChebyshev() *Chebyshev {
	return &Chebyshev{}
}

func (self *Chebyshev) Distance(vectorX *mat64.Dense, vectorY *mat64.Dense) float64 {
	r1, c1 := vectorX.Dims()
	r2, c2 := vectorY.Dims()
	if r1 != r2 || c1 != c2 {
		panic(mat64.ErrShape)
	}

	subVector := mat64.NewDense(0, 0, nil)
	subVector.Sub(vectorX, vectorY)

	max := float64(0)

	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			max = math.Max(max, math.Abs(vectorX.At(i, j)-vectorY.At(i, j)))
		}
	}

	return max
}
