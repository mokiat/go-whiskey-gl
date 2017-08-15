package shader

import (
	"errors"

	"github.com/go-gl/gl/v2.1/gl"
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
	return nil
}
