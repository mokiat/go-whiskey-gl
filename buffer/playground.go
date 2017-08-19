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

type Float32DataWriter interface {
	PutValue(value float32)
	PutValue2(a, b float32)
	PutValue3(a, b, c float32)
	PutValue4(a, b, c, d float32)
}

func NewFloat32DataWriter(playground Float32DataPlayground, stride int) Float32DataWriter {
	return &float32DataWriter{
		playground: playground,
		offset:     0,
		stride:     stride,
	}
}

type float32DataWriter struct {
	playground Float32DataPlayground
	offset     int
	stride     int
}

func (w *float32DataWriter) PutValue(value float32) {
	w.playground.PutFloat32(w.offset, value)
	w.offset += w.stride
}

func (w *float32DataWriter) PutValue2(a, b float32) {
	w.playground.PutFloat32(w.offset+0, a)
	w.playground.PutFloat32(w.offset+1, b)
	w.offset += w.stride
}

func (w *float32DataWriter) PutValue3(a, b, c float32) {
	w.playground.PutFloat32(w.offset+0, a)
	w.playground.PutFloat32(w.offset+1, b)
	w.playground.PutFloat32(w.offset+2, c)
	w.offset += w.stride
}

func (w *float32DataWriter) PutValue4(a, b, c, d float32) {
	w.playground.PutFloat32(w.offset+0, a)
	w.playground.PutFloat32(w.offset+1, b)
	w.playground.PutFloat32(w.offset+2, c)
	w.playground.PutFloat32(w.offset+3, d)
	w.offset += w.stride
}

type UInt16DataPlayground interface {
	DataPlayground
	PutUInt16(index int, value uint16)
}

func DedicatedUInt16DataPlayground(count int) UInt16DataPlayground {
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

type UInt16DataWriter interface {
	PutValue(value uint16)
}

func NewUInt16DataWriter(playground UInt16DataPlayground, stride int) UInt16DataWriter {
	return &uint16DataWriter{
		playground: playground,
		offset:     0,
		stride:     stride,
	}
}

type uint16DataWriter struct {
	playground UInt16DataPlayground
	offset     int
	stride     int
}

func (w *uint16DataWriter) PutValue(value uint16) {
	w.playground.PutUInt16(w.offset, value)
	w.offset += w.stride
}
