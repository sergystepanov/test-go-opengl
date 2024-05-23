package app

import (
	"runtime"
	"testing"
)

func init() {
	runtime.LockOSThread()
}

func TestGlut(t *testing.T) {
	StartGlut()
}
