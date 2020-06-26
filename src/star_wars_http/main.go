package main

import (
	"fmt"
	"star_wars_http/repository"
	"star_wars_http/types"
)

func main() {

	vehicles := repository.GetVehicles("https://swapi.dev/api/vehicles/")
	people := repository.GetPeople("https://swapi.dev/api/people/")

	oneSeater := vehicles.OneSeater()

	printStarWars("----------Star Wars Vehicles-----------", convertVehicles(oneSeater))
	printStarWars("----------Star Wars People-------------", convertPeople(people.Results))
}

func convertPeople(p []types.People) []printable {

	b := make([]printable, len(p), len(p))
	for i := range p {
		b[i] = p[i]
	}
	return b
}

func convertVehicles(p []types.Vehicle) (b []printable) {

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
