package cards

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	shell "github.com/ipfs/go-ipfs-api"
)

// Address => (1-4) => Card
var Cards = make(map[common.Address]map[int]*Card)
var Decks = make(map[common.Address]DeckJSON)

type DeckJSON struct {
	Card1 string `json:"card-1"`
	Card2 string `json:"card-2"`
	Card3 string `json:"card-3"`
	Card4 string `json:"card-4"`
}

type Card struct {
	TokenId int
	Image   string `json:"image"`
	Health  int    `json:"health"`
	Damage  int    `json:"damage"`
	X       float64
	Y       float64
	// 0 - unchoosen, 1 - choosing, 2 - move
	State int
}

func InitDeck(addr common.Address, deckURI string) {
	resp, err := http.Get(deckURI)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var deck DeckJSON

	err = json.Unmarshal(body, &deck)
	if err != nil {
		log.Fatal(err)
	}
	Decks[addr] = deck
	initCards(addr, deck)
}

func initCards(addr common.Address, deck DeckJSON) {
	sh := shell.NewShell("localhost:5001")
	v := reflect.ValueOf(deck)
	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		// values will be an array that stores all cards jsons
		values[i] = v.Field(i).Interface()
	}

	Cards[addr] = make(map[int]*Card)
	// fetch images from JSONs
	for i := 0; i < len(values); i++ {
		out, _ := os.Create("./cards/imgs/" + strconv.Itoa(i+1) + ".png")
		defer out.Close()
		cardJSON := fmt.Sprintf("%v", values[i])
		resp, err := http.Get(cardJSON)
		if err != nil {
			log.Fatal(err)
		}
		var card = Card{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &card)
		fmt.Println("here")
		sh.Get(card.Image, out.Name())
		fmt.Println("here")
		Cards[addr][i+1] = &card
		fmt.Println(Cards[addr][i+1])
	}
}

// get cards of `addr`
func GetCards(addr common.Address) (path []*Card) {
	c := make([]*Card, 4)
	for i := 1; i < 5; i++ {
		c[i-1] = Cards[addr][i+1]
	}

	return c
}

func DelTemp() {
	os.RemoveAll("./cards/imgs")
}
