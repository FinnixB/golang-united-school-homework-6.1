package golang_united_school_homework

import (
	"math"
)

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (this Triangle) CalcPerimeter() float64 {
	return 3.0 * this.Side
}

func (this Triangle) CalcArea() float64 {
	return this.Side * this.Side * math.Sqrt(3) / 4
}
