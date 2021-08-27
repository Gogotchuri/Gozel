package render_types

import 	"github.com/go-gl/gl/v4.6-core/gl"

type ShaderDataType int
type ShaderType int32

const (
	VertexShader   ShaderType = gl.VERTEX_SHADER
	FragmentShader ShaderType = gl.FRAGMENT_SHADER
)


const (
	NoneSDT ShaderDataType = iota
	Float1
	Float2
	Float3
	Float4
	Int1
	Int2
	Int3
	Int4
	Mat2
	Mat3
	Mat4
)

func getShaderDataTypeSize(t ShaderDataType) int32{
	sizeOfFloat := int32(4)
	sizeOfInt := int32(4)
	switch t{
		case Float1:   return sizeOfFloat
		case Float2:   return sizeOfFloat * 2
		case Float3:   return sizeOfFloat * 3
		case Float4:   return sizeOfFloat * 4
		case Int1:     return sizeOfInt
		case Int2:     return sizeOfInt * 2
		case Int3:     return sizeOfInt * 3
		case Int4:     return sizeOfInt * 4
		case Mat2:     return sizeOfFloat * 4 * 4
		case Mat3:     return sizeOfInt * 3 * 3
		case Mat4:     return sizeOfFloat * 4 * 4
		//TODO Others should be added
	}
	panic("shader datatype size not defined")
	return 0
}

func getShaderDataTypeCount(t ShaderDataType) int32{
	switch t{
		case Float1:   return 1
		case Float2:   return 2
		case Float3:   return 3
		case Float4:   return 4
		case Int1:     return 1
		case Int2:     return 2
		case Int3:     return 3
		case Int4:     return 4
		case Mat2:     return 2 * 2
		case Mat3:     return 3 * 3
		case Mat4:     return 4 * 4
		//TODO Others should be added
	}
	panic("shader datatype size not defined")
	return 0
}