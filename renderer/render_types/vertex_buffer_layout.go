package render_types

type LayoutElement struct {
	SDType       ShaderDataType
	Name         string
	ElementCount int32
	Normalized   bool
	Size, Offset int32
}

func CreateLayoutElement(SDType ShaderDataType, name string, normalized bool) LayoutElement {
	return LayoutElement{
		SDType:       SDType,
		Name:         name,
		ElementCount: getShaderDataTypeCount(SDType),
		Normalized:   normalized,
		Size:         getShaderDataTypeSize(SDType),
		Offset:       0,
	}
}

type VertexBufferLayout struct {
	stride   int32
	elements []LayoutElement
}

func (vbl *VertexBufferLayout) Push(le LayoutElement)  {
	le.Offset = vbl.stride
	vbl.elements = append(vbl.elements, le)
	vbl.stride += le.Size
}

func (vbl *VertexBufferLayout) GetStride() int32 {
	return vbl.stride
}

func (vbl *VertexBufferLayout) GetElements() []LayoutElement {
	return vbl.elements
}