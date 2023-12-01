package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Game struct {
	Player *Player
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Vector struct {
	X float64
	Y float64
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{
		Player: NewPlayer(),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
