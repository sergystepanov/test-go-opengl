package app

import (
	"os"
	"runtime"
	"testing"

	"github.com/faiface/mainthread"
)

func init() {
	runtime.LockOSThread()
}

func TestMain(m *testing.M) {
	mainthread.Run(func() {
		os.Exit(m.Run())
	})
}

func TestSDL2(t *testing.T) {
	Start()
}
