package buffer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mokiat/go-whiskey/logging"
)

type IndexBuffer struct {
	Buffer
}

func NewIndexBuffer() *IndexBuffer {
	return &IndexBuffer{
		Buffer: DefaultBuffer(),
	}
}

func (b *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, b.ID)
}

func (b *IndexBuffer) CreateData(playground DataPlayground) {
	logging.Printf("creating index buffer data (size: %d)", playground.Size())
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, playground.Size(), gl.Ptr(playground.Data()), gl.STATIC_DRAW)
}
