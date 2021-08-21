package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
)

func CreateWindow(width, height int, title string) (render_types.Window, error) {
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil, fmt.Errorf("platform none")
	case render_types.OpenGL:
		return opengl.CreateGLWindow(width, height, title)
	}
	return nil, fmt.Errorf("platform is undefined")
}
