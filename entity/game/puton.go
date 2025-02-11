package game

import "fmt"

func (g *Game) PutOm(args []string) string {
	onChair := g.Player.Location.ListItemFromHolder("на стуле")

	if len(args) == 0 {
		if g.Player.Backpack {
			return "рюкзак уже надет"
		} else if len(onChair) > 0 {
			return fmt.Sprintf("что? [%s]", outputItems(onChair))
		} else {
			return "нет предметов"
		}
	}

	answer := "вы надели:"
	if args[0] == "рюкзак" && g.Player.Location.IsItemThere(args[0]) {
		if _, ok := g.Player.Location.RemoveItemByName(args[0]); ok {
			g.Player.Backpack = true
			return fmt.Sprintf("%s %s", answer, args[0])
		}
	}
	return fmt.Sprintf("недоступно: %s", args[0])

}
