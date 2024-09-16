package main

import (
	"cabbooking/controllers"
	"fmt"
)

/*
Cab booking
---------
Models
	Rider -> name, id
	Cab -> id, driver name, isAvailable,
	Trip -> rider, cab, start, end locations, price
	Location -> x, y

Strategy -> pricing (simple - 10$), cabMatchingStrategy (algo)

Services
	Rider -> addRider, getRider
	Cabs -> addCab, getCabs(), updateLocation, updateAvailiability
	Trips -> book, tripHistory,

Controller
	CabsController ->
	RiderController -> addRider, getRider, fetchHistory, bookTrip

---------

*/

func main() {

	rHandler := controllers.GetRiderController()
	cHandler := controllers.NewCabController()

	_ = rHandler.AddRider("rider1", "rider1")
	_ = cHandler.AddCab("cab1", "cab1")
	_ = cHandler.UpdateCabAvailability("cab1", true)
	err := rHandler.Book("rider1", 0, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hello and welcome!")

}
