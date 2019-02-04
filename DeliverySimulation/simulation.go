package main

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
	*Transportation
}

//Pickup is a type of vehicle to be used for the delivery
type Pickup struct {
	*Transportation
	isPrivate bool
}

//TrainCar is a type of vehicle to be used for the delivery
type TrainCar struct {
	*Transportation
	railway string
}

func main() {

}
