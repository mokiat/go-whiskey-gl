package buffer

import (
	"errors"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mokiat/go-whiskey/logging"
)

const InvalidID = 0

type Buffer struct {
	ID uint32
}

func DefaultBuffer() Buffer {
	return Buffer{
		ID: InvalidID,
	}
}

func (b *Buffer) Allocate() error {
	gl.GenBuffers(1, &b.ID)
	if b.ID == InvalidID {
		return errors.New("failed to allocate buffer")
	}
	logging.Printf("buffer allocated (id: %d)", b.ID)
	return nil
}

func (b *Buffer) Release() {
	gl.DeleteBuffers(1, &b.ID)
	b.ID = InvalidID
}
