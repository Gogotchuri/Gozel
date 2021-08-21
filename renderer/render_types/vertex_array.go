package render_types

type VertexArray interface {
	AddVertexBuffer(vi VertexBuffer)
	SetIndexBuffer(ib IndexBuffer)

	Bind()
	UnBind()

	GetVertexBuffers() []VertexBuffer
	GetIndexBuffer()   IndexBuffer

	Destroy()
}