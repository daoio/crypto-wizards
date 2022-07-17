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
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		ebitenutil.DebugPrint(screen, "S")
		Choosing(2)
	case inpututil.IsKeyJustReleased(ebiten.KeyS):
		ebitenutil.DebugPrint(screen, "S")
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		ebitenutil.DebugPrint(screen, "D")
		Choosing(3)
	case inpututil.IsKeyJustReleased(ebiten.KeyD):
		ebitenutil.DebugPrint(screen, "D")
	case inpututil.IsKeyJustPressed(ebiten.KeyF):
		ebitenutil.DebugPrint(screen, "F")
		Choosing(4)
	case inpututil.IsKeyJustReleased(ebiten.KeyF):
		ebitenutil.DebugPrint(screen, "F")
	}
}

func CheckMove(screen *ebiten.Image) {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyEnter):
		Move(Players[1])
	}
}
