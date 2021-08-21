package opengl

import (
	"Gozel/renderer/render_types"
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"time"
)

var _ render_types.GraphicsContext = &GLContext{}

type GLContext struct {
	Window *glfw.Window
}

func CreateGLContext(window render_types.Window) (*GLContext, error) {
	return &GLContext{Window: window.GetBaseWindow().(*glfw.Window)}, nil
}

func (c *GLContext) SwapBuffers() {
	now := time.Now()
	c.Window.SwapBuffers()
	fmt.Println("context swap buffer time ", time.Since(now))
}

func (c *GLContext) Init() {
	c.Window.MakeContextCurrent()

}

