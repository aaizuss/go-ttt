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
