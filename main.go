package main

import (
	"github.com/sergystepanov/test-go-opengl/app"
	"github.com/sergystepanov/test-go-opengl/internal/thread"
)

func run() {
	thread.Call(app.MainGLFW)
}

func main() {
	thread.Run(run)
}
