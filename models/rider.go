package models

// Rider -> name, id

type IRider interface {
	GetName() string
	SetName(name string)
	GetId() string
	SetId(id string)
}

type Rider struct {
	Name string
	Id   string
}

func (r *Rider) GetName() string {
	return r.Name
}

func (r *Rider) SetName(name string) {
	r.Name = name
}

func (r *Rider) GetId() string {
	return r.Id
}

func (r *Rider) SetId(id string) {
	r.Id = id
}
