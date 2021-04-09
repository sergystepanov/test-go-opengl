package app

import (
	"log"
	"runtime"
	"time"

	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func StartGLFW() {
	mainthread.Call(MainGLFW)
}

func MainGLFW() {
	log.Printf("[Start]")
	waitSec := 3.0

	var window *glfw.Window

	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	if runtime.GOOS == "darwin" {
		log.Printf("Oh no, it seems that you are running it on macOS")
		//glfw.WindowHint(glfw.StencilBits, 8)
		//glfw.WindowHint(glfw.RedBits, 8)
		//glfw.WindowHint(glfw.GreenBits, 8)
		//glfw.WindowHint(glfw.BlueBits, 8)
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	}

	window, err := glfw.CreateWindow(320, 240, "Test", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	defer window.Destroy()

	if err = gl.Init(); err != nil {
		panic(err)
	}

	printGlInfo()

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	w, h := window.GetSize()
	gl.Viewport(0, 0, int32(w), int32(h))

	start := time.Now()
	running := true
	for running && !window.ShouldClose() {
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

		window.SwapBuffers()
		if time.Since(start).Seconds() > waitSec {
			running = false
		}
	}

	log.Printf("[End]")
}
