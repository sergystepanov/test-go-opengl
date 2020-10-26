package app

import (
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

func Start() {
	mainthread.Call(Main)
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

	if err = gl.Init(); err != nil {
		panic(err)
	}

	printGlInfo()

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	w, h := window.GetSize()
	gl.Viewport(0, 0, w, h)

	start := time.Now()
	running := true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.Begin(gl.TRIANGLES)
		gl.Color3f(1.0, 1.0, 0.0)
		gl.Vertex2f(0.0, 0.75)
		gl.Vertex2f(0.75, -0.75)
		gl.Vertex2f(-0.75, -0.75)
		gl.Color3f(0.0, 0.0, 0.0)
		gl.Vertex2f(-0.40, 0.0)
		gl.Vertex2f(0.0, -0.77)
		gl.Vertex2f(0.40, 0.0)
		gl.End()

		window.GLSwap()
		if time.Since(start).Seconds() > waitSec {
			running = false
		}
	}

	log.Printf("[End]")
}

func printGlInfo() {
	log.Printf("[OpenGL] Version: %v", gl.GoStr(gl.GetString(gl.VERSION)))
	log.Printf("[OpenGL] Vendor: %v", gl.GoStr(gl.GetString(gl.VENDOR)))
	log.Printf("[OpenGL] Renderer: %v", gl.GoStr(gl.GetString(gl.RENDERER)))
	log.Printf("[OpenGL] GLSL Version: %v", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
}
