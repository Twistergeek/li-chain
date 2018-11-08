package blockchain

import (
	"../block"
)

type Blockchain struct {
	Blocks []*block.Block
}
