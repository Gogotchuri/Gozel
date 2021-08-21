package opengl

import (
	"Gozel/renderer/render_types"
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"image"
	"image/draw"
	"os"
)

var _ render_types.Texture = &GLTexture{}

type GLTexture struct {
	GLTextureID uint32
	Filepath string

	Width, Height int
}

func (t *GLTexture) Bind(slot uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + slot)
	gl.BindTexture(gl.TEXTURE_2D, t.GLTextureID)
}

func (t *GLTexture) UnBind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t *GLTexture) GetWidth() int {
	return t.Width
}

func (t *GLTexture) GetHeight() int {
	return t.Height
}

func (t *GLTexture) Destroy() {
	gl.DeleteTextures(1, &t.GLTextureID)
}

func CreateGLTexture(texPath string) (*GLTexture, error) {
	tex, w, h, err := loadTexture(texPath)
	if err != nil {
		return nil, err
	}
	return &GLTexture{
		GLTextureID: tex,
		Filepath:    texPath,
		Width:       w,
		Height:      h,
	}, nil
}


func loadTexture(path string) (texture uint32, width, height int, err error) {
	imgFile, err := os.Open(path)
	if err != nil {
		return 0,0,0, fmt.Errorf("texture %s not found: %s", path, err.Error())
	}
	defer imgFile.Close()
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return 0,0,0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, 0, 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	gl.GenTextures(1, &texture)

	//Bind
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	//Unbind
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return texture, rgba.Rect.Size().X, rgba.Rect.Size().Y, nil
}