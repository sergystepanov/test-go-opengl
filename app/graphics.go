package app

import (
	"log"

	"github.com/sergystepanov/test-go-opengl/internal/gl"
)

func iGL(w, h int) error {
	if err := gl.Init(); err != nil {
		return err
	}

	printGlInfo()

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, int32(w), int32(h))
	return nil
}

func triforce() {
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
}

func printGlInfo() {
	log.Printf("[OpenGL] Version: %v", gl.GoStr(gl.GetString(gl.VERSION)))
	log.Printf("[OpenGL] Vendor: %v", gl.GoStr(gl.GetString(gl.VENDOR)))
	log.Printf("[OpenGL] Renderer: %v", gl.GoStr(gl.GetString(gl.RENDERER)))
	log.Printf("[OpenGL] GLSL Version: %v", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
}
