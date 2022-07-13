// game and objects initialization
package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/spf13/viper"
)

// ebiten's game interface implementation
type Game struct {
	background *ebiten.Image
}

// updates screen each tic
func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//opt := ebiten.DrawImageOptions{}

	// get all cards
	c := GetCards()

	CheckChoose(screen)
	switch {
	default:
		// render background
		//screen.DrawImage(g.background, &opt)

		// render cards on background
		for i := 0; i < len(c); i++ {
			c[i].Draw(screen)
		}
	}
}

// returns screen dimension
func (g *Game) Layout(w, h int) (int, int) {
	return viper.GetInt("width"), viper.GetInt("height")
}

func NewGame() *Game {
	back, _, err := ebitenutil.NewImageFromFile(viper.GetString("background"))
	if err != nil {
		log.Fatal(err)
	}
	g := &Game{
		background: back,
	}

	return g
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
	})
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	game := NewGame()
	ebiten.SetWindowTitle(viper.GetString("title"))
	ebiten.SetWindowSize(viper.GetInt("width"), viper.GetInt("height"))

	// start game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
