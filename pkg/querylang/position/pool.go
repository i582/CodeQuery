package position

const DefaultBlockSize = 1024

type Pool struct {
	block []Pos
	off   int
}

func NewPool(blockSize int) *Pool {
	return &Pool{
		block: make([]Pos, blockSize),
	}
}

func (p *Pool) Get() *Pos {
	if len(p.block) == 0 {
		return nil
	}

	if len(p.block) == p.off {
		p.block = make([]Pos, len(p.block))
		p.off = 0
	}

	p.off++

	return &p.block[p.off-1]
}
