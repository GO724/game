package player

import (
	"fmt"
	"game/entity/item"
	"game/entity/location"
	"game/entity/world"
)

type Player struct {
	Location     *location.Place
	PrevLocation *location.Place // last move : from
	NextLocation string          // last move : to (if door close), or empty
	Items        map[string]*item.Item
	Backpack     bool
}

func New(world *world.World, spawnLocation string) *Player {
	spawnPlace := world.GetLocationByName(spawnLocation)

	if spawnPlace == nil {
		panic(fmt.Sprintf("spawn location not found : %s", spawnLocation))
	}

	return &Player{
		Location: spawnPlace,
		Items:    make(map[string]*item.Item, 3),
	}
}

func (p *Player) SetItems(items map[string]*item.Item) {
	p.Items = items
}

func (p *Player) AddItem(i item.Item) (answer string, result bool) {
	if !p.Backpack {
		answer = "некуда класть"
		result = false
		return
	}

	//if p.Location.HaveItem()
	thisItem, ok := p.Items[i.Name]
	if !ok {
		p.Items[i.Name] = &i
		thisItem = p.Items[i.Name]
	}
	thisItem.Quantity++

	answer = "предмет добавлен в инвентарь: " + i.Name
	result = true
	return
}

func (p *Player) RemoveItem(i item.Item) (answer string, result bool) {
	if !p.Backpack {
		answer = "нет такого"
		result = false
		return
	}
	thisItem, ok := p.Items[i.Name]
	if !ok {
		answer = "нет такого"
		result = false
		return
	}
	thisItem.Quantity--

	answer = "предмет удален из инвентаря: " + i.Name
	result = true
	return
}

func (p *Player) ItemByName(iName string) *item.Item {
	if !p.Backpack {
		return nil
	}
	item, ok := p.Items[iName]
	if !ok {
		return nil
	}
	return item
}
