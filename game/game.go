// game logic and actions
package game

import (
	"encoding/json"

	"github.com/daoio/crypto-wizards/cards"
	"github.com/ethereum/go-ethereum/common"
)

type Player struct {
	Owner common.Address `json:"owner"`
	Cards cards.Card     `json:"cards"`
	Act   string         `json:"action"`
}

type Players map[common.Address]*Player

type GameState struct {
	MyAddress common.Address `json:"-"`
	IsServer  bool           `json:"-"`
	Players   `json:"players"`
}

type Event struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type EventConnect struct {
	Player
}

// choosing card
type EventChoosing struct {
	// who chooses
	Owner common.Address `json:"owner"`
	// which card (1-4), see cards.Cards
	Card int `json:"card"`
}

type EventUnchoosen struct {
	Owner common.Address `json:"owner"`
	Card  int            `json:"card"`
}

// move card
type EventMove struct {
	// who moves
	Owner common.Address `json:"owner"`
	// which card moves
	Crd cards.Card `json:"card"`
}

// calculating health - damage, etc
type EventBattle struct {
	Owner1 common.Address `json:"owner1"`
	Owner2 common.Address `json:"owner2"`
	// move card1 from owner1
	Crd1 cards.Card `json:"card1"`
	// move card2 from owner2
	Crd2 cards.Card `json:"card2"`
}

type EventInit struct {
	Address common.Address `json:"owner"`
	Players Players        `json:"players"`
}

const EventTypeConnect = "connect"
const EventTypeInit = "init"
const EventTypeChoosing = "choosing"
const EventTypeUnchoosen = "unchoosen"
const EventTypeMove = "move"
const EventTypeBattle = "battle"

func (state *GameState) HandleEvent(event *Event) {
	switch event.Type {
	case EventTypeConnect:
		str, _ := json.Marshal(event.Data)
		var ev EventConnect
		json.Unmarshal(str, &ev)

		state.Players[ev.Owner] = &ev.Player
	case EventTypeInit:
		str, _ := json.Marshal(event.Data)
		var ev EventInit
		json.Unmarshal(str, &ev)

		if !state.IsServer {
			state.MyAddress = ev.Address
			state.Players = ev.Players
		}
	case EventTypeChoosing:
		str, _ := json.Marshal(event.Data)
		var ev EventChoosing
		json.Unmarshal(str, &ev)

		// get card by id
		card := cards.Cards[ev.Owner][ev.Card]
		if card.State == 0 {
			card.Y -= 75
			// choosen
			card.State = 1
		}
	case EventTypeUnchoosen:
		str, _ := json.Marshal(event.Data)
		var ev EventChoosing
		json.Unmarshal(str, &ev)

		// get card by id
		card := cards.Cards[ev.Owner][ev.Card]
		if card.State == 1 {
			card.Y += 75
			// unchoosen
			card.State = 0
		}
	case EventTypeMove:

	}
}
