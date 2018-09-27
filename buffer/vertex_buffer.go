package buffer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mokiat/go-whiskey/logging"
)

type VertexBuffer struct {
	Buffer
}

func NewVertexBuffer() *VertexBuffer {
	return &VertexBuffer{
		Buffer: DefaultBuffer(),
	}
}

func (b *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, b.ID)
}

func (b *VertexBuffer) CreateData(playground DataPlayground) {
	logging.Printf("creating vertex buffer data (size: %d)", playground.Size())
	gl.BufferData(gl.ARRAY_BUFFER, playground.Size(), gl.Ptr(playground.Data()), gl.STATIC_DRAW)
}
