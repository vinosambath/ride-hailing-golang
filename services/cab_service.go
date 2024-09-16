package services

import (
	"cabbooking/models"
	"fmt"
	"sync"
)

// Cabs -> addCab, getCabs(), updateLocation, updateAvailiability

type CabService interface {
	AddCab(cab models.ICab) error
	GetCabById(id string) (models.ICab, error)
	UpdateCabLocation(id string, location models.Location) error
	UpdateCabAvailability(id string, isAvailable bool) error
	GetAvailableCable() ([]*models.ICab, error)
}

type CabServiceImpl struct {
	cabs map[string]*models.ICab
}

var CabServiceInstance CabService
var CabServiceSingleton sync.Once

func NewCabService() CabService {
	CabServiceSingleton.Do(func() {
		CabServiceInstance = &CabServiceImpl{
			cabs: make(map[string]*models.ICab),
		}
	})
	return CabServiceInstance
}

func (c *CabServiceImpl) AddCab(cab models.ICab) error {
	c.cabs[cab.GetId()] = &cab
	return nil
}

func (c *CabServiceImpl) GetCabById(id string) (models.ICab, error) {
	return *c.cabs[id], nil
}

func (c *CabServiceImpl) UpdateCabLocation(id string, location models.Location) error {
	cab := *c.cabs[id]
	cab.SetCurrentLocation(location)
	return nil
}

func (c *CabServiceImpl) UpdateCabAvailability(id string, isAvailable bool) error {
	fmt.Println(c.cabs)
	cab := *c.cabs[id]
	cab.SetIsAvailable(isAvailable)
	return nil
}

func (c *CabServiceImpl) GetAvailableCable() ([]*models.ICab, error) {
	availableCabs := make([]*models.ICab, 0)
	for cab := range c.cabs {
		cabReference := *c.cabs[cab]
		if cabReference.GetIsAvailable() {
			availableCabs = append(availableCabs, &cabReference)
		}
	}

	return availableCabs, nil
}
