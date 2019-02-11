package main

import (
	"math"
	"math/rand"
	"sync"
)

var waitGroup sync.WaitGroup

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

//Stack is a stack data structure implementation for Triangles
type Stack struct {
	semaphore chan bool
	triangles []Triangle
}

//Push adds the triangle to the top of the stack
func (stack *Stack) Push(triangle Triangle) {
	stack.triangles = append(stack.triangles, triangle)
}

//Peek returns the top triangle on the stack without removing it
func (stack *Stack) Peek() Triangle {
	return stack.triangles[len(stack.triangles)-1]
}

//Pop removes the top triangle from the stack and returns it
func (stack *Stack) Pop() Triangle {
	topTriangle := stack.Peek()
	stack.triangles = stack.triangles[0 : len(stack.triangles)-1]
	return topTriangle
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
