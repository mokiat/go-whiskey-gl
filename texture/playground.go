package texture

type FlatDataPlayground interface {
	Width() int
	Height() int
	Data() []byte
}

type RGBAFlatDataPlayground interface {
	FlatDataPlayground
	SetTexel(x, y int, rgba RGBA)
	Texel(x, y int) RGBA
	SetData([]byte)
}

func DedicatedRGBAFlatDataPlayground(width, height int) RGBAFlatDataPlayground {
	return &rgbaFlatDataPlayground{
		width:  width,
		height: height,
		data:   make([]byte, width*height*4),
	}
}

type RGBA struct {
	R byte
	G byte
	B byte
	A byte
}

type rgbaFlatDataPlayground struct {
	width  int
	height int
	data   []byte
}

func (p *rgbaFlatDataPlayground) Width() int {
	return p.width
}

func (p *rgbaFlatDataPlayground) Height() int {
	return p.height
}

func (p *rgbaFlatDataPlayground) Data() []byte {
	return p.data
}

func (p *rgbaFlatDataPlayground) SetTexel(x, y int, rgba RGBA) {
	offset := (y*p.width + x) * 4
	p.data[offset+0] = rgba.R
	p.data[offset+1] = rgba.G
	p.data[offset+2] = rgba.B
	p.data[offset+3] = rgba.A
}

func (p *rgbaFlatDataPlayground) Texel(x, y int) RGBA {
	offset := (y*p.width + x) * 4
	return RGBA{
		R: p.data[offset+0],
		G: p.data[offset+1],
		B: p.data[offset+2],
		A: p.data[offset+3],
	}
}

func (p *rgbaFlatDataPlayground) SetData(data []byte) {
	if len(data) != p.width*p.height*4 {
		panic("Invalid data size!")
	}
	if copy(p.data, data) != p.width*p.height*4 {
		panic("Incorrect source data size!")
	}
}
