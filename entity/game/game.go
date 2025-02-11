package game

import (
	"game/entity/player"
	"game/entity/world"
)

const CmdLookAround string = "осмотреться"
const CmdGo string = "идти"
const CmdPutOn string = "надеть"
const CmdTake string = "взять"
const CmdApply string = "применить"
const CmdHave string = "есть"

type Game struct {
	Player               *player.Player
	World                *world.World
	LookAroundHandlers   map[string]func(g *Game) string
	AfterMoveMsgHandlers map[string]func(g *Game) string
}

func New(w *world.World, lookAround, afterMove map[string]func(g *Game) string) *Game {
	return &Game{
		Player:               player.New(w, "кухня"),
		World:                w,
		LookAroundHandlers:   lookAround,
		AfterMoveMsgHandlers: afterMove,
	}
}

func (g *Game) HandleCommand(input string) string {

	command, args := Parce(input)

	switch command {
	case CmdLookAround:
		return g.LookAround()

	case CmdGo:
		return g.GoPlace(args)

	case CmdPutOn:
		return g.PutOm(args)

	case CmdTake:
		return g.Take(args)

	case CmdApply:
		return g.Apply(args)

	case CmdHave:
		return g.Have(args)
	default:
		return g.Unknown(args)

	}
}

func (g *Game) Unknown(args []string) string {

	return "неизвестная команда"
}
