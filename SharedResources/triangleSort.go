package main

import (
	"math"
	"math/rand"
)

//Point defines a point with an x and y coordinate
type Point struct {
	x float64
	y float64
}

//Triangle defines the 3 Points of a triangle
type Triangle struct {
	A Point
	B Point
	C Point
}

func triangles10000() [10000]Triangle {
	var tableau [10000]Triangle
	rand.Seed(2120)
	for i := 0; i < 10000; i++ {
		tableau[i].A = Point{rand.Float64() * 100., rand.Float64() * 100.}
		tableau[i].B = Point{rand.Float64() * 100., rand.Float64() * 100.}
		tableau[i].C = Point{rand.Float64() * 100., rand.Float64() * 100.}
	}
	return tableau
}

//Perimeter calculates the perimeter of a Triangle
func (t Triangle) Perimeter() float64 {

	var perimeter float64

	//Calculating triangle side lengths using coordinates
	sideLengthA := CalculateLength(t.B.x, t.A.x, t.B.y, t.A.y)
	sideLengthB := CalculateLength(t.C.x, t.B.x, t.C.y, t.B.y)
	sideLengthC := CalculateLength(t.A.x, t.C.x, t.A.y, t.C.y)

	perimeter = sideLengthA + sideLengthB + sideLengthC

	return perimeter

}

//Area calculates the area of a Triangle
func (t Triangle) Area() float64 {

	var area float64

	sideLengthA := CalculateLength(t.B.x, t.A.x, t.B.y, t.A.y)
	sideLengthB := CalculateLength(t.C.x, t.B.x, t.C.y, t.B.y)
	sideLengthC := CalculateLength(t.A.x, t.C.x, t.A.y, t.C.y)

	//Heron's Formula for calculating area of a triangle
	average := (sideLengthA + sideLengthB + sideLengthC) / 2
	area = math.Sqrt(average * (average - sideLengthA) * (average - sideLengthB) * (average - sideLengthC))

	return area
}

//CalculateLength returns the length of one side of the triangle
func CalculateLength(x2 float64, x1 float64, y2 float64, y1 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {

}
