package controllers

import (
	"ride-hailing-golang/models"
	"ride-hailing-golang/services"
	"ride-hailing-golang/strategy"
)

type IRiderController interface {
	AddRider(id string, name string) error
	Book(riderId string, sourceX float64, sourceY float64, destX float64, destY float64) error
}

type RiderController struct {
	riderService services.RiderService
	tripService  services.TripService
}

func GetRiderController() IRiderController {
	return &RiderController{
		riderService: services.NewRiderService(),
		tripService:  services.NewTripService(strategy.DefaultCabMatchingStrategy{}, &strategy.DefaultPricingStrategy{}),
	}
}

func (rc *RiderController) AddRider(id string, name string) error {
	return rc.riderService.AddRider(&models.Rider{
		Id:   id,
		Name: name,
	})
}

func (rc *RiderController) Book(riderId string, sourceX float64, sourceY float64, destX float64, destY float64) error {
	rider, err := rc.riderService.GetRider(riderId)
	if err != nil {
		return err
	}
	_, err = rc.tripService.BookTrip(*rider, models.Location{sourceX, sourceY}, models.Location{destX, destY})
	if err != nil {
		return err
	}
	return nil
}
