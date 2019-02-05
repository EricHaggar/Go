package main

import "errors"

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
		""}
}

func (truck *Truck) addLoad(trip Trip) error {

	//initialize needed variables
	var distance int
	var timeRequired int

	/*Error checks*/
	if truck.destination != "" {
		return errors.New("Error: Other destination")
	}

	if trip.weight > truck.capacity {
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

	if trip.deadline > timeRequired {
		return errors.New("Error: Other destination")
	}

	//Passed all error checks
	truck.destination = trip.destination
	truck.capacity = trip.weight

	return nil
}

func main() {

}
