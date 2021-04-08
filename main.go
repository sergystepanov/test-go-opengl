package main

import (
	"github.com/faiface/mainthread"
	"github.com/sergystepanov/test-go-opengl/app"
)

func run() {
	mainthread.Call(app.MainGLFW)
}

func main() {
	mainthread.Run(run)
}
