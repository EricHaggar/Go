package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Cell is one location in the transportation problem with a unit cost and number of items delivered
type Cell struct {
	cost           float64
	itemsDelivered int
	row            int
	column         int
}

//Route is a regroupment of cells that forms a path
type Route struct {
	path []Cell
}

//Warehouse contains the warehouse properties
type Warehouse struct {
	name   string
	demand int
	column int
}

//Factory contains the factory properties
type Factory struct {
	name   string
	supply int
	row    int
}

//Transportation contains the problem to be solved
type Transportation struct {
	numOfWarehouses int
	numOfFactories  int
	cells           [][]Cell
	warehouses      []Warehouse
	factories       []Factory
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getFileName(fileDescription string) string {
	var fileName string
	fmt.Printf("Please enter the name of the %s text file (without .txt): ", fileDescription)
	_, err := fmt.Scanf("%s \n", &fileName)
	for err != nil {
		fmt.Print("Invalid Input! Please enter a valid string: ")
		_, err = fmt.Scanf("%s \n", &fileName)
	}
	return fileName
}

func openFile(fileName string) []string {
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func initProblem(problemDescription []string) {

}

func main() {
	fmt.Print("Welcome to the Stepping Stone Method Calculator! \n")
	fileName := getFileName("problem description")
	problemDescription := openFile(fileName + ".txt")
	initProblem(problemDescription)
	fileName = getFileName("initial solution")

}
