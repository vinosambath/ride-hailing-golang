package services

import (
	"fmt"
	"ride-hailing-golang/models"
	"ride-hailing-golang/strategy"
	"sync"
)

type TripService interface {
	BookTrip(rider models.IRider, from models.Location, to models.Location) (*models.Trip, error)
}

type tripServiceImpl struct {
	tripsByUser         map[string][]*models.Trip
	cabMatchingStrategy strategy.ICabMatchingStrategy
	pricingStrategy     strategy.IPricingStrategy
	cabService          CabService
}

var TripServiceInstance TripService
var TripServiceInit sync.Once

func NewTripService(cabMatchingStrategy strategy.ICabMatchingStrategy, pricingStrategy strategy.IPricingStrategy) TripService {
	TripServiceInit.Do(func() {
		TripServiceInstance = &tripServiceImpl{
			tripsByUser:         make(map[string][]*models.Trip),
			cabMatchingStrategy: cabMatchingStrategy,
			pricingStrategy:     pricingStrategy,
			cabService:          NewCabService(),
		}
	})

	return TripServiceInstance
}

func (t *tripServiceImpl) BookTrip(rider models.IRider, from models.Location, to models.Location) (*models.Trip, error) {
	cabs, err := t.cabService.GetAvailableCable()
	if err != nil {
		return nil, err
	}

	bestCab, err := t.cabMatchingStrategy.MatchCabToRider(rider, cabs, from, to)

	if err != nil {
		return nil, err
	}

	price, err := t.pricingStrategy.GetComputedPrice(from, to)
	if err != nil {
		return nil, err
	}

	trip := &models.Trip{
		Rider: rider,
		Cab:   *bestCab,
		Price: price,
		From:  from,
		To:    to,
	}

	fmt.Println("trip", trip)

	return trip, nil
}
