package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
)

func CreateTexture(texPath string) (render_types.Texture, error) {
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil, fmt.Errorf("platform none")
	case render_types.OpenGL:
		return opengl.CreateGLTexture(texPath)
	}
	return nil, fmt.Errorf("platform is undefined")
}