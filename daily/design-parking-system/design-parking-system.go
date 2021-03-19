package design_parking_system

type ParkingSystem struct {
	space [4]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		space: [4]int{0, big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.space[carType] <= 0 {
		return false
	}
	this.space[carType] = this.space[carType] - 1
	return true
}
