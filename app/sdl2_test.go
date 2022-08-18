package app

import (
	"os"
	"runtime"
	"testing"

	"github.com/sergystepanov/test-go-opengl/internal/thread"
)

func init() {
	runtime.LockOSThread()
}

func TestMain(m *testing.M) {
	thread.Run(func() {
		os.Exit(m.Run())
	})
}

func TestSDL2(t *testing.T) {
	Start()
}
