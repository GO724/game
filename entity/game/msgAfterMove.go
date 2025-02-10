package game

import (
	"fmt"
)

func MakeAfterMoveHandlers() map[string]func(g *Game) string {

	mapLookAroundHandler := make(map[string]func(g *Game) string, 4)

	// {1, "осмотреться", "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"},
	// {01, "осмотреться", "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"},
	// {19, "осмотреться", "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"}, // состояние изменилось
	lookAroundKitchen := func(g *Game) string {
		view := "кухня, ничего интересного."

		if len(g.Player.Location.Exits) > 0 {
			view = fmt.Sprintf("%s можно пройти - %v", view, outputExits(g.Player.Location.Exits))
		}

		return view
	}
	mapLookAroundHandler["кухня"] = lookAroundKitchen

	// {07, "осмотреться", "на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор"},
	// {10, "осмотреться", "на столе: ключи, конспекты. можно пройти - коридор"}, // состояние изменилось
	// {16, "осмотреться", "пустая комната. можно пройти - коридор"}, // состояние изменилось
	afterMoveLookAroundRoom := func(g *Game) string {
		view := "ты в своей комнате."
		if len(g.Player.Location.Exits) > 0 {
			view = fmt.Sprintf("%s можно пройти - %s", view, outputExits(g.Player.Location.Exits))
		}

		return view
	}
	mapLookAroundHandler["комната"] = afterMoveLookAroundRoom

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
