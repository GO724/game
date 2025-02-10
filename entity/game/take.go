package game

import "fmt"

func (g *Game) Take(args []string) string {
	onTable := g.Player.Location.ListItemFromHolder("на столе")

	if len(args) == 0 {
		return fmt.Sprintf("что? [%s]", outputItems(onTable))
	}

	// if g.Player.ItemByName(args[0]).Have() {
	// 	return fmt.Sprintf("уже есть в инвентаре: %s", args[0])
	// }

	if g.Player.Location.IsItemThere(args[0]) {
		if g.Player.Backpack {
			if Item, ok := g.Player.Location.RemoveItemByName(args[0]); ok {
				if answer, ok := g.Player.AddItem(*Item); ok {
					return answer
				}
			}
		} else {
			return "некуда класть"
		}
	} else {
		return "нет такого"
	}

	return fmt.Sprintf("takeItem error : %v", args)
}
