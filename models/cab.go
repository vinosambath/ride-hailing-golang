package models

// Cab -> id, driver name, isAvailable,

type ICab interface {
	GetId() string
	SetId(id string)
	GetName() string
	SetName(name string)
	GetIsAvailable() bool
	SetIsAvailable(isAvailable bool)
	GetCurrentLocation() Location
	SetCurrentLocation(location Location)
}

type Cab struct {
	Id              string
	Name            string
	IsAvailable     bool
	CurrentLocation Location
}

func (c *Cab) GetId() string {
	return c.Id
}

func (c *Cab) SetId(id string) {
	c.Id = id
}

func (c *Cab) GetName() string {
	return c.Name
}

func (c *Cab) SetName(name string) {
	c.Name = name
}

func (c *Cab) GetIsAvailable() bool {
	return c.IsAvailable
}

func (c *Cab) SetIsAvailable(isAvailable bool) {
	c.IsAvailable = isAvailable
}

func (c *Cab) GetCurrentLocation() Location {
	return c.CurrentLocation
}

func (c *Cab) SetCurrentLocation(location Location) {
	c.CurrentLocation = location
}
