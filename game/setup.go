package game

type Player struct {
	marker  string
	isHuman bool
}

const (
	humans     = "1"
	humanFirst = "2"
	aiFirst    = "3"
)

func (game *Game) SetupPlayers() {
	choice := game.ui.GetGameChoice()

	switch choice {
	case humans:
		game.players = humanHuman()
	case humanFirst:
		game.players = humanAi()
	case aiFirst:
		game.players = aiHuman()
	default:
		game.players = humanAi()
	}
}

func humanHuman() []Player {
	p1 := Player{marker: "x", isHuman: true}
	p2 := Player{marker: "o", isHuman: true}
	return []Player{p1, p2}
}

func aiHuman() []Player {
	p1 := Player{marker: "x", isHuman: false}
	p2 := Player{marker: "o", isHuman: true}
	return []Player{p1, p2}
}

func humanAi() []Player {
	p1 := Player{marker: "x", isHuman: true}
	p2 := Player{marker: "o", isHuman: false}
	return []Player{p1, p2}
}
