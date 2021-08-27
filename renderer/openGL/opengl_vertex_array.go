package opengl

import (
	"Gozel/renderer/render_types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var _ render_types.VertexArray = &GLVertexArray{}

type GLVertexArray struct {
	vertexArrayID     uint32
	vertexBufferIndex uint32

	vertexBuffers []render_types.VertexBuffer
	indexBuffer   render_types.IndexBuffer
}

func CreateGLVertexArray() *GLVertexArray {
	vaoID := uint32(0)
	gl.GenVertexArrays(1, &vaoID)
	return &GLVertexArray{
		vertexArrayID:     vaoID,
		vertexBufferIndex: 0,
		vertexBuffers:     []render_types.VertexBuffer{},
		indexBuffer:       nil,
	}
}

func (va *GLVertexArray) AddVertexBuffer(vb render_types.VertexBuffer) {
	va.vertexBuffers = append(va.vertexBuffers, vb)
	/*Bind this vao*/
	va.Bind()
	/*Bind vertex/index buffer*/
	vb.Bind()

	elements := vb.GetLayout().GetElements()
	for _, elem := range elements {
		//Enable attrib array
		gl.EnableVertexAttribArray(va.vertexBufferIndex)
		//pass layout
		gl.VertexAttribPointer(va.vertexBufferIndex, elem.ElementCount,
			getGLBaseType(elem.SDType), elem.Normalized, vb.GetLayout().GetStride(), gl.PtrOffset(int(elem.Offset)))
		//Incrementing for every subsequent vertex attribute
		va.vertexBufferIndex++
	}
}

func (va *GLVertexArray) SetIndexBuffer(ib render_types.IndexBuffer) {
	va.Bind()
	ib.Bind()
	va.indexBuffer = ib
}

func (va *GLVertexArray) Bind() {
	gl.BindVertexArray(va.vertexArrayID)
}

func (va *GLVertexArray) UnBind() {
	gl.BindVertexArray(0)
}

func (va *GLVertexArray) GetVertexBuffers() []render_types.VertexBuffer {
	return va.vertexBuffers
}

func (va *GLVertexArray) GetIndexBuffer() render_types.IndexBuffer {
	return va.indexBuffer
}

func (va *GLVertexArray) Destroy() {
	gl.DeleteVertexArrays(1, &va.vertexArrayID)
}

func getGLBaseType(dataType render_types.ShaderDataType) uint32 {
	switch dataType {
	case render_types.Float1:
		return gl.FLOAT
	case render_types.Float2:
		return gl.FLOAT
	case render_types.Float3:
		return gl.FLOAT
	case render_types.Float4:
		return gl.FLOAT
	case render_types.Int1:
		return gl.INT
	case render_types.Int2:
		return gl.INT
	case render_types.Int3:
		return gl.INT
	case render_types.Int4:
		return gl.INT
	case render_types.Mat2:
		return gl.FLOAT
	case render_types.Mat3:
		return gl.FLOAT
	case render_types.Mat4:
		return gl.FLOAT
	}
	panic("Gl Type undefined")
}
