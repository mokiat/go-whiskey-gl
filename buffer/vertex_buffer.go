package buffer

import "github.com/go-gl/gl/v2.1/gl"

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
	gl.BufferData(gl.ARRAY_BUFFER, playground.Size(), gl.Ptr(playground.Data()), gl.STATIC_DRAW)
}
