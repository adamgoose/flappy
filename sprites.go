package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var sprites = map[string]*sdl.Rect{
	"background_day":   {X: 0, Y: 0, H: 256, W: 144},
	"background_night": {X: 146, Y: 0, H: 256, W: 144},
	"get_ready":        {X: 295, Y: 59, H: 25, W: 92},
	"player_0":         {X: 3, Y: 491, H: 12, W: 17},
	"player_1":         {X: 31, Y: 491, H: 12, W: 17},
	"player_2":         {X: 59, Y: 491, H: 12, W: 17},
	"pipe_down":        {X: 56, Y: 323, H: 160, W: 26},
	"pipe_up":          {X: 84, Y: 323, H: 160, W: 26},
}

var sprite *sdl.Texture

func loadSprites(renderer *sdl.Renderer) {
	img, err := sdl.LoadBMP("assets/sprite.bmp")
	if err != nil {
		panic(fmt.Errorf("loading sprite: %v", err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from sprite: %v", err))
	}

	sprite = tex
}
