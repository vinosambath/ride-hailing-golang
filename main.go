package main

import (
	"fmt"
	"os"
	"os/signal"
	"ride-hailing-golang/controllers"
	"ride-hailing-golang/routes"
	"syscall"
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
	routes.InitRouter()
	err := rHandler.Book("rider1", 0, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hello and welcome!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

}
