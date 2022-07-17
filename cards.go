// all about cards
package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/spf13/viper"
)

var Cards = make(map[int]*Card)

type Card struct {
	// TODO: ipfsUrl : string
	Img    *ebiten.Image
	Health int
	Damage int
	w      int // width
	h      int // height
	x      int
	y      int
}

func (c *Card) Draw(target *ebiten.Image) {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Scale(2, 2)
	opt.GeoM.Translate(float64(c.x), float64(c.y))
	target.DrawImage(c.Img, &opt)
}

func GetCards() []*Card {
	c := make([]*Card, 4)

	for i := 0; i < 4; i++ {
		c[i] = Cards[i+1]
	}

	return c
}

// pop card up when user chooses it
func Choosing(n int) {
	switch {
	case CardStates[Cards[n]].Choosen && !CardStates[Cards[n]].Move:
		Cards[n].y -= 75
		PlayerStates[Players[1]].Choosing = true
		PlayerStates[Players[1]].Card = n
		CardStates[Cards[n]].Choosen = false
	case !CardStates[Cards[n]].Choosen && !CardStates[Cards[n]].Move:
		Cards[n].y += 75
		PlayerStates[Players[1]].Choosing = false
		PlayerStates[Players[1]].Card = 0
		CardStates[Cards[n]].Choosen = true
	}
}

func init() {
	// init coordinates
	x := 450
	y := 570

	for i := 1; i < 5; i++ {
		img, _, err := ebitenutil.NewImageFromFile("./sprites/cards/" + strconv.Itoa(i) + ".png")
		if err != nil {
			log.Fatal(err)
		}
		Cards[i] = &Card{
			Img:    img,
			Health: viper.GetInt("card-" + strconv.Itoa(i) + ".Health"),
			Damage: viper.GetInt("card-" + strconv.Itoa(i) + ".Damage"),
			w:      viper.GetInt("card-" + strconv.Itoa(i) + ".w"),
			h:      viper.GetInt("card-" + strconv.Itoa(i) + ".h"),
			x:      (x + 2),
			y:      y,
		}

		x += 50
	}
}
