package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.2-core/gl"
)

const InvalidID = 0

type Shader struct {
	ID uint32
}

func DefaultShader() Shader {
	return Shader{
		ID: InvalidID,
	}
}

// XXX: Would be nice if this worked with byte-array-based playground
// objects so that allocation is not needed
func (s *Shader) SetSourceCode(source string) {
	sources, free := gl.Strs(source + "\x00")
	defer free()
	gl.ShaderSource(s.ID, 1, sources, nil)
}

func (s *Shader) Compile() error {
	gl.CompileShader(s.ID)

	if s.getCompileStatus() == gl.FALSE {
		log := s.getShaderLog()
		return fmt.Errorf("Failed to compile shader: %s", log)
	}

	return nil
}

func (s *Shader) getCompileStatus() int32 {
	var status int32
	gl.GetShaderiv(s.ID, gl.COMPILE_STATUS, &status)
	return status
}

func (s *Shader) getShaderLog() string {
	var logLength int32
	gl.GetShaderiv(s.ID, gl.INFO_LOG_LENGTH, &logLength)

	log := strings.Repeat("\x00", int(logLength+1))
	gl.GetShaderInfoLog(s.ID, logLength, nil, gl.Str(log))
	return log
}

func (s *Shader) Release() {
	gl.DeleteShader(s.ID)
	s.ID = InvalidID
}
