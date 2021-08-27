package opengl

import (
	"Gozel/renderer/render_types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var _ render_types.VertexBuffer = &GLVertexBuffer{}

type DrawHint uint32

const (
	StreamDraw  DrawHint = gl.STREAM_DRAW
	StaticDraw  DrawHint = gl.STATIC_DRAW
	DynamicDraw DrawHint = gl.DYNAMIC_DRAW
)

type GLVertexBuffer struct {
	BufferID uint32
	Count    int
	Layout   render_types.VertexBufferLayout
}

func CreateGLVertexBuffer(vertices []float32, unitSize int, count int, hint DrawHint) *GLVertexBuffer {
	vb := &GLVertexBuffer{
		Count: count,
	}

	gl.GenBuffers(1, &vb.BufferID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.BufferID)
	gl.BufferData(gl.ARRAY_BUFFER, unitSize*count, gl.Ptr(vertices), uint32(hint))

	return vb
}

func (b *GLVertexBuffer) GetLayout() *render_types.VertexBufferLayout {
	return &b.Layout
}

func (b *GLVertexBuffer) SetLayout(vbi render_types.VertexBufferLayout) {
	b.Layout = vbi
}

func (b *GLVertexBuffer) GetCount() int32 {
	return int32(b.Count)
}

func (b *GLVertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, b.BufferID)
}

func (b *GLVertexBuffer) UnBind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, b.BufferID)
}

func (b *GLVertexBuffer) Destroy() {
	gl.DeleteBuffers(1, &b.BufferID)
}

var _ render_types.IndexBuffer = &GLIndexBuffer{}

type GLIndexBuffer struct {
	BufferID uint32
	Count    int
}

func CreateGLIndexBuffer(indices []int32, count int, hint DrawHint) *GLIndexBuffer {
	ib := &GLIndexBuffer{
		Count: count,
	}

	gl.GenBuffers(1, &ib.BufferID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.BufferID)
	//TODO fixed size of 4, might be different
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*count, gl.Ptr(indices), uint32(hint))

	return ib
}

func (ib *GLIndexBuffer) GetCount() int32 {
	return int32(ib.Count)
}

func (ib *GLIndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.BufferID)
}

func (ib *GLIndexBuffer) UnBind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (ib *GLIndexBuffer) Destroy() {
	gl.DeleteBuffers(1, &ib.BufferID)
}
