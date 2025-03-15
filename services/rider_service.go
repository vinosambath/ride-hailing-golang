package services

import (
	"errors"
	"ride-hailing-golang/models"
	"sync"
)

// Rider -> addRider, getRider

type RiderService interface {
	AddRider(rider models.IRider) error
	GetRider(riderId string) (*models.IRider, error)
}

type riderServiceImpl struct {
	riders map[string]*models.IRider
}

var RiderServiceInstance RiderService
var RiderServiceSingleton sync.Once

func NewRiderService() RiderService {
	RiderServiceSingleton.Do(func() {
		RiderServiceInstance = &riderServiceImpl{
			riders: make(map[string]*models.IRider),
		}
	})
	return RiderServiceInstance
}

func (r *riderServiceImpl) AddRider(rider models.IRider) error {
	_, ok := r.riders[rider.GetId()]
	if ok {
		return errors.New("Rider already exists")
	}
	r.riders[rider.GetId()] = &rider
	return nil
}

func (r *riderServiceImpl) GetRider(riderId string) (*models.IRider, error) {
	return r.riders[riderId], nil
}
