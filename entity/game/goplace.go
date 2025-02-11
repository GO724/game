package game

import "fmt"

func (g *Game) GoPlace(args []string) string {
	if len(args) == 0 {
		return "куда?"
	}

	place, ok := g.World.Places[args[0]]
	if !ok {
		return "неизвестное место"
	}

	if !g.Player.Location.CheckWay(args[0]) {
		return fmt.Sprintf("нет пути в %s", args[0])
	}

	if !g.Player.Location.CheckDoor(args[0]) {
		g.Player.NextLocation = args[0]
		return "дверь закрыта"
	}

	g.Player.PrevLocation = g.Player.Location
	g.Player.Location = place // update player location
	g.Player.NextLocation = ""

	afterMoveMsg := g.AfterMoveMsgHandlers[g.Player.Location.Name]

	return afterMoveMsg(g)
}
