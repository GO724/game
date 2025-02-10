package game

import (
	"fmt"
)

func MakeLookAroundHandlers() map[string]func(g *Game) string {

	mapLookAroundHandler := make(map[string]func(g *Game) string, 4)

	// {1, "осмотреться", "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"},
	// {01, "осмотреться", "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"},
	// {19, "осмотреться", "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"}, // состояние изменилось
	lookAroundKitchen := func(g *Game) string {
		onTable := g.Player.Location.ListItemFromHolder("на столе")

		view := "ты находишься на кухне,"

		if len(onTable) > 0 {
			view = fmt.Sprintf("%s на столе: %s, надо", view, outputItems(onTable))
		}
		if !g.
			Player.Backpack {
			view = fmt.Sprintf("%s собрать рюкзак и", view)
		}
		view = fmt.Sprintf("%s идти в универ.", view)
		if len(g.Player.Location.Exits) > 0 {
			view = fmt.Sprintf("%s можно пройти - %v", view, outputExits(g.Player.Location.Exits))
		}

		return view
	}
	mapLookAroundHandler["кухня"] = lookAroundKitchen

	// {07, "осмотреться", "на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор"},
	// {10, "осмотреться", "на столе: ключи, конспекты. можно пройти - коридор"}, // состояние изменилось
	// {16, "осмотреться", "пустая комната. можно пройти - коридор"}, // состояние изменилось
	lookAroundRoom := func(g *Game) string {
		onTable := g.Player.Location.ListItemFromHolder("на столе")
		onChair := g.Player.Location.ListItemFromHolder("на стуле")

		notEmptyChair := len(onChair) > 0
		notEmptyTable := len(onTable) > 0
		view := ""

		if notEmptyChair && notEmptyTable {
			view = fmt.Sprintf("на столе: %s, на стуле: %s.", outputItems(onTable), outputItems(onChair))
		} else if notEmptyChair {
			view = fmt.Sprintf("на стуле: %s.", outputItems(onChair))
		} else if notEmptyTable {
			view = fmt.Sprintf("на столе: %s.", outputItems(onTable))
		} else { // empty chair & table
			view = "пустая комната."
		}

		if len(g.Player.Location.Exits) > 0 {
			view = fmt.Sprintf("%s можно пройти - %s", view, outputExits(g.Player.Location.Exits))
		}

		return view
	}
	mapLookAroundHandler["комната"] = lookAroundRoom

	// {20, "идти коридор", },
	lookAroundHall := func(g *Game) string {
		view := "ничего интересного."
		if len(g.Player.Location.Exits) > 0 {
			view = fmt.Sprintf("%s можно пройти - %s", view, outputExits(g.Player.Location.Exits))
		}

		return view
	}
	mapLookAroundHandler["коридор"] = lookAroundHall

	// {21, "идти улица", "дверь закрыта"},                                  // условие не удовлетворено
	// {25, "идти улица", "на улице весна. можно пройти - домой"},
	lookAroundStreet := func(g *Game) string {
		return "на улице весна. можно пройти - домой"
	}
	mapLookAroundHandler["улица"] = lookAroundStreet

	return mapLookAroundHandler
}
