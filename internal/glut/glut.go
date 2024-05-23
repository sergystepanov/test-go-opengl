package glut

/*
#cgo windows pkg-config: freeglut
#cgo linux LDFLAGS: -lGL -lGLU -lglut
#cgo darwin pkg-config: glut
#cgo darwin LDFLAGS: -framework OpenGL

#include <stdlib.h>
#include <GL/freeglut.h>

extern void goDisplay();

static void register_display() {
	glutDisplayFunc(&goDisplay);
}

static void unregister_display() {
	glutDisplayFunc(0);
}
*/
import "C"
import (
	"os"
	"strings"
	"unsafe"
)

const RGB = C.GLUT_RGB
const Single = C.GLUT_SINGLE

type tWindowCallbacks struct {
	display func()
}

var (
	windowCallbacks = make(map[int]*tWindowCallbacks)
)

func Init() {
	argc := C.int(len(os.Args))
	argv := make([]*C.char, argc)
	argv_o := make([]*C.char, argc)

	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
	}
	copy(argv_o, argv)

	defer func() {
		if int(argc) < len(os.Args) {
			for i, j := 0, 0; i < int(argc); i++ {
				argv_c := C.GoString(argv[i])

				for strings.Compare(os.Args[j], argv_c) != 0 {
					j = j + 1
				}
				os.Args[i] = os.Args[j]
				j = j + 1
			}
			os.Args = os.Args[:argc]
		}
		for _, arg := range argv_o {
			C.free(unsafe.Pointer(arg))
		}
	}()

	C.glutInit(&argc, &argv[0])
}

func InitDisplayMode(mode uint) {
	C.glutInitDisplayMode(C.uint(mode))
}

func InitWindowSize(width, height int) {
	C.glutInitWindowSize(C.int(width), C.int(height))
}

func CreateWindow(title string) (windowId int) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	windowId = int(C.glutCreateWindow(ctitle))
	windowCallbacks[windowId] = new(tWindowCallbacks)

	return windowId
}

func DestroyWindow(windowId int) {
	C.glutDestroyWindow(C.int(windowId))
	delete(windowCallbacks, windowId)
}

func DisplayFunc(display func()) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].display = display
	if display != nil {
		C.register_display()
	} else {
		C.unregister_display()
	}
}

func MainLoop()      { C.glutMainLoop() }
func MainLoopEvent() { C.glutMainLoopEvent() }
func PostRedisplay() { C.glutPostRedisplay() }
func SwapBuffers()   { C.glutSwapBuffers() }

//export goDisplay
func goDisplay() {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].display()
}
