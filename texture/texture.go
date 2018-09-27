package texture

import (
	"errors"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mokiat/go-whiskey/logging"
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
	logging.Printf("texture allocated (id: %d)", t.ID)
	return nil
}

func (t *Texture) Release() {
	gl.DeleteTextures(1, &t.ID)
	t.ID = InvalidID
}
