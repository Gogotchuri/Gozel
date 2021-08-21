package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
)

func CreateShader(name, vertexSrc, fragmentSrc string) (render_types.Shader, error) {
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil, fmt.Errorf("platform none")
	case render_types.OpenGL:
		return opengl.CreateGLShader(name, vertexSrc, fragmentSrc)
	}
	return nil, fmt.Errorf("platform is undefined")
}

func CreateShaderFromFile(filename string) (render_types.Shader, error) {
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil, fmt.Errorf("platform none")
	case render_types.OpenGL:
		return opengl.CreateGLShaderFromFile(filename)
	}
	return nil, fmt.Errorf("platform is undefined")
}

