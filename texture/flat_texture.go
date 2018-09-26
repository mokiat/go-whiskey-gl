package texture

import "github.com/go-gl/gl/v3.2-core/gl"

type FlatTexture struct {
	Texture
}

func NewFlatTexture() *FlatTexture {
	return &FlatTexture{}
}

func (t *FlatTexture) Allocate() error {
	if err := t.Texture.Allocate(); err != nil {
		return err
	}
	t.Bind()
	// TODO: Make configurable
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	return nil
}

func (t *FlatTexture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.ID)
}

func (t *FlatTexture) CreateData(playground FlatDataPlayground) {
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(playground.Width()),
		int32(playground.Height()),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(playground.Data()),
	)
}
