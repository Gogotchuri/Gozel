package opengl

import (
	"Gozel/renderer/render_types"
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var _ render_types.Window = &GLWindow{}

type GLWindow struct {
	GLFWWindow *glfw.Window
	Context    render_types.GraphicsContext
	vsync      bool
}

func CreateGLWindow(width, height int, title string) (*GLWindow, error) {
	/* Initialize the library */
	err := glfw.Init()
	if err != nil {
		return nil, err
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	/* Create a windowed mode window and its OpenGL context */
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, fmt.Errorf("GLFW window initialization failed")
	}

	/* Make the window's context current */
	//window.MakeContextCurrent()
	gc := &GLContext{window}
	gc.Init()
	glfw.SwapInterval(0)

	/*Viewports for scaling and callback*/
	cb := func(w *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	}
	window.SetFramebufferSizeCallback(cb)

	return &GLWindow{GLFWWindow: window, Context: gc}, err
}

func (w *GLWindow) IsOpen() bool {
	return !w.GLFWWindow.ShouldClose()
}

func (w *GLWindow) Close() {
	w.GLFWWindow.SetShouldClose(true)
	glfw.Terminate()
}

func (w *GLWindow) OnUpdate() {
	w.Context.SwapBuffers()
}

func (w *GLWindow) GetBaseWindow() interface{} {
	return w.GLFWWindow
}

func (w *GLWindow) IsVSync() bool {
	return w.vsync
}

func (w *GLWindow) SetVSync(b bool) {
	interval := 0
	if b {
		interval = 1
	}
	glfw.SwapInterval(interval)
	w.vsync = b
}
