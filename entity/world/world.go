package world

import "game/entity/location"

type World struct {
	Places map[string]*location.Place
}

func New() *World {
	return &World{
		Places: make(map[string]*location.Place, 5),
	}
}

func (w *World) AddPlace(p *location.Place) {
	w.Places[p.Name] = p
}

func (w *World) DeletePlace(p *location.Place) {
	delete(w.Places, p.Name)
}

func (w *World) GetLocationByName(locName string) *location.Place {
	return w.Places[locName]
}
