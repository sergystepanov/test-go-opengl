package app

import (
	"log"
	"time"
	"unsafe"

	"github.com/sergystepanov/test-go-opengl/internal/gl"
	"github.com/sergystepanov/test-go-opengl/internal/glut"
	"github.com/sergystepanov/test-go-opengl/internal/thread"
)

func StartGlut() {
	thread.Call(MainGlut)
}

func MainGlut() {
	log.Printf("[Start]")
	waitSec := 1.0
	ww, wh := 320, 240

	glut.Init()

	glut.InitDisplayMode(glut.Single | glut.RGB)
	glut.InitWindowSize(ww, wh)
	wid := glut.CreateWindow("GLUT")
	defer glut.DestroyWindow(wid)
	if err := iGL(ww, wh); err != nil {
		panic(err)
	}
	running := true
	pixels := make([]byte, ww*wh*3)
	glut.DisplayFunc(func() {
		triforce()
		glut.SwapBuffers()
		if !running {
			gl.ReadPixels(0, 0, int32(ww), int32(wh), gl.RGB, gl.UNSIGNED_BYTE, unsafe.Pointer(&pixels[0]))
			toRGBA(pixels, ww, wh, "./glut.png")
		}
	})
	glut.MainLoopEvent()
	go glut.MainLoop()

	start := time.Now()
	for running {
		if time.Since(start).Seconds() > waitSec {
			running = false
			glut.PostRedisplay()
			glut.MainLoopEvent()
		}
	}
	log.Printf("[End]")
}
