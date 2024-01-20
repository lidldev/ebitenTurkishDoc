package main

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	rplatformer "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	IdleSprite      *ebiten.Image
	RightSprite     *ebiten.Image
	LeftSprite      *ebiten.Image
	BackgroundImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(rplatformer.MainChar_png))
	if err != nil {
		panic(err)
	}
	IdleSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Right_png))
	if err != nil {
		panic(err)
	}
	RightSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Left_png))
	if err != nil {
		panic(err)
	}
	LeftSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Background_png))
	if err != nil {
		panic(err)
	}
	BackgroundImage = ebiten.NewImageFromImage(img)
}

type char struct {
	x  int
	y  int
	vx int
	vy int
}

const (
	groundY = 380
	unit    = 16
)

const (
	screenWidth  = 960
	screenHeight = 540
)

func (c *char) tryJump() {
	if c.y == groundY*unit {
		c.vy = -10 * unit
	}
}

func (c *char) update() {
	c.x += c.vx
	c.y += c.vy

	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.vx > 0 {
		c.vx -= 2
	} else if c.vx < 0 {
		c.vx += 2
	}
	if c.vy < 20*unit {
		c.vy += 8
	}
}

func (c *char) draw(screen *ebiten.Image) {
	s := IdleSprite
	if c.vx > 0 {
		s = RightSprite
	} else if c.vx < 0 {
		s = LeftSprite
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(s, op)
}

type Game struct {
	gopher *char
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(BackgroundImage, op)

	g.gopher.draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	if g.gopher == nil {
		g.gopher = &char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.gopher.vx = -5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.gopher.vx = 5 * unit
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.gopher.tryJump()
	}

	g.gopher.update()
	return nil
}

func main() {
	ebiten.SetWindowTitle("Ebiten Türkçe Dökümantasyon")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
