package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
)

func CreateVertexBuffer(vertices []float32,  unitSize int, count int, hint opengl.DrawHint) (render_types.VertexBuffer, error) {
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil, fmt.Errorf("platform none")
	case render_types.OpenGL:
		return opengl.CreateGLVertexBuffer(vertices, unitSize, count, hint)
	}
	return nil, fmt.Errorf("platform is undefined")
}

func CreateIndexBuffer(indices []int32, count int, hint opengl.DrawHint) (render_types.IndexBuffer, error) {
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil, fmt.Errorf("platform none")
	case render_types.OpenGL:
		return opengl.CreateGLIndexBuffer(indices, count, hint)
	}
	return nil, fmt.Errorf("platform is undefined")
}