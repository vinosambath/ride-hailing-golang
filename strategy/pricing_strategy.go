package strategy

import (
	"fmt"
	"ride-hailing-golang/models"
)

type IPricingStrategy interface {
	GetComputedPrice(from models.Location, to models.Location) (float64, error)
}

type PricingStrategy struct{}

func (p *PricingStrategy) GetComputedPrice(from models.Location, to models.Location) (float64, error) {
	return 0, fmt.Errorf("not implemented")
}
