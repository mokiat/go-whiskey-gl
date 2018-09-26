package shader

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.2-core/gl"
)

type Program struct {
	ID uint32
}

func NewProgram() *Program {
	return &Program{
		ID: InvalidID,
	}
}

func (p *Program) Allocate() error {
	p.ID = gl.CreateProgram()
	if p.ID == InvalidID {
		return errors.New("Failed to allocate program!")
	}
	return nil
}

func (p *Program) AttachVertexShader(shader *VertexShader) {
	gl.AttachShader(p.ID, shader.ID)
}

func (p *Program) AttachFragmentShader(shader *FragmentShader) {
	gl.AttachShader(p.ID, shader.ID)
}

func (p *Program) LinkProgram() error {
	gl.LinkProgram(p.ID)

	if p.getLinkStatus() == gl.FALSE {
		log := p.getProgramLog()
		return fmt.Errorf("Failed to link program: %s", log)
	}

	return nil
}

func (p *Program) getLinkStatus() int32 {
	var status int32
	gl.GetProgramiv(p.ID, gl.LINK_STATUS, &status)
	return status
}

func (p *Program) getProgramLog() string {
	var logLength int32
	gl.GetProgramiv(p.ID, gl.INFO_LOG_LENGTH, &logLength)

	log := strings.Repeat("\x00", int(logLength+1))
	gl.GetProgramInfoLog(p.ID, logLength, nil, gl.Str(log))
	return log
}

func (p *Program) GetAttributeLocation(name string) uint32 {
	location := gl.GetAttribLocation(p.ID, gl.Str(name+"\x00"))
	// we do the casting since all other gl functions expect uint32
	return uint32(location)
}

func (p *Program) GetUniformLocation(name string) int32 {
	return gl.GetUniformLocation(p.ID, gl.Str(name+"\x00"))
}

func (p *Program) Use() {
	gl.UseProgram(p.ID)
}

func (p *Program) Release() {
	gl.DeleteProgram(p.ID)
	p.ID = InvalidID
}
