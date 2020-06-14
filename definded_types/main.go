package main

import "fmt"

type Gallons float64
type Liters float64

func (l Liters) toGallons() Gallons {
	return Gallons(l * 0.264)
}

func (l Liters) String() string {
	return fmt.Sprintf("%.2f", float64(l)) + " liters"
}

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.785)
}

func (g *Gallons) Double() {
	*g *= 2
}

func (g Gallons) String() string {
	return fmt.Sprintf("%.2f", float64(g)) + " gallons"
}

func main() {
	carFuel := Gallons(10.0)
	busFuel := Liters(240.0)

	fmt.Println(carFuel, busFuel)
	fmt.Println(carFuel.toLiters(), busFuel.toGallons())

	carFuel.Double()
	fmt.Println(carFuel, busFuel)
}
