package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
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

//Stack is a stack data structure implementation for Triangles
type Stack struct {
	semaphore chan int
	triangles []Triangle
}

//Push adds the triangle to the top of the stack
func (stack *Stack) Push(triangle Triangle) {
	stack.semaphore <- 1 //acquire semaphore
	stack.triangles = append(stack.triangles, triangle)
	<-stack.semaphore //release semaphore
}

//Peek returns the top triangle on the stack without removing it
func (stack *Stack) Peek() Triangle {
	stack.semaphore <- 1
	topTriangle := stack.triangles[len(stack.triangles)-1]
	<-stack.semaphore
	return topTriangle
}

//Pop removes the top triangle from the stack and returns it
func (stack *Stack) Pop() Triangle {
	stack.semaphore <- 1
	topTriangle := stack.Peek()
	stack.triangles = stack.triangles[0 : len(stack.triangles)-1]
	<-stack.semaphore
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

func classifyTriangle(highRatio *Stack, lowRatio *Stack, ratioThreshold float64, triangles []Triangle, waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	for i := range triangles {

		ratio := triangles[i].Perimeter() / triangles[i].Area()

		if ratio < ratioThreshold {
			lowRatio.Push(triangles[i])
		} else {
			highRatio.Push(triangles[i])
		}
	}
}

func main() {

	var waitGroup sync.WaitGroup
	ratioThreshold := 1.0

	triangles := triangles10000()

	//Create semaphores
	lowStackSemaphore := make(chan int, 1)
	highStackSemaphore := make(chan int, 1)

	//Create stacks
	lowStack := Stack{lowStackSemaphore, []Triangle{}}
	highStack := Stack{highStackSemaphore, []Triangle{}}

	for i := 0; i < 10; i++ {

		waitGroup.Add(1)
		go classifyTriangle(&highStack, &lowStack, ratioThreshold, triangles[i*10:(i*10)+1000], &waitGroup)
	}

	waitGroup.Wait()

	//Goroutines are finished executing

	lowStackTop := lowStack.Peek()
	highStackTop := highStack.Peek()

	fmt.Print("\n")
	fmt.Println("\t\t\tResults")
	fmt.Println("--------------------------------------------------------")
	fmt.Println()
	fmt.Printf("LowRatio Stack size : %d triangles\n", len(lowStack.triangles))
	fmt.Println("Triangle at the top of the LowRatio stack:")
	fmt.Printf("Coordinates: Point A (%f, %f), Point B (%f, %f), Point C (%f, %f) \n", lowStackTop.A.x, lowStackTop.A.y, lowStackTop.B.x, lowStackTop.B.y, lowStackTop.C.x, lowStackTop.C.y)

	fmt.Println()
	fmt.Printf("HighRatio Stack size : %d triangles\n", len(highStack.triangles))
	fmt.Println("Triangle at the top of the HighRatio stack:")
	fmt.Printf("Coordinates: Point A (%f, %f), Point B (%f, %f), Point C (%f, %f) \n", highStackTop.A.x, highStackTop.A.y, highStackTop.B.x, highStackTop.B.y, highStackTop.C.x, highStackTop.C.y)

}
