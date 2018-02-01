package game

type Player struct {
	marker  string
	isHuman bool
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

func (game *Game) SetupPlayers() {
	game.ui.Show("welcome")
	choice := game.ui.GetGameChoice()

	switch choice {
	case "1":
		game.players = humanHuman()
	case "2":
		game.players = humanAi()
	case "3":
		game.players = aiHuman()
	default:
		game.players = humanAi()
	}
}
