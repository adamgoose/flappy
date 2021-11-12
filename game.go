package main

import (
	"fmt"

	"github.com/adamgoose/flappy/components"
	"github.com/adamgoose/flappy/lib"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	R        *sdl.Renderer
	Active   bool
	Score    int
	Objects  []*lib.Object
	PipePool *PipePool

	player   int
	getReady int
}

func NewGame(renderer *sdl.Renderer) *Game {
	g := &Game{
		R:      renderer,
		Active: false,
		Score:  0,
	}

	g.PushObject(NewBackground())
	g.player = g.PushObject(NewPlayer())
	g.getReady = g.PushObject(NewGetReady())

	g.PipePool = NewPipePool(g, 5)

	return g
}

func (g *Game) PushObject(o *lib.Object) int {
	id := len(g.Objects)
	g.Objects = append(g.Objects, o)
	return id
}

func (g *Game) Start() {
	p := g.Objects[g.player]
	p.GetComponent(&components.Flight{}).(*components.Flight).Active = true

	gr := g.Objects[g.getReady]
	gr.Active = false

	g.Active = true
}

func (g *Game) Update(delta float64) (err error) {
	keys := sdl.GetKeyboardState()

	if !g.Active && keys[sdl.SCANCODE_SPACE] == 1 {
		g.Start()
	}

	for _, o := range g.Objects {
		if o.Active {
			err = o.Update(delta)
			if err != nil {
				fmt.Println("updating object:", err)
				return
			}
			err = o.Draw(g.R)
			if err != nil {
				fmt.Println("drawing object:", o)
				return
			}
		}
	}
	return nil
}
