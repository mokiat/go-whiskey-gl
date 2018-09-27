package shader

import (
	"errors"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mokiat/go-whiskey/logging"
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
	logging.Printf("fragment shader allocated (id: %d)", s.ID)
	return nil
}
