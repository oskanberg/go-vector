package vector

import (
	"math"
	"math/rand"
)

type Vector2D struct {
	X, Y float64
}

func (s *Vector2D) Normalised() *Vector2D {
	mag := s.Magnitude()
	if mag == 0 {
		return &Vector2D{
			X: s.X,
			Y: s.Y,
		}
	}
	return &Vector2D{
		X: s.X / mag,
		Y: s.Y / mag,
	}
}

func (s *Vector2D) Magnitude() float64 {
	l2 := (s.X * s.X) + (s.Y * s.Y)
	if l2 == 0 {
		return 0
	}
	return math.Sqrt(l2)
}

func (s *Vector2D) Subtract(v *Vector2D) *Vector2D {
	return &Vector2D{
		X: s.X - v.X,
		Y: s.Y - v.Y,
	}
}

func (s *Vector2D) Add(v *Vector2D) *Vector2D {
	return &Vector2D{
		X: s.X + v.X,
		Y: s.Y + v.Y,
	}
}

func (s *Vector2D) Dot(v *Vector2D) float64 {
	return (s.X * v.X) + (s.Y * v.Y)
}

func (s *Vector2D) Rotated(radians float64) *Vector2D {
	cr := math.Cos(radians)
	sr := math.Sin(radians)
	return &Vector2D{
		X: (s.X * cr) - (s.Y * sr),
		Y: (s.X * sr) + (s.Y * cr),
	}
}

func (s *Vector2D) Multiplied(by float64) *Vector2D {
	return &Vector2D{
		X: s.X * by,
		Y: s.Y * by,
	}
}

func (s *Vector2D) Divided(by float64) *Vector2D {
	return &Vector2D{
		X: s.X / by,
		Y: s.Y / by,
	}
}

func (s *Vector2D) Wrap(XLimit, YLimit float64) *Vector2D {
	newX := math.Mod(s.X, XLimit)
	if newX < 0 {
		newX = XLimit - newX
	}

	newY := math.Mod(s.Y, YLimit)
	if newY < 0 {
		newY = YLimit - newY
	}

	// newX := math.Min(s.X, XLimit)
	// if newX < 0 {
	// 	newX = 0
	// }

	// newY := math.Min(s.Y, YLimit)
	// if newY < 0 {
	// 	newY = 0
	// }

	return &Vector2D{
		X: newX,
		Y: newY,
	}
}

func NewVector2d(x, y float64) *Vector2D {
	return &Vector2D{
		X: x,
		Y: y,
	}
}

func NewRandomUnitVector() *Vector2D {
	unit := &Vector2D{
		X: 1,
		Y: 0,
	}
	return unit.Rotated(rand.Float64() * 2 * math.Pi)
}
