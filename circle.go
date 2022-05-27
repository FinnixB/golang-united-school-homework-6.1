package golang_united_school_homework

import (
	"math"
)

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (this Circle) CalcPerimeter() float64 {
	return 2.0 * math.Pi * this.Radius
}

func (this Circle) CalcArea() float64 {
	return math.Pi * this.Radius * this.Radius
}
