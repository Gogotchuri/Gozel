package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
)

func CreateVertexBuffer(vertices []float32, unitSize int, count int, hint opengl.DrawHint) render_types.VertexBuffer {
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil
	case render_types.OpenGL:
		return opengl.CreateGLVertexBuffer(vertices, unitSize, count, hint)
	}
	return nil
}

func CreateIndexBuffer(indices []int32, count int, hint opengl.DrawHint) render_types.IndexBuffer {
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil
	case render_types.OpenGL:
		return opengl.CreateGLIndexBuffer(indices, count, hint)
	}
	return nil
}
