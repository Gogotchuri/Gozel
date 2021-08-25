package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
)

func CreateGraphicsContext(window render_types.Window) render_types.GraphicsContext {
	switch render_types.CurrentPlatform {
	case render_types.None:
		fmt.Println("platform none")
		return nil
	case render_types.OpenGL:
		return opengl.CreateGLContext(window)
	}
	fmt.Println("platform undefined")
	return nil

}
