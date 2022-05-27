package golang_united_school_homework

import (
	"errors"
)

var (
	errorBoxIsFull         = errors.New("Box is full")
	errorBoxIsEmpty        = errors.New("Box is empty")
	errorBoxOutRange       = errors.New("Box shape index is out of range")
	errorShapeIndex        = errors.New("Shape index not exists")
	errorShapeNotExists    = errors.New("Shape does not exists in box")
	errorThereAreNoCircles = errors.New("There are no circles in box")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	if shapesCapacity > 0 {
		return &box{
			shapesCapacity: shapesCapacity,
			shapes:         make([]Shape, 0),
		}
	}

	return nil
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}

	return errorBoxIsFull
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i > b.shapesCapacity {
		return nil, errorBoxOutRange
	}

	if i >= len(b.shapes) {
		return nil, errorShapeIndex
	}

	return b.shapes[i], nil
}

func (b *box) DeleteByIndex(i int) error {
	_, err := b.GetByIndex(i)

	if err == nil {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	}

	return err
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	s, err := b.GetByIndex(i)

	if err == nil {
		return s, b.DeleteByIndex(i)
	}

	return s, err
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	s, err := b.GetByIndex(i)

	if err == nil {
		b.shapes[i] = shape
	}

	return s, err
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	result := 0.0

	for i := 0; i < len(b.shapes); i++ {
		result = result + b.shapes[i].CalcPerimeter()
	}

	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	result := 0.0

	for i := 0; i < len(b.shapes); i++ {
		result = result + b.shapes[i].CalcArea()
	}

	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	haveCircles := false

	for i := len(b.shapes) - 1; i >= 0; i-- {
		_, ok := b.shapes[i].(*Circle)
		if ok == true {
			b.DeleteByIndex(i)
			haveCircles = true
		}
	}

	if haveCircles {
		return nil
	}

	return errorThereAreNoCircles
}
