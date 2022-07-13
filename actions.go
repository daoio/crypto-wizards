package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func CheckChoose(screen *ebiten.Image) {
	switch {
	// REFACTOR!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		ebitenutil.DebugPrint(screen, "A")
		Choosing(1)
	case inpututil.IsKeyJustReleased(ebiten.KeyA):
		ebitenutil.DebugPrint(screen, "A")
		CancelChoose(1)
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		ebitenutil.DebugPrint(screen, "S")
		Choosing(2)
	case inpututil.IsKeyJustReleased(ebiten.KeyS):
		ebitenutil.DebugPrint(screen, "S")
		CancelChoose(2)
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		ebitenutil.DebugPrint(screen, "D")
		Choosing(3)
	case inpututil.IsKeyJustReleased(ebiten.KeyD):
		ebitenutil.DebugPrint(screen, "D")
		CancelChoose(3)
	case inpututil.IsKeyJustPressed(ebiten.KeyF):
		ebitenutil.DebugPrint(screen, "F")
		Choosing(4)
	case inpututil.IsKeyJustReleased(ebiten.KeyF):
		ebitenutil.DebugPrint(screen, "F")
		CancelChoose(4)
	}
}
