package lib

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Vector struct {
	X, Y float64
}

type Object struct {
	Position   Vector
	Rotation   float64
	Active     bool
	Components []Component
}

func (o *Object) AddComponent(new Component) {
	for _, existing := range o.Components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	o.Components = append(o.Components, new)
}

func (o *Object) Draw(renderer *sdl.Renderer) error {
	for _, comp := range o.Components {
		err := comp.OnDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Object) Update(delta float64) error {
	for _, comp := range o.Components {
		err := comp.OnUpdate(delta)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Object) GetComponent(withType Component) Component {
	typ := reflect.TypeOf(withType)
	for _, comp := range o.Components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
}
