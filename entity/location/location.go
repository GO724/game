package location

import (
	"fmt"
	"game/entity/item"
)

type Place struct {
	Name         string
	ItemsHolders map[string]struct{}
	Items        map[string]*item.Item
	Exits        map[string]*struct {
		Index      int  // index in output list
		InDoor     bool // is enterance door?
		InDoorOpen bool // is enterance door open?
		InDoorKey  bool // is enterance door key needed?
	}
}

func NewPlace(name string) *Place {
	return &Place{
		Name:         name,
		ItemsHolders: make(map[string]struct{}),
		Items:        make(map[string]*item.Item),
		Exits: make(map[string]*struct {
			Index      int  // index in output list
			InDoor     bool // is enterance door?
			InDoorOpen bool // is enterance door open?
			InDoorKey  bool // is enterance door key needed?
		}, 3),
	}
}

func (p *Place) AddItemsHolder(itemsHolder string) {
	p.ItemsHolders[itemsHolder] = struct{}{}
}

func (p *Place) DeleteItemsHolder(itemsHolder string) {
	delete(p.ItemsHolders, itemsHolder)
}

func (p *Place) AddItem(item *item.Item) {
	p.Items[item.Name] = item
}

func (p *Place) DeleteItem(item *item.Item) {
	delete(p.Items, item.Name)
}

func (p *Place) RemoveItemByName(itemName string) (*item.Item, bool) {
	// at any items holder
	thisItem, ok := p.Items[itemName]
	if !ok {
		return nil, false
	}
	result := thisItem.Decrease()

	if thisItem.Empty() {
		delete(p.Items, itemName)
	}

	return thisItem, result
}

func (p *Place) ListItemsHolrers() []string {
	return make([]string, 0)
}

func (p *Place) ListItemFromHolder(itemsHolder string) []*item.Item {
	result := make([]*item.Item, 0, 3)

	for _, v := range p.Items {
		if v.PlaceHolder == itemsHolder {
			result = append(result, v)
		}
	}

	return result
}

func (p *Place) IsItemThere(itemName string) bool {
	// at any items holder
	thisItem, ok := p.Items[itemName]
	if !ok {
		return false
	}

	if thisItem.Empty() {
		delete(p.Items, itemName)
		return false
	}

	return true
}

func (p *Place) AddExit(index int, toPlace string, door struct {
	Index      int
	InDoor     bool
	InDoorOpen bool
	InDoorKey  bool
}) {
	door.Index = index
	p.Exits[toPlace] = &door
}

func (p *Place) DeleteExit(toPlace string) {
	delete(p.Exits, toPlace)
}

func (p *Place) SetDoor(toPlace string, inDoorOpen, inDoorKey bool) bool {
	if door, ok := p.Exits[toPlace]; ok {
		door.InDoorKey = inDoorKey
		door.InDoorOpen = inDoorOpen
		return true
	}
	return false
}

func (p *Place) OpenDoor(toPlace string, inDoorKey bool) bool {
	if door, ok := p.Exits[toPlace]; ok {
		if door.InDoorKey {
			if inDoorKey {
				if !door.InDoorOpen {
					door.InDoorOpen = true
				}
			}
		} else {
			if !door.InDoorOpen {
				door.InDoorOpen = true
			}
		}
		return door.InDoorOpen
	}
	return false
}

func (p *Place) CheckDoor(toPlace string) bool {
	if door, ok := p.Exits[toPlace]; ok {
		if door.InDoor {
			return door.InDoorOpen
		}
	}
	return true
}

func (p *Place) CheckWay(toPlace string) bool {
	_, ok := p.Exits[toPlace]
	return ok
}

func (p *Place) OnEnter() string {
	fmt.Println("Enter to", p.Name)
	return "onEnter"
}

func (p *Place) OnExit() string {
	fmt.Println("Exit from", p.Name)
	return "onExit"
}
