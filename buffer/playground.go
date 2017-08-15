package buffer

import "math"

type DataPlayground interface {
	Size() int
	Data() []byte
}

type Float32DataPlayground interface {
	DataPlayground
	PutFloat32(index int, value float32)
}

func DedicatedFloat32DataPlayground(count int) Float32DataPlayground {
	return &float32DataPlayground{
		data: make([]byte, count*4),
	}
}

type float32DataPlayground struct {
	data []byte
}

func (p *float32DataPlayground) Size() int {
	return len(p.data)
}

func (p *float32DataPlayground) Data() []byte {
	return p.data
}

func (p *float32DataPlayground) PutFloat32(index int, value float32) {
	bits := math.Float32bits(value)
	p.data[index*4+0] = byte(bits)
	p.data[index*4+1] = byte(bits >> 8)
	p.data[index*4+2] = byte(bits >> 16)
	p.data[index*4+3] = byte(bits >> 24)
}

type UInt16Playground interface {
	DataPlayground
	PutUInt16(index int, value uint16)
}

func DedicatedUInt16DataPlayground(count int) UInt16Playground {
	return &uint16DataPlayground{
		data: make([]byte, count*2),
	}
}

type uint16DataPlayground struct {
	data []byte
}

func (p *uint16DataPlayground) Size() int {
	return len(p.data)
}

func (p *uint16DataPlayground) Data() []byte {
	return p.data
}

func (p *uint16DataPlayground) PutUInt16(index int, value uint16) {
	p.data[index*2+0] = byte(value)
	p.data[index*2+1] = byte(value >> 8)
}
