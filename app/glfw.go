package app

import (
	"log"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sergystepanov/test-go-opengl/internal/thread"
)

func StartGLFW() {
	thread.Call(MainGLFW)
}

func MainGLFW() {
	log.Printf("[Start]")
	waitSec := 3.0

	var window *glfw.Window

	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	//if runtime.GOOS == "darwin" {
	//	log.Printf("Oh no, it seems that you are running it on macOS")
	//	glfw.WindowHint(glfw.StencilBits, 8)
	//	glfw.WindowHint(glfw.RedBits, 8)
	//	glfw.WindowHint(glfw.GreenBits, 8)
	//	glfw.WindowHint(glfw.BlueBits, 8)
	//	glfw.WindowHint(glfw.AlphaBits, 8)
	//	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	//	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	//	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	//	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	//}

	window, err := glfw.CreateWindow(320, 240, "Test", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	defer window.Destroy()

	if err = iGL(window.GetSize()); err != nil {
		panic(err)
	}

	start := time.Now()
	running := true
	for running && !window.ShouldClose() {
		triforce()
		window.SwapBuffers()
		if time.Since(start).Seconds() > waitSec {
			running = false
		}
	}

	log.Printf("[End]")
}
