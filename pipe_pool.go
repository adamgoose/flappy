package main

import "github.com/adamgoose/flappy/lib"

type PipePool struct {
	Game *Game

	pool []int
}

func NewPipePool(g *Game, count int) *PipePool {
	pool := []int{}

	for i := 0; i < count; i++ {
		pool = append(pool, g.PushObject(NewPipe()))
	}

	return &PipePool{
		Game: g,
		pool: pool,
	}
}

func (p *PipePool) GetPipe() (*lib.Object, bool) {
	for _, i := range p.pool {
		pipe := p.Game.Objects[i]
		if !pipe.Active {
			return pipe, true
		}
	}

	return nil, false
}
