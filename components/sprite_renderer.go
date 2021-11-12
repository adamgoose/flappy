package components

import (
	"github.com/adamgoose/flappy/lib"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteRenderer struct {
	lib.DefaultBehavior
	Texture *sdl.Texture
	Source  *sdl.Rect
	Scale   float64

	X, Y, H, W float64
}

func NewSpriteRenderer(container *lib.Object, tex *sdl.Texture, src *sdl.Rect, scale float64) *SpriteRenderer {
	return &SpriteRenderer{
		DefaultBehavior: lib.DefaultBehavior{
			Container: container,
		},
		Texture: tex,
		Source:  src,
		Scale:   scale,
		X:       float64(src.X),
		Y:       float64(src.Y),
		H:       float64(src.H),
		W:       float64(src.W),
	}
}

func (sr *SpriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	// Converting coordinates to top left of sprite
	x := sr.Container.Position.X - sr.W/2.0*sr.Scale
	y := sr.Container.Position.Y - sr.H/2.0*sr.Scale

	renderer.CopyEx(
		sr.Texture,
		sr.Source,
		&sdl.Rect{X: int32(x), Y: int32(y), W: sr.Source.W * int32(sr.Scale), H: sr.Source.H * int32(sr.Scale)},
		sr.Container.Rotation,
		&sdl.Point{X: sr.Source.W / 2, Y: sr.Source.H / 2},
		sdl.FLIP_NONE)

	return nil
}
