package main

import (
	_ "image/png"
	"log"

	"github.com/daoio/crypto-wizards/cards"
	_ "github.com/daoio/crypto-wizards/eth"
	"github.com/daoio/crypto-wizards/server"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	/*c := cards.GetCards()
	x := 80.0
	y := 250.0

	for i := 0; i < len(c); i++ {
		img, _, err := ebitenutil.NewImageFromFile(c[i])
		if err != nil {
			log.Fatal(err)
		}
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 1)
		x += 50
		op.GeoM.Translate(x, y)
		screen.DrawImage(img, &op)
	}*/
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	// eth initialized => first user has authenticated, then it's time to start server and upgrade it to ws
	server.StartServer()
	server.FindNewGame()
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	// delete temporary files
	cards.DelTemp()
}
