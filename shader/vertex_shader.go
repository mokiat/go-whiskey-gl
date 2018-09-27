package shader

import (
	"errors"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mokiat/go-whiskey/logging"
)

type VertexShader struct {
	Shader
}

func NewVertexShader() *VertexShader {
	return &VertexShader{
		Shader: DefaultShader(),
	}
}

func (s *VertexShader) Allocate() error {
	s.ID = gl.CreateShader(gl.VERTEX_SHADER)
	if s.ID == InvalidID {
		return errors.New("Failed to allocate shader!")
	}
	logging.Printf("vertex shader allocated (id: %d)", s.ID)
	return nil
}
