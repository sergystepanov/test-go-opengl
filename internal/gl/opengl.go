package gl

/*
#cgo windows CFLAGS: -DTAG_WINDOWS
#cgo darwin CFLAGS: -DTAG_DARWIN
#cgo linux freebsd openbsd CFLAGS: -DTAG_POSIX
#cgo egl,linux egl,freebsd egl,openbsd egl,windows CFLAGS: -DTAG_EGL

#cgo egl,windows LDFLAGS: -lEGL
#cgo egl,darwin  LDFLAGS: -lEGL

#cgo !gles2,darwin        LDFLAGS: -framework OpenGL
#cgo gles2,darwin         LDFLAGS: -lGLESv2
#cgo !gles2,windows       LDFLAGS: -lopengl32
#cgo gles2,windows        LDFLAGS: -lGLESv2
#cgo !egl,linux !egl,freebsd !egl,openbsd pkg-config: gl
#cgo egl,linux egl,freebsd egl,openbsd    pkg-config: egl

// Check the EGL tag first as it takes priority over the platform's default
// configuration of WGL/GLX/CGL.
#if defined(TAG_EGL)
#include <stdlib.h>
#include <EGL/egl.h>
void* GlowGetProcAddress_gl21(const char* name) {
    return eglGetProcAddress(name);
}
#elif defined(TAG_WINDOWS)
#define WIN32_LEAN_AND_MEAN 1
#include <windows.h>
#include <stdlib.h>
static HMODULE ogl32dll = NULL;
void* GlowGetProcAddress_gl21(const char* name) {
    void* pf = wglGetProcAddress((LPCSTR) name);
    if (pf) {
        return pf;
    }
    if (ogl32dll == NULL) {
        ogl32dll = LoadLibraryA("opengl32.dll");
    }
    return GetProcAddress(ogl32dll, (LPCSTR) name);
}
#elif defined(TAG_DARWIN)
#include <stdlib.h>
#include <dlfcn.h>
void* GlowGetProcAddress_gl21(const char* name) {
    return dlsym(RTLD_DEFAULT, name);
}
#elif defined(TAG_POSIX)
#include <stdlib.h>
#include <GL/glx.h>
void* GlowGetProcAddress_gl21(const char* name) {
    return glXGetProcAddress((const GLubyte *) name);
}
#endif

#if defined(_WIN32) && !defined(APIENTRY) && !defined(__CYGWIN__) && !defined(__SCITECH_SNAP__)
#ifndef WIN32_LEAN_AND_MEAN
#define WIN32_LEAN_AND_MEAN 1
#endif

#include <windows.h>

#endif
#ifndef APIENTRY
#define APIENTRY
#endif
#ifndef APIENTRYP
#define APIENTRYP APIENTRY *
#endif
#ifndef GLAPI
#define GLAPI extern
#endif

#include <KHR/khrplatform.h>

typedef unsigned int GLenum;
typedef unsigned int GLbitfield;
typedef khronos_uint8_t GLubyte;
typedef int GLint;
typedef int GLsizei;
typedef khronos_float_t GLfloat;
typedef double GLdouble;
#ifdef __APPLE__
typedef void *GLhandleARB;
#else
typedef unsigned int GLhandleARB;
#endif

typedef void (APIENTRYP GPCLEARCOLOR)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);

typedef void (APIENTRYP GPENABLE)(GLenum cap);

typedef void (APIENTRYP GPCLEARDEPTH)(GLdouble depth);

typedef void (APIENTRYP GPDEPTHFUNC)(GLenum func);

typedef void (APIENTRYP GPVIEWPORT)(GLint x, GLint y, GLsizei width, GLsizei height);

typedef const GLubyte *(APIENTRYP GPGETSTRING)(GLenum name);

typedef void  (APIENTRYP GPCLEAR)(GLbitfield mask);

typedef void  (APIENTRYP GPBEGIN)(GLenum mode);

typedef void  (APIENTRYP GPEND)();

typedef void  (APIENTRYP GPCOLOR3F)(GLfloat red, GLfloat green, GLfloat blue);

typedef void  (APIENTRYP GPVERTEX2F)(GLfloat x, GLfloat y);

static void glowClearColor(GPCLEARCOLOR fnptr, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    (*fnptr)(red, green, blue, alpha);
}

static void glowClearDepth(GPCLEARDEPTH fnptr, GLdouble depth) { (*fnptr)(depth); }

static void glowEnable(GPENABLE fnptr, GLenum cap) { (*fnptr)(cap); }

static void glowDepthFunc(GPDEPTHFUNC fnptr, GLenum func) { (*fnptr)(func); }

static void glowViewport(GPVIEWPORT fnptr, GLint x, GLint y, GLsizei width, GLsizei height) {
    (*fnptr)(x, y, width, height);
}

static const GLubyte *glowGetString(GPGETSTRING fnptr, GLenum name) {
    return (*fnptr)(name);
}

static void glowClear(GPCLEAR fnptr, GLbitfield mask) { (*fnptr)(mask); }

static void glowBegin(GPBEGIN fnptr, GLenum mode) {
    (*fnptr)(mode);
}

static void glowEnd(GPEND fnptr) { (*fnptr)(); }

static void glowColor3f(GPCOLOR3F fnptr, GLfloat red, GLfloat green, GLfloat blue) {
    (*fnptr)(red, green, blue);
}

static void glowVertex2f(GPVERTEX2F fnptr, GLfloat x, GLfloat y) {
    (*fnptr)(x, y);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

const (
	DEPTH_TEST               = 0x0B71
	LEQUAL                   = 0x0203
	VENDOR                   = 0x1F00
	VERSION                  = 0x1F02
	RENDERER                 = 0x1F01
	SHADING_LANGUAGE_VERSION = 0x8B8C
	COLOR_BUFFER_BIT         = 0x00004000
	DEPTH_BUFFER_BIT         = 0x00000100
	TRIANGLES                = 0x0004
)

var (
	gpClearColor C.GPCLEARCOLOR
	gpClearDepth C.GPCLEARDEPTH
	gpDepthFunc  C.GPDEPTHFUNC
	gpEnable     C.GPENABLE
	gpViewport   C.GPVIEWPORT
	gpGetString  C.GPGETSTRING
	gpClear      C.GPCLEAR
	gpBegin      C.GPBEGIN
	gpEnd        C.GPEND
	gpColor3f    C.GPCOLOR3F
	gpVertex2f   C.GPVERTEX2F
)

func Init() error { return InitWithProcAddrFunc(getProcAddress) }

func InitWithProcAddrFunc(getProcAddr func(name string) unsafe.Pointer) error {
	if gpEnable = (C.GPENABLE)(getProcAddr("glEnable")); gpEnable == nil {
		return errors.New("glEnable")
	}
	if gpClearColor = (C.GPCLEARCOLOR)(getProcAddr("glClearColor")); gpClearColor == nil {
		return errors.New("glClearColor")
	}
	if gpClearDepth = (C.GPCLEARDEPTH)(getProcAddr("glClearDepth")); gpClearDepth == nil {
		return errors.New("glClearDepth")
	}
	if gpDepthFunc = (C.GPDEPTHFUNC)(getProcAddr("glDepthFunc")); gpDepthFunc == nil {
		return errors.New("glDepthFunc")
	}
	if gpViewport = (C.GPVIEWPORT)(getProcAddr("glViewport")); gpViewport == nil {
		return errors.New("glViewport")
	}
	if gpGetString = (C.GPGETSTRING)(getProcAddr("glGetString")); gpGetString == nil {
		return errors.New("glGetString")
	}
	if gpClear = (C.GPCLEAR)(getProcAddr("glClear")); gpClear == nil {
		return errors.New("glClear")
	}
	if gpBegin = (C.GPBEGIN)(getProcAddr("glBegin")); gpBegin == nil {
		return errors.New("glBegin")
	}
	if gpEnd = (C.GPEND)(getProcAddr("glEnd")); gpEnd == nil {
		return errors.New("glEnd")
	}
	if gpColor3f = (C.GPCOLOR3F)(getProcAddr("glColor3f")); gpColor3f == nil {
		return errors.New("glColor3f")
	}
	if gpVertex2f = (C.GPVERTEX2F)(getProcAddr("glVertex2f")); gpVertex2f == nil {
		return errors.New("glVertex2f")
	}
	return nil
}

func ClearColor(red float32, green float32, blue float32, alpha float32) {
	C.glowClearColor(gpClearColor, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
func ClearDepth(depth float64) { C.glowClearDepth(gpClearDepth, (C.GLdouble)(depth)) }
func DepthFunc(fn uint32)      { C.glowDepthFunc(gpDepthFunc, (C.GLenum)(fn)) }
func Enable(cap uint32)        { C.glowEnable(gpEnable, (C.GLenum)(cap)) }
func Viewport(x int32, y int32, width int32, height int32) {
	C.glowViewport(gpViewport, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func GetString(name uint32) *uint8 {
	ret := C.glowGetString(gpGetString, (C.GLenum)(name))
	return (*uint8)(ret)
}
func Clear(mask uint32) { C.glowClear(gpClear, (C.GLbitfield)(mask)) }
func Begin(mode uint32) { C.glowBegin(gpBegin, (C.GLenum)(mode)) }
func End()              { C.glowEnd(gpEnd) }
func Color3f(red float32, green float32, blue float32) {
	C.glowColor3f(gpColor3f, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
}
func Vertex2f(x float32, y float32) {
	C.glowVertex2f(gpVertex2f, (C.GLfloat)(x), (C.GLfloat)(y))
}

func getProcAddress(name string) unsafe.Pointer {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.GlowGetProcAddress_gl21(cname)
}

func GoStr(str *uint8) string { return C.GoString((*C.char)(unsafe.Pointer(str))) }
