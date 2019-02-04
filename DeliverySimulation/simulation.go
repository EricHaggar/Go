package main

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

func main() {

}
