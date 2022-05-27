package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (this Rectangle) CalcPerimeter() float64 {
	return this.Height*2.0 + this.Weight*2.0
}

func (this Rectangle) CalcArea() float64 {
	return this.Height * this.Weight
}
