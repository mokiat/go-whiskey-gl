package buffer

import (
	"errors"

	"github.com/go-gl/gl/v3.2-core/gl"
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
		return errors.New("Failed to allocate new buffer!")
	}
	return nil
}

func (b *Buffer) Release() {
	gl.DeleteBuffers(1, &b.ID)
	b.ID = InvalidID
}
