package strategy

import (
	"fmt"
	"ride-hailing-golang/models"
)

type DefaultCabMatchingStrategy struct{}

func (s DefaultCabMatchingStrategy) MatchCabToRider(
	_rider models.IRider,
	candidateCabs []*models.ICab,
	_from models.Location,
	_to models.Location) (*models.ICab, error) {

	if len(candidateCabs) == 0 {
		return nil, fmt.Errorf("no candidate Cabs found")
	}

	return candidateCabs[0], nil
}
