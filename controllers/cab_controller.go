package controllers

import (
	"cabbooking/models"
	"cabbooking/services"
)

type ICabController interface {
	AddCab(id string, name string) error
	UpdateCabLocation(id string, location models.Location) error
	UpdateCabAvailability(id string, isAvailable bool) error
}

type CabController struct {
	cabService services.CabService
}

func NewCabController() ICabController {
	return &CabController{
		cabService: services.NewCabService(),
	}
}

func (c *CabController) AddCab(id string, name string) error {
	err := c.cabService.AddCab(&models.Cab{Id: id, Name: name})
	return err
}

func (c *CabController) UpdateCabLocation(id string, location models.Location) error {
	return c.cabService.UpdateCabLocation(id, location)
}

func (c *CabController) UpdateCabAvailability(id string, isAvailable bool) error {
	return c.cabService.UpdateCabAvailability(id, isAvailable)
}
