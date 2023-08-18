package sign

func NewParam() *Param {
	return &Param{}
}

type Param struct {
}

func (p *Param) Set() *Param {
	return p
}

func (p *Param) Build() *Param {
	return p
}
