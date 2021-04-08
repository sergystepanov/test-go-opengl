package app

import (
	"github.com/faiface/mainthread"
	"os"
	"runtime"
	"testing"
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
