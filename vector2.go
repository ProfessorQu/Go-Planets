package main

import "math"

type Vector2 struct {
	X, Y float64
}

func (self Vector2) Distance(other Vector2) float64 {
	diff_x := self.X - other.X
	diff_y := self.Y - other.Y
	return math.Sqrt(diff_x*diff_x + diff_y*diff_y)
}

func (self Vector2) Magnitude() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y)
}

func (self *Vector2) Add(other Vector2) {
	self.X += other.X
	self.Y += other.Y
}

func (self Vector2) Direction(other Vector2) Vector2 {
	v := Vector2{0, 0}
	v.X = other.X - self.X
	v.Y = other.Y - self.Y

	return v.Normalized()
}

func (self Vector2) Normalized() Vector2 {
	v := Vector2{0, 0}
	magnitude := self.Magnitude()
	v.X = self.X / magnitude
	v.Y = self.Y / magnitude

	return v
}
