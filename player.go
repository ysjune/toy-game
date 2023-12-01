package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"toy-game/assets"
)

type Player struct {
	Position Vector
	Sprite   *ebiten.Image
	Rotation float64
}

func NewPlayer() *Player {

	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		Position: pos,
		Sprite:   sprite,
	}
}
func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Rotation += speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.Sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.Rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.Position.X, p.Position.Y)

	screen.DrawImage(p.Sprite, op)
}
