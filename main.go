package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	planets       []*Planet
	width, height int
}

func (g *Game) Update() error {
	for _, planet := range g.planets {
		planet.Update(g)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 127, 255, 255})
	for _, planet := range g.planets {
		ebitenutil.DrawCircle(screen, planet.Position.X, planet.Position.Y, planet.Mass, planet.Color)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func InitGame(width, height int) Game {
	planets := []*Planet{}

	for i := 0; i < 5; i++ {
		x := rand.Float64() * 640
		y := rand.Float64() * 480
		mass := 20.0 + rand.Float64()*30.0
		color := color.RGBA{uint8(rand.Uint32()), uint8(rand.Uint32()), uint8(rand.Uint32()), 255}
		planets = append(planets, &Planet{
			Vector2{x, y},
			Vector2{y / 1000, x / 1000},
			mass, color})
	}

	game := Game{planets: planets, width: width, height: height}

	return game
}

func main() {
	rand.Seed(time.Now().UnixNano())

	width := 800
	height := 600

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Planets")

	game := InitGame(width, height)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
