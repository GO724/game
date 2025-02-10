package game

import "fmt"

func (g *Game) Apply(args []string) string {
	// apply args[0] to args[1]
	item := g.Player.ItemByName(args[0])
	if !item.Have() {
		return fmt.Sprintf("нет предмета в инвентаре - %s", args[0])
	}

	if args[1] != "дверь" {
		return "не к чему применить"
	}

	// if g.Player.NextLocation == "" {
	// 	return "не к чему применить"
	// }

	doorOpen := false
	for _, door := range g.Player.Location.Exits {
		if door.InDoor {
			if door.InDoorKey {
				if !door.InDoorOpen {
					door.InDoorOpen = true
					doorOpen = true
				}
			}
		}
	}
	// if !g.Player.Location.OpenDoor(g.Player.NextLocation, item.Have()) {
	// 	return "дверь закрыта"
	// }

	if doorOpen {
		return "дверь открыта"
	}

	return "дверь закрыта"
}
