package main

import (
	"fmt"
)

func main() {

	vehicles := getVehicles("https://swapi.dev/api/vehicles/")
	oneSeater := vehicles.OneSeater()

	people := getPeople("https://swapi.dev/api/people/")

	printStarWars("----------Star Wars Vehicles-----------", convertVehicles(oneSeater))
	printStarWars("----------Star Wars People-----------", convertPeople(people.Results))
}

func convertPeople(p []people) []printable {

	b := make([]printable, len(p), len(p))
	for i := range p {
		b[i] = p[i]
	}
	return b
}

func convertVehicles(p []vehicle) (b []printable) {

	b = make([]printable, len(p), len(p))
	for i := range p {
		b[i] = p[i]
	}
	return
}

func printStarWars(title string, starWarsPrintable []printable) {

	var cp consolePrinter

	fmt.Println()
	fmt.Println(title)
	fmt.Println()
	for _, p := range starWarsPrintable {
		cp.Println(p)
	}
}

type printable interface {
	Print() string
}

//Trying to wrap an IO function
type consolePrinter struct{}

func (cp consolePrinter) Println(p printable) {
	fmt.Println(p.Print())
}
