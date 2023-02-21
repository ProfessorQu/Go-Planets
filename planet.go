package main

import (
	"image/color"
	"math/rand"
)

type Planet struct {
	Position Vector2
	Velocity Vector2
	Mass     float64
	Color    color.Color
}

func (self *Planet) Update(game *Game) {
	for _, other := range game.planets {
		if other == self {
			continue
		} else if self.Position.Distance(other.Position) < IGNORE_DIST {
			continue
		}

		// Calculating the acceleration
		distance := self.Position.Distance(other.Position)
		new_direction := self.Position.Direction(other.Position)

		acceleration := ((other.Mass * self.Mass) / (distance * distance))
		acceleration *= GRAVITY_CONST
		acceleration /= self.Mass

		// Apply the acceleration
		self.Velocity.X += new_direction.X * acceleration
		self.Velocity.Y += new_direction.Y * acceleration
	}

	self.Position.Add(self.Velocity)

	// Wrapping around the screen
	if self.Position.X < 0 {
		self.Position.X = float64(WIDTH)
	} else if self.Position.X > float64(WIDTH) {
		self.Position.X = 0
	}
	if self.Position.Y < 0 {
		self.Position.Y = float64(HEIGHT)
	} else if self.Position.Y > float64(HEIGHT) {
		self.Position.Y = 0
	}
}

func Random() Planet {
	// Random position
	x := rand.Float64() * float64(WIDTH)
	y := rand.Float64() * float64(HEIGHT)

	// Random velocity
	vel_x := -MAX_START_SPEED + rand.Float64()*(2*MAX_START_SPEED)
	vel_y := -MAX_START_SPEED + rand.Float64()*(2*MAX_START_SPEED)

	// Random mass and color
	mass := MIN_RADIUS + rand.Float64()*(MAX_RADIUS-MIN_RADIUS)
	color := color.RGBA{
		uint8(rand.Uint32()),
		uint8(rand.Uint32()),
		uint8(rand.Uint32()),
		255}

	return Planet{
		Vector2{x, y},
		Vector2{vel_x, vel_y},
		mass, color}

}
