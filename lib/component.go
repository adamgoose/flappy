package lib

import "github.com/veandco/go-sdl2/sdl"

type Component interface {
	OnUpdate(delta float64) error
	OnDraw(renderer *sdl.Renderer) error
	OnCollide(other *Component)
}

type DefaultBehavior struct {
	Container *Object
}

func (*DefaultBehavior) OnUpdate(delta float64) error {
	return nil
}

func (*DefaultBehavior) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (*DefaultBehavior) OnCollide(other *Component) {
	//
}
