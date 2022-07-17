package main

var Players = make(map[int]*Player)

type Player struct {
	ID int
}

func GetPlayers() []*Player {
	p := make([]*Player, 1)

	p[0] = Players[1]

	return p
}

func Move(p *Player) {
	if PlayerStates[p].Choosing && !CardStates[Cards[PlayerStates[p].Card]].Move {
		// move card on the table
		Cards[PlayerStates[p].Card].y -= 250
		// lock next card's movements
		CardStates[Cards[PlayerStates[p].Card]].Move = true
	}
}

func init() {
	Players[1] = &Player{ID: 1}
}
