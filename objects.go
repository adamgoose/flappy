package main

import (
	"github.com/adamgoose/flappy/components"
	"github.com/adamgoose/flappy/lib"
	"github.com/veandco/go-sdl2/sdl"
)

func NewPlayer() *lib.Object {
	p := &lib.Object{}
	p.Position = lib.Vector{X: screenWidth / 4, Y: screenHeight / 2}
	p.Active = true

	sr := components.NewSpriteRenderer(p, sprite, sprites["player_1"], scale)
	p.AddComponent(sr)

	f := components.NewFlight(p, screenHeight)
	p.AddComponent(f)

	return p
}

func NewBackground() *lib.Object {
	p := &lib.Object{}
	p.Position = lib.Vector{X: screenWidth / 2, Y: screenHeight / 2}
	p.Active = true

	sr := components.NewSpriteRenderer(p, sprite, sprites["background_night"], scale)
	p.AddComponent(sr)

	return p
}

func NewGetReady() *lib.Object {
	p := &lib.Object{}
	p.Position = lib.Vector{X: screenWidth / 2, Y: screenHeight / 4}
	p.Active = true

	sr := components.NewSpriteRenderer(p, sprite, sprites["get_ready"], scale)
	p.AddComponent(sr)

	return p
}

func NewPipe() *lib.Object {
	p := &lib.Object{}
	p.Position = lib.Vector{X: screenWidth / 4 * 3, Y: screenHeight / 2}
	p.Active = false

	spacing := float64(125)
	csr := components.NewCompoundSpriteRenderer(p, sprite,
		[]*sdl.Rect{sprites["pipe_down"], sprites["pipe_up"]},
		[]lib.Vector{{X: 0, Y: spacing * -1 * scale}, {X: 0, Y: spacing * scale}},
		scale,
	)
	p.AddComponent(csr)

	return p
}
