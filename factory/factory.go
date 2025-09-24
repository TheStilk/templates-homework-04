package factory

import (
	"hw04/vehicle"
)

type VehicleFactory func() vehicle.Vehicle

func NewCarFactory(brand, model, fuel string) VehicleFactory {
	return func() vehicle.Vehicle {
		return &vehicle.Car{
			Brand:  brand,
			Model:  model,
			Fuel:   fuel,
		}
	}
}

func NewMotorcycleFactory(motoType string, engineSize int) VehicleFactory {
	return func() vehicle.Vehicle {
		return &vehicle.Motorcycle{
			Type:       motoType,
			EngineSize: engineSize,
		}
	}
}

func NewTruckFactory(capacity int) VehicleFactory {
	return func() vehicle.Vehicle {
		return &vehicle.Truck{
			Capacity: capacity,
		}
	}
}

func NewBusFactory(capacity int) VehicleFactory {
	return func() vehicle.Vehicle {
		return &vehicle.Bus{
			Capacity: capacity,
		}
	}
}
