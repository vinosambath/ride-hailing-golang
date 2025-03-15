package strategy

import "ride-hailing-golang/models"

type DefaultPricingStrategy struct{}

func (d *DefaultPricingStrategy) GetComputedPrice(from models.Location, to models.Location) (float64, error) {
	return 10.0, nil
}
