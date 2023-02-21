package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const NUM_PLANETS = 10

const MIN_RADIUS = 20
const MAX_RADIUS = 50

const MAX_START_SPEED = 5

const GRAVITY_CONST = 50
const IGNORE_DIST = 200

const WIDTH = 640
const HEIGHT = 480

var BG_COLOR = color.RGBA{0, 127, 255, 255}

type Game struct {
	planets []*Planet
}

func (g *Game) Update() error {
	for _, planet := range g.planets {
		planet.Update(g)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BG_COLOR)
	for _, planet := range g.planets {
		ebitenutil.DrawCircle(screen, planet.Position.X, planet.Position.Y, planet.Mass, planet.Color)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func InitGame() Game {
	planets := []*Planet{}

	for i := 0; i < NUM_PLANETS; i++ {
		planet := Random()
		planets = append(planets, &planet)
	}

	game := Game{planets: planets}

	return game
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(WIDTH, HEIGHT)

	ebiten.SetWindowTitle("Planets")

	game := InitGame()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
