package main

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func main() {
	var winTitle = "Test"
	var winWidth, winHeight int32 = 160, 120
	var window *sdl.Window
	var context sdl.GLContext
	//var event sdl.Event
	//var running bool
	var err error

	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow(
		winTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight,
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

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, winWidth, winHeight)

	//running = true
	//for running {
	//	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	//		switch t := event.(type) {
	//		case *sdl.QuitEvent:
	//			running = false
	//		case *sdl.MouseMotionEvent:
	//			fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
	//		}
	//	}

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Begin(gl.TRIANGLES)
	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex2f(0.5, 0.0)
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Vertex2f(-0.5, -0.5)
	gl.Color3f(0.0, 0.0, 1.0)
	gl.Vertex2f(-0.5, 0.5)
	gl.End()

	window.GLSwap()
	time.Sleep(3 * time.Second)
	//}
	fmt.Printf("[End]")
}
