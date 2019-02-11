package main

import (
	"errors"
	"fmt"
	"strings"
)

//Transporter is an interface which defines addLoad and print method signatures
type Transporter interface {
	addLoad(trip Trip) error
	print()
}

//Trip defines a trip from Ottawa to Montreal or Toronto
type Trip struct {
	destination string
	weight      float32
	deadline    int
}

//Transportation defines variables that will be used by Truck, Pickup and TrainCar
type Transportation struct {
	vehicle     string
	name        string
	destination string
	speed       float32
	capacity    float32
	load        float32
}

//Truck is a type of vehicle to be used for the delivery
type Truck struct {
	Transportation
}

//Pickup is a type of vehicle to be used for the delivery
type Pickup struct {
	Transportation
	isPrivate bool
}

//TrainCar is a type of vehicle to be used for the delivery
type TrainCar struct {
	Transportation
	railway string
}

//NewTruck returns an initialized Truck
func NewTruck() *Truck {
	return &Truck{Transportation{
		vehicle:     "",
		name:        "",
		destination: "",
		speed:       40,
		capacity:    10,
		load:        0,
	}}
}

//NewPickUp returns an initialized Pickup
func NewPickUp() *Pickup {
	return &Pickup{Transportation{
		vehicle:     "",
		name:        "",
		destination: "",
		speed:       60,
		capacity:    2,
		load:        0,
	},
		true}
}

//NewTrainCar returns an initialized TrainCar
func NewTrainCar() *TrainCar {
	return &TrainCar{Transportation{
		vehicle:     "",
		name:        "",
		destination: "",
		speed:       30,
		capacity:    30,
		load:        0,
	},
		"CNR"}
}

func (truck *Truck) addLoad(trip Trip) error {

	//initialize needed variables
	var distance int
	var timeRequired int
	var totalLoad float32

	/*Error checks*/
	//Checks if vehicle already has a destination
	if truck.destination != trip.destination && truck.destination != "" {
		return errors.New("Error: Other destination")
	}

	totalLoad = trip.weight + truck.load

	//Checks if added weight from trip would exceed the vehicle capacity
	if totalLoad > truck.capacity {
		return errors.New("Error: Out of capacity")
	}

	distance = SetDistance(trip)

	//Calculate time required to destination
	timeRequired = GetTimeRequired(truck.speed, distance)

	//Checks if deadline is feasible with respect to the required travel time
	if trip.deadline < timeRequired {
		return errors.New("Error: Can't meet deadline")
	}

	//Passed all error checks
	truck.load = totalLoad

	if truck.destination == "" {
		truck.destination = trip.destination
	}

	//return nil if there's no error
	return nil
}

func (pickup *Pickup) addLoad(trip Trip) error {

	//initialize needed variables
	var distance int
	var timeRequired int
	var totalLoad float32

	/*Error checks*/
	//Checks if vehicle already has a destination
	if pickup.destination != trip.destination && pickup.destination != "" {
		return errors.New("Error: Other destination")
	}

	totalLoad = trip.weight + pickup.load

	//Checks if added weight from trip would exceed the vehicle capacity
	if totalLoad > pickup.capacity {
		return errors.New("Error: Out of capacity")
	}

	distance = SetDistance(trip)

	//Calculate time required to destination
	timeRequired = GetTimeRequired(pickup.speed, distance)

	//Checks if deadline is feasible with respect to the required travel time
	if trip.deadline < timeRequired {
		return errors.New("Error: Can't meet deadline")
	}

	//Passed all error checks
	pickup.load = totalLoad

	if pickup.destination == "" {
		pickup.destination = trip.destination
	}

	//return nil if there's no error
	return nil
}

func (trainCar *TrainCar) addLoad(trip Trip) error {

	//initialize needed variables
	var distance int
	var timeRequired int
	var totalLoad float32

	/*Error checks*/
	//Checks if vehicle already has a destination
	if trainCar.destination != trip.destination && trainCar.destination != "" {
		return errors.New("Error: Other destination")
	}

	totalLoad = trip.weight + trainCar.load

	//Checks if added weight from trip would exceed the vehicle capacity
	if totalLoad > trainCar.capacity {
		return errors.New("Error: Out of capacity")
	}

	distance = SetDistance(trip)

	//Calculate time required to destination
	timeRequired = GetTimeRequired(trainCar.speed, distance)

	//Checks if deadline is feasible with respect to the required travel time
	if trip.deadline < timeRequired {
		return errors.New("Error: Can't meet deadline")
	}

	//Passed all error checks
	trainCar.load = totalLoad

	if trainCar.destination == "" {
		trainCar.destination = trip.destination
	}

	//return nil if there's no error
	return nil
}

//Print trip information for each type of vehicle
func (truck *Truck) print() {
	fmt.Printf("%v to %v with %f tons\n", truck.name, truck.destination, truck.load)
}
func (pickup *Pickup) print() {
	fmt.Printf("%v to %v with %f tons (Private: %v)\n", pickup.name, pickup.destination, pickup.load, pickup.isPrivate)
}
func (trainCar *TrainCar) print() {
	fmt.Printf("%v to %v with %f tons (%v)\n", trainCar.name, trainCar.destination, trainCar.load, trainCar.railway)
}

//NewTorontoTrip creates a new Trip to Toronto
func NewTorontoTrip(weight float32, deadline int) *Trip {
	return &Trip{
		destination: "Toronto",
		weight:      weight,
		deadline:    deadline,
	}
}

//NewMontrealTrip creates a new Trip to Montreal
func NewMontrealTrip(weight float32, deadline int) *Trip {
	return &Trip{
		destination: "Montreal",
		weight:      weight,
		deadline:    deadline,
	}
}

//SetDistance sets the distance with respect to the destination
func SetDistance(trip Trip) (distance int) {

	//Set distance from Montreal to 200 km and 400 km for Toronto
	if trip.destination == "Montreal" {
		distance = 200.0
	} else {
		distance = 400.0 //Toronto
	}
	return
}

//GetTimeRequired returns the travel time with respect to the speed and distance
func GetTimeRequired(speed float32, distance int) (timeRequired int) {
	return int(speed) / distance
}

func main() {

	//Creating 2 Trucks, 3 Pickups and 1 TrainCar
	truckA := NewTruck()
	truckA.name = "Truck A"

	truckB := NewTruck()
	truckB.name = "Truck B"

	pickUpA := NewPickUp()
	pickUpA.name = "Pickup A"

	pickUpB := NewPickUp()
	pickUpB.name = "Pickup B"

	pickUpC := NewPickUp()
	pickUpC.name = "Pickup C"

	trainCarA := NewTrainCar()
	trainCarA.name = "TrainCar A"

	transporters := []Transporter{truckA, truckB, pickUpA, pickUpB, pickUpC, trainCarA}

	//Define needed variables
	var weight float32
	var deadline int
	var trip *Trip
	var trips []*Trip
	var destination string

	for {
		fmt.Print("Destination: (t)oronto, (m)ontreal, else exit? ")
		fmt.Scanln(&destination)

		destination = strings.ToLower(destination)

		if destination[0] == 'm' { //Montreal
			destination = "Montreal"
		} else if destination[0] == 't' { //Toronto
			destination = "Toronto"
		} else { //Quit
			fmt.Print("Not going to TO or Montreal, bye! \n")
			break
		}

		fmt.Print("Weight: ")
		fmt.Scanln(&weight)

		fmt.Print("Deadline (in hours): ")
		fmt.Scanln(&deadline)

		if destination == "Montreal" {
			trip = NewMontrealTrip(weight, deadline)
		} else {
			trip = NewTorontoTrip(weight, deadline)
		}

		for i := range transporters {

			err := transporters[i].addLoad(*trip)

			if err == nil {
				trips = append(trips, trip)
				break
			} else {
				fmt.Println(err)
				if i == len(transporters)-1 {
					fmt.Println("Sorry, none of the Transporters can deliver!")
					break
				}
			}
		}
	}

	fmt.Print("Trips: [")

	for i := range trips {
		if i == len(trips)-1 {
			fmt.Printf("{%v %g %d}", trips[i].destination, trips[i].weight, trips[i].deadline)
		} else {
			fmt.Printf("{%v %g %d} ", trips[i].destination, trips[i].weight, trips[i].deadline)
		}
	}
	fmt.Print("]\n")

	fmt.Println("Vehicles :")

	for i := range transporters {
		transporters[i].print()
	}

}
