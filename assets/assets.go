package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("player.png")

//var MeteorSprites = mustLoadImages("meteors/*.png")
//var LaserSprite = mustLoadImage("laser.png")
//var ScoreFont = mustLoadFont("font.ttf")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
