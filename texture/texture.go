package texture

import (
	"errors"

	"github.com/go-gl/gl/v3.2-core/gl"
)

const InvalidID = 0

type Texture struct {
	ID uint32
}

func DefaultTexture() Texture {
	return Texture{
		ID: InvalidID,
	}
}

func (t *Texture) Allocate() error {
	gl.GenTextures(1, &t.ID)
	if t.ID == InvalidID {
		return errors.New("Failed to allocate texture!")
	}
	return nil
}

func (t *Texture) Release() {
	gl.DeleteTextures(1, &t.ID)
	t.ID = InvalidID
}
