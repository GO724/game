package game

func (g *Game) LookAround() string {
	lookAround := g.LookAroundHandlers[g.Player.Location.Name]

	return lookAround(g)
}
