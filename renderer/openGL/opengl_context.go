package opengl

import (
	"Gozel/renderer/render_types"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var _ render_types.GraphicsContext = &GLContext{}

type GLContext struct {
	Window *glfw.Window
}

func CreateGLContext(window render_types.Window) *GLContext {
	return &GLContext{Window: window.GetBaseWindow().(*glfw.Window)}
}

func (c *GLContext) SwapBuffers() {
	c.Window.SwapBuffers()
}

func (c *GLContext) Init() {
	c.Window.MakeContextCurrent()

}
