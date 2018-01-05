package main

import (
	"fmt"
	"math"
	"reflect"
)

type shapes interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type triangle struct {
	side1 float64
	side2 float64
	side3 float64
}

type parallelogram struct {
	height float64
	width  float64
}

type elipse struct {
	majoraxis float64
	minoraxis float64
}

func main() {
	fmt.Println("Area of shapes.")

	allshapes := []shapes{}

	allshapes = append(allshapes, rectangle{10, 5},
		circle{10},
		square{70},
		triangle{15, 22, 19},
		parallelogram{10, 5},
		elipse{10, 5})

	for _, shapeitem := range allshapes {
		outputarea(shapeitem)
	}
}

func outputarea(s shapes) {
	fmt.Printf("Shape: %-15s Area: %10.2f\t Perimeter: %10.2f\n", reflect.TypeOf(s).Name(), s.area(), s.perimeter())
}

func (shape rectangle) area() float64 {
	return shape.height * shape.width
}

func (shape circle) area() float64 {
	return (shape.radius * shape.radius) * math.Pi
}

func (shape square) area() float64 {
	return math.Pow(shape.side, 2)
}

func (shape triangle) area() float64 {
	p := (shape.side1 + shape.side2 + shape.side3) / 2
	return math.Sqrt(p * (p - shape.side1) * (p - shape.side2) * (p - shape.side3))
}

func (shape parallelogram) area() float64 {
	return shape.height * shape.width
}

func (shape elipse) area() float64 {
	return math.Pi * shape.majoraxis * shape.minoraxis
}

func (shape rectangle) perimeter() float64 {
	return (shape.height + shape.width) * 2
}

func (shape circle) perimeter() float64 {
	return 2 * (math.Pi * shape.radius)
}

func (shape square) perimeter() float64 {
	return shape.side * 4
}

func (shape triangle) perimeter() float64 {
	return shape.side1 + shape.side2 + shape.side3
}

func (shape parallelogram) perimeter() float64 {
	return (shape.height + shape.width) * 2
}

func (shape elipse) perimeter() float64 {
	sumaxis := shape.majoraxis + shape.minoraxis
	h := math.Pow(shape.majoraxis-shape.minoraxis, 2) / math.Pow(sumaxis, 2)
	return (math.Pi * sumaxis) * (1 + ((3 * h) / (10 + math.Sqrt(4-(3*h)))))
}
