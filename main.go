package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"hw04/factory"
)

func main() {
	fmt.Printf("Transport Factory System\n")
	fmt.Printf(
		"Now u have available types like:\n(car, motorcycle, truck, bus)\n",
	)
	fmt.Print("Enter vehicle type: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	vehicleType := strings.TrimSpace(strings.ToLower(input))

	var vehicleFactory factory.VehicleFactory

	switch vehicleType {
	case "car":
		fmt.Print("Enter brand: ")
		brand, _ := reader.ReadString('\n')
		fmt.Print("Enter model: ")
		model, _ := reader.ReadString('\n')
		fmt.Print("Enter fuel type (gasoline/diesel/electric): ")
		fuel, _ := reader.ReadString('\n')
		vehicleFactory = factory.NewCarFactory(
			strings.TrimSpace(brand),
			strings.TrimSpace(model),
			strings.TrimSpace(fuel),
		)

	case "motorcycle":
		fmt.Print("Enter type (sport/touring): ")
		motoType, _ := reader.ReadString('\n')
		fmt.Print("Enter engine size (cubes): ")
		engineStr, _ := reader.ReadString('\n')
		engineSize, _ := strconv.Atoi(strings.TrimSpace(engineStr))
		vehicleFactory = factory.NewMotorcycleFactory(
			strings.TrimSpace(motoType),
			engineSize,
		)

	case "truck":
		fmt.Print("Enter capacity (tons): ")
		capStr, _ := reader.ReadString('\n')
		capacity, _ := strconv.Atoi(strings.TrimSpace(capStr))
		vehicleFactory = factory.NewTruckFactory(capacity)

	case "bus":
		fmt.Print("Enter passenger capacity: ")
		capStr, _ := reader.ReadString('\n')
		capacity, _ := strconv.Atoi(strings.TrimSpace(capStr))
		vehicleFactory = factory.NewBusFactory(capacity)

	default:
		fmt.Printf("Unsupported vehicle type: %s\n", vehicleType)
		return
	}

	v := vehicleFactory()
	fmt.Println("\nCreated vehicle:")
	v.Drive()
	v.Refuel()
}
