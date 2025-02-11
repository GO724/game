package game

import (
	"fmt"
)

func MakeAfterMoveHandlers() map[string]func(g *Game) string {

	mapAfterMoveHandler := make(map[string]func(g *Game) string, 4)

	afterMoveKitchen := func(g *Game) string {
		view := "кухня, ничего интересного."

		if len(g.Player.Location.Exits) > 0 {
			view = fmt.Sprintf("%s можно пройти - %v", view, outputExits(g.Player.Location.Exits))
		}

		return view
	}
	mapAfterMoveHandler["кухня"] = afterMoveKitchen

	afterMoveLookAroundRoom := func(g *Game) string {
		view := "ты в своей комнате."
		if len(g.Player.Location.Exits) > 0 {
			view = fmt.Sprintf("%s можно пройти - %s", view, outputExits(g.Player.Location.Exits))
		}

		return view
	}
	mapAfterMoveHandler["комната"] = afterMoveLookAroundRoom

	afterMoveHall := func(g *Game) string {
		view := "ничего интересного."
		if len(g.Player.Location.Exits) > 0 {
			view = fmt.Sprintf("%s можно пройти - %s", view, outputExits(g.Player.Location.Exits))
		}

		return view
	}
	mapAfterMoveHandler["коридор"] = afterMoveHall

	afterMoveStreet := func(g *Game) string {
		return "на улице весна. можно пройти - домой"
	}
	mapAfterMoveHandler["улица"] = afterMoveStreet

	return mapAfterMoveHandler
}
