package strategy

import (
	"fmt"
	"ride-hailing-golang/models"
)

type ICabMatchingStrategy interface {
	MatchCabToRider(rider models.IRider, candidateCabs []*models.ICab, _from models.Location, _to models.Location) (*models.ICab, error)
}

type CabmatchingStrategy struct{}

func (c *CabmatchingStrategy) MatchCabToRider(
	_rider models.IRider,
	candidateCabs []*models.ICab,
	_from models.Location,
	_to models.Location) (*models.ICab, error) {
	return nil, fmt.Errorf("Not implemented")
}
