package app

import (
	"log"
	"time"

	"github.com/sergystepanov/test-go-opengl/internal/thread"
	"github.com/veandco/go-sdl2/sdl"
)

func Start() {
	thread.Call(Main)
}

func Main() {
	log.Printf("[Start]")
	waitSec := 3.0

	var window *sdl.Window
	var context sdl.GLContext
	var event sdl.Event

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Test",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		320, 240,
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	w, h := window.GetSize()
	if err = iGL(int(w), int(h)); err != nil {
		panic(err)
	}

	start := time.Now()
	running := true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		triforce()
		window.GLSwap()
		if time.Since(start).Seconds() > waitSec {
			running = false
		}
	}

	log.Printf("[End]")
}
