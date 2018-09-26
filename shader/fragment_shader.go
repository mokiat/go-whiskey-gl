package shader

import (
	"errors"

	"github.com/go-gl/gl/v3.2-core/gl"
)

type FragmentShader struct {
	Shader
}

func NewFragmentShader() *FragmentShader {
	return &FragmentShader{
		Shader: DefaultShader(),
	}
}

func (s *FragmentShader) Allocate() error {
	s.ID = gl.CreateShader(gl.FRAGMENT_SHADER)
	if s.ID == InvalidID {
		return errors.New("Failed to allocate shader!")
	}
	return nil
}
