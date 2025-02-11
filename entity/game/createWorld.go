package game

import (
	i "game/entity/item"
	l "game/entity/location"
	w "game/entity/world"
)

func CreateWorld() *w.World {
	withoutDoor := struct {
		Index      int
		InDoor     bool
		InDoorOpen bool
		InDoorKey  bool
	}{
		InDoor:     false,
		InDoorOpen: false,
		InDoorKey:  false,
	}

	closedDoor := struct {
		Index      int
		InDoor     bool
		InDoorOpen bool
		InDoorKey  bool
	}{
		InDoor:     true,
		InDoorOpen: false,
		InDoorKey:  true,
	}

	kitchen := l.NewPlace("кухня")
	kitchen.AddItemsHolder("на столе")
	kitchen.AddItem(i.New("чай", "на столе", 1))
	kitchen.AddExit(1, "коридор", withoutDoor)

	room := l.NewPlace("комната")
	room.AddItemsHolder("на столе")
	room.AddItemsHolder("на стуле")
	room.AddItem(i.New("ключи", "на столе", 1))
	room.AddItem(i.New("конспекты", "на столе", 1))
	room.AddItem(i.New("рюкзак", "на стуле", 1))
	room.AddExit(1, "коридор", withoutDoor)

	hall := l.NewPlace("коридор")
	hall.AddExit(1, "кухня", withoutDoor)
	hall.AddExit(2, "комната", withoutDoor)
	hall.AddExit(3, "улица", closedDoor)

	street := l.NewPlace("улица")
	street.AddExit(1, "коридор", withoutDoor)

	world := w.New()
	world.AddPlace(kitchen)
	world.AddPlace(room)
	world.AddPlace(hall)
	world.AddPlace(street)

	return world
}
