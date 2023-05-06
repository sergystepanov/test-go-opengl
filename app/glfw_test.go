package app

import (
	"runtime"
	"testing"
)

func init() {
	runtime.LockOSThread()
}

func TestGLFW(t *testing.T) {
	StartGLFW()
}
