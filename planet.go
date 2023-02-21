package main

import (
	"image/color"
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
		} else if self.Position.Distance(other.Position) < 200 {
			continue
		}

		distance := self.Position.Distance(other.Position)
		new_direction := self.Position.Direction(other.Position)

		acceleration := ((other.Mass * self.Mass) / (distance * distance))
		acceleration *= 50
		acceleration /= self.Mass

		self.Velocity.X += new_direction.X * acceleration
		self.Velocity.Y += new_direction.Y * acceleration
	}

	self.Position.Add(self.Velocity)

	if self.Position.X < 0 {
		self.Position.X = float64(game.width)
	} else if self.Position.X > float64(game.width) {
		self.Position.X = 0
	}
	if self.Position.Y < 0 {
		self.Position.Y = float64(game.height)
	} else if self.Position.Y > float64(game.height) {
		self.Position.Y = 0
	}
}
