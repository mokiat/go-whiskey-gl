package texture

import "github.com/go-gl/gl/v2.1/gl"

type CubeTexture struct {
	Texture
}

func NewCubeTexture() *CubeTexture {
	return &CubeTexture{}
}

func (t *CubeTexture) Allocate() error {
	if err := t.Texture.Allocate(); err != nil {
		return err
	}
	t.Bind()
	// TODO: Make configurable
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	return nil
}

func (t *CubeTexture) Bind() {
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, t.ID)
}

func (t *CubeTexture) CreateData(playground CubeDataPlayground) {
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

	setImage := func(target uint32, side CubeSide) {
		imgData := playground.Data(side)
		gl.TexImage2D(
			target,
			0,
			gl.RGBA,
			int32(playground.Size()),
			int32(playground.Size()),
			0,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			gl.Ptr(imgData),
		)
	}
	setImage(gl.TEXTURE_CUBE_MAP_POSITIVE_X, CubeSideRight)
	setImage(gl.TEXTURE_CUBE_MAP_NEGATIVE_X, CubeSideLeft)
	setImage(gl.TEXTURE_CUBE_MAP_POSITIVE_Z, CubeSideFront)
	setImage(gl.TEXTURE_CUBE_MAP_NEGATIVE_Z, CubeSideBack)
	// top and bottom are flipped due to opengl's renderman issue
	setImage(gl.TEXTURE_CUBE_MAP_POSITIVE_Y, CubeSideBottom)
	setImage(gl.TEXTURE_CUBE_MAP_NEGATIVE_Y, CubeSideTop)
}
