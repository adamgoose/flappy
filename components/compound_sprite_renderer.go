package components

import (
	"github.com/adamgoose/flappy/lib"
	"github.com/veandco/go-sdl2/sdl"
)

type CompoundSpriteRenderer struct {
	lib.DefaultBehavior
	Texture   *sdl.Texture
	Sources   []*sdl.Rect
	Positions []lib.Vector
	Scale     float64
}

func NewCompoundSpriteRenderer(container *lib.Object, tex *sdl.Texture, srcs []*sdl.Rect, poss []lib.Vector, scale float64) *CompoundSpriteRenderer {
	return &CompoundSpriteRenderer{
		DefaultBehavior: lib.DefaultBehavior{
			Container: container,
		},
		Texture:   tex,
		Sources:   srcs,
		Positions: poss,
		Scale:     scale,
	}
}

func (r *CompoundSpriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	orig := r.Container.Position
	for i := range r.Sources {
		src := r.Sources[i]
		pos := r.Positions[i]

		x := orig.X + pos.X - float64(src.W)/2*r.Scale
		y := orig.Y + pos.Y - float64(src.H)/2*r.Scale

		renderer.Copy(
			r.Texture,
			src,
			&sdl.Rect{X: int32(x), Y: int32(y), W: src.W * int32(r.Scale), H: src.H * int32(r.Scale)},
		)
	}

	return nil
}
