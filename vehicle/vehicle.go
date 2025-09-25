// Package vehicle
package vehicle

import "fmt"

type Vehicle interface {
	Drive()
	Refuel()
}

type Car struct {
	Brand string
	Model string
	Fuel  string
}

func (c *Car) Drive() {
	fmt.Printf("Car %s %s is driving on %s fuel\n", c.Brand, c.Model, c.Fuel)
}

func (c *Car) Refuel() {
	fmt.Printf("Refueling car %s %s with %s\n", c.Brand, c.Model, c.Fuel)
}

type Motorcycle struct {
	Type       string
	EngineSize int
}

func (m *Motorcycle) Drive() {
	fmt.Printf("Motorcycle %s, %d cubes\n", m.Type, m.EngineSize)
}

func (m *Motorcycle) Refuel() {
	fmt.Printf("Refueling motorcycle %s with gasoline\n", m.Type)
}

type Truck struct {
	Capacity int
}

func (t *Truck) Drive() {
	fmt.Printf("Truck capacity: %d tons\n", t.Capacity)
}

func (t *Truck) Refuel() {
	fmt.Printf("Refueling heavy truck with diesel\n")
}

type Bus struct {
	Capacity int
}

func (b *Bus) Drive() {
	fmt.Printf("Bus capacity %d is transporting passengers\n", b.Capacity)
}

func (b *Bus) Refuel() {
	fmt.Printf("Refueling bus with gas\n")
}
