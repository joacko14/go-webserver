package main

import (
	"fmt"
	"webapp/geoFinder"
)

func main() {
	// gorillaserver.Start()
	var lat, long float64

	fmt.Print("Enter latitude: ")
	fmt.Scan(&lat)

	fmt.Print("Enter longitude: ")
	fmt.Scan(&long)

	fmt.Println(geoFinder.FindMe(lat, long))
}
