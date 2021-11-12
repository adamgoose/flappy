package components

import (
	"math"

	"github.com/adamgoose/flappy/lib"
	"github.com/veandco/go-sdl2/sdl"
)

type Flight struct {
	lib.DefaultBehavior
	Active       bool
	RotUpSpeed   float64
	RotDownSpeed float64
	Gravity      float64
	Thrust       float64
	ScreenHeight float64
}

func NewFlight(container *lib.Object, screenHeight float64) *Flight {
	return &Flight{
		DefaultBehavior: lib.DefaultBehavior{
			Container: container,
		},
		Active:       false,
		ScreenHeight: screenHeight,
		RotUpSpeed:   20,
		RotDownSpeed: -5,
		Gravity:      -10,
		Thrust:       25,
	}
}

func (f *Flight) OnUpdate(delta float64) error {
	if f.Active == false {
		return nil
	}

	keys := sdl.GetKeyboardState()

	dy := f.Gravity
	if keys[sdl.SCANCODE_SPACE] == 1 {
		dy += f.Thrust
	}

	f.Container.Position.Y = cap(f.Container.Position.Y-dy*delta, 0, f.ScreenHeight)

	dr := f.RotDownSpeed
	if dy >= 0 {
		dr += f.RotUpSpeed
	}
	f.Container.Rotation = cap(f.Container.Rotation-dr*delta, -45, 45)

	return nil
}

func cap(val, min, max float64) float64 {
	return math.Min(math.Max(val, min), max)
}
