package pow

import (
  "math"
	"math/big"
  "../../block"
)

const targetBits = 16
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *block.Block
	target *big.Int
}
