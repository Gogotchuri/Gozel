package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
)

func CreateGraphicsContext(window render_types.Window) (render_types.GraphicsContext, error){
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil, fmt.Errorf("platform none")
	case render_types.OpenGL:
		return opengl.CreateGLContext(window)
	}
	return nil, fmt.Errorf("platform is undefined")

}