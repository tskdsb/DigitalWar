package types

type Map struct {
	Blocks []*Block
}

func NewMap(size int) *Map {
	return &Map{
		Blocks: make([]*Block, 0, size),
	}
}

func (m *Map) AddBlock(block ...*Block) {
	m.Blocks = append(m.Blocks, block...)
}
