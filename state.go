package main

// state maps
var CardStates = make(map[*Card]*CardState)
var PlayerStates = make(map[*Player]*PlayerState)

type GameState struct {
	// importing ethereum account via private key
	Import bool
}

type CardState struct {
	Choosen bool
	Move    bool
	Lost    bool
	Won     bool
}

type PlayerState struct {
	Choosing bool
	// id of choosen card
	Card int
	Move bool
}

func init() {
	c := GetCards()
	p := GetPlayers()

	for i := 0; i < len(p); i++ {
		PlayerStates[p[i]] = &PlayerState{
			Choosing: false,
			Card:     0,
			Move:     false,
		}
	}

	for i := 0; i < len(c); i++ {
		CardStates[c[i]] = &CardState{
			Choosen: true,
		}
	}
}
