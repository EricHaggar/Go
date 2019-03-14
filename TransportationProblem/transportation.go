package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Global Variable Declaration

//Cell is one location in the transportation problem with a unit cost and number of items delivered
type Cell struct {
	cost           int
	itemsDelivered int
	row            int
	column         int
}

//Route is a regroupment of cells that forms a path
type Route struct {
	path []Cell
}

//Factory contains the factory properties
type Factory struct {
	name   string
	supply int
	row    int
}

//Warehouse contains the warehouse properties
type Warehouse struct {
	name   string
	demand int
	column int
}

//Transportation contains the problem to be solved
type Transportation struct {
	numOfFactories  int
	numOfWarehouses int
	factories       []Factory
	warehouses      []Warehouse
	costs           [][]int
	cells           [][]Cell
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println()
		log.Fatal(err)
	}
}

func getFileName(fileDescription string) string {
	var fileName string
	fmt.Printf("Enter %s text file name (without .txt): ", fileDescription)
	_, err := fmt.Scanf("%s \n", &fileName)
	for err != nil {
		fmt.Print("Invalid Input! Please enter a valid string: ")
		_, err = fmt.Scanf("%s \n", &fileName)
	}
	return fileName
}

func openFile(fileName string) []string {
	file, err := os.Open(fileName)
	errorCheck(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func newTransportation(problemDescription, initSolution []string) *Transportation {

	var (
		numOfFactories  int
		numOfWarehouses int
		factoryNames    []string
		warehouseNames  []string
		supplies        []int
		demands         []int
		costs           [][]int
		cells           [][]Cell
	)

	for index, line := range problemDescription {

		line := removeSpaces(line)
		size := len(line)

		if index == 0 {
			numOfFactories = stringToInteger(line[0])
			numOfWarehouses = stringToInteger(line[1])
		} else if index == 1 {

			for i := 0; i < numOfWarehouses; i++ {
				warehouseNames = append(warehouseNames, line[i])
			}
		} else if index < (numOfFactories + 2) {

			factoryNames = append(factoryNames, line[0])
			costRow := make([]int, numOfWarehouses)

			for i := 0; i < numOfWarehouses; i++ {
				cost := stringToInteger(line[i+1])
				costRow[i] = cost
			}
			supplies = append(supplies, stringToInteger(line[size-1]))
			costs = append(costs, costRow)

		} else {
			for i := 0; i < numOfWarehouses; i++ {
				demands = append(demands, stringToInteger(line[i+1]))
			}
		}
	}

	itemsDelivered := getItemsDelivered(initSolution, numOfFactories, numOfWarehouses)
	cells = initCells(costs, itemsDelivered, numOfFactories, numOfWarehouses)
	factories := initFactories(factoryNames, supplies)
	warehouses := initWarehouses(warehouseNames, demands)

	return &Transportation{numOfFactories, numOfWarehouses, factories, warehouses, costs, cells}

}

func initFactories(factoryNames []string, supplies []int) []Factory {

	var factories []Factory
	size := len(factoryNames)

	for i := 0; i < size; i++ {
		factory := Factory{factoryNames[i], supplies[i], i}
		factories = append(factories, factory)
	}

	return factories
}

func initWarehouses(warehouseNames []string, demands []int) []Warehouse {

	var warehouses []Warehouse
	size := len(warehouseNames)

	for i := 0; i < size; i++ {
		warehouse := Warehouse{warehouseNames[i], demands[i], i}
		warehouses = append(warehouses, warehouse)
	}

	return warehouses
}

func initCells(costs [][]int, itemsDelivered [][]int, rows int, columns int) [][]Cell {

	var cells [][]Cell

	for i := 0; i < rows; i++ {
		cellsRow := make([]Cell, columns)
		for j := 0; j < columns; j++ {
			cell := Cell{costs[i][j], itemsDelivered[i][j], i, j}
			cellsRow[j] = cell
		}
		cells = append(cells, cellsRow)
	}
	return cells
}

func getItemsDelivered(strArr []string, numOfFactories int, numOfWarehouses int) [][]int {

	var items [][]int

	for index, line := range strArr {

		line := removeSpaces(line)

		if index > 1 && index < (numOfFactories+2) {
			itemRow := make([]int, numOfWarehouses)
			for i := 0; i < numOfWarehouses; i++ {
				item := stringToInteger(line[i+1])
				itemRow[i] = item
			}
			items = append(items, itemRow)
		}

	}
	return items
}

func removeSpaces(s string) []string {
	split := strings.Fields(s)
	return split
}

func stringToInteger(s string) int {
	value, err := strconv.Atoi(s)
	errorCheck(err)
	return value
}

func (t *Transportation) printTransportation() {

	fmt.Printf("---Transportation--- \n")
	fmt.Print("Number of Factories: ")
	fmt.Println(t.numOfFactories)
	fmt.Print("Number of Warehouses: ")
	fmt.Println(t.numOfWarehouses)
	fmt.Print("Factories: ")
	fmt.Println(t.factories)
	fmt.Print("Warehouses: ")
	fmt.Println(t.warehouses)
	fmt.Print("Costs: ")
	fmt.Println(t.costs)
	fmt.Print("Cells: ")
	fmt.Println(t.cells)

}

func main() {
	fmt.Print("***Welcome to the Stepping Stone Method Calculator!*** \n\n")
	//fileName := getFileName("problem description")
	//problemDescription := openFile(fileName + ".txt")
	problemDescription := openFile("input.txt")
	initSolution := openFile("initSol.txt")

	t := newTransportation(problemDescription, initSolution)
	t.printTransportation()
}
