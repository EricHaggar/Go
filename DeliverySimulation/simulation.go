package main

import (
	"errors"
	"fmt"
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
	if truck.destination != "" {
		return errors.New("Error: Other destination")
	}

	totalLoad = trip.weight + truck.load

	//Checks if added weight from trip would exceed the vehicle capacity
	if totalLoad > truck.capacity {
		return errors.New("Error: Out of capacity")
	}

	//Set distance from Montreal to 200 km and 400 km for Toronto
	if trip.destination == "Montreal" {
		distance = 200.0
	} else {
		distance = 400.0 //Toronto
	}

	//Calculate time required
	timeRequired = (distance) / int(truck.speed)

	//Checks if deadline is feasible with respect to the required travel time
	if trip.deadline < timeRequired {
		return errors.New("Error: Other destination")
	}

	//Passed all error checks
	truck.destination = trip.destination
	truck.load = truck.load + trip.weight

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
	if pickup.destination != "" {
		return errors.New("Error: Other destination")
	}

	totalLoad = trip.weight + pickup.load

	//Checks if added weight from trip would exceed the vehicle capacity
	if totalLoad > pickup.capacity {
		return errors.New("Error: Out of capacity")
	}

	//Set distance from Montreal to 200 km and 400 km for Toronto
	if trip.destination == "Montreal" {
		distance = 200.0
	} else {
		distance = 400.0 //Toronto
	}

	//Calculate time required
	timeRequired = (distance) / int(pickup.speed)

	//Checks if deadline is feasible with respect to the required travel time
	if trip.deadline < timeRequired {
		return errors.New("Error: Other destination")
	}

	//Passed all error checks
	pickup.destination = trip.destination
	pickup.load = pickup.load + trip.weight

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
	if trainCar.destination != "" {
		return errors.New("Error: Other destination")
	}

	totalLoad = trip.weight + trainCar.load

	//Checks if added weight from trip would exceed the vehicle capacity
	if totalLoad > trainCar.capacity {
		return errors.New("Error: Out of capacity")
	}

	//Set distance from Montreal to 200 km and 400 km for Toronto
	if trip.destination == "Montreal" {
		distance = 200.0
	} else {
		distance = 400.0 //Toronto
	}

	//Calculate time required
	timeRequired = (distance) / int(trainCar.speed)

	//Checks if deadline is feasible with respect to the required travel time
	if trip.deadline < timeRequired {
		return errors.New("Error: Other destination")
	}

	//Passed all error checks
	trainCar.destination = trip.destination
	trainCar.load = trainCar.load + trip.weight

	//return nil if there's no error
	return nil
}

//Print trip information for each type of vehicle
func (truck *Truck) print() {
	fmt.Printf("%v to %v with %f tons\n", truck.name, truck.destination, truck.load)
}
func (pickup *Pickup) print() {
	fmt.Printf("%v to %v with %f tons (%+v)\n", pickup.name, pickup.destination, pickup.load, pickup.isPrivate)
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

func main() {

}
