package game

import "fmt"

func (g *Game) Have(args []string) string {
	if len(args) == 0 {
		return outputPlayerItems(g.Player.Items)
	}

	item := g.Player.ItemByName(args[0])
	if item.Have() {
		return fmt.Sprintf("%s : есть  %d", args[0], item.Quantity)
	}

	return "нет"
}
