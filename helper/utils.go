package utils

import (
	"errors"
	"github.com/beacon-chain/types"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/blake2s"
)

const MAX_VALIDATORS = 2 ^ 24

// Shuffle shuffles the validator set based on same seed, this algorithm will be used to select
// shard committees proposer and attesters for blocks.
func Shuffle(seed common.Hash, validatorCount int) ([]types.ValidatorRecord, error) {
	if validatorCount <= MAX_VALIDATORS {
		return nil, errors.New("not enough validators to shuffle")
	}
	validators := make([]types.ValidatorRecord, validatorCount)

	for i := 0; i < validatorCount; i++ {
		hashSeed, _ := blake2s.New256(seed[:])
		hashSeedByte := hashSeed.Sum(nil)
		for j := 0; j < 30; j = j + 3 {
			numToSwap := int(hashSeedByte[j] + hashSeedByte[j+1] + hashSeedByte[j+2])
			remaining := validatorCount - i
			if remaining == 0 {
				break
			}
			posToSwap := numToSwap%remaining + i
			validators[i], validators[posToSwap] = validators[posToSwap], validators[i]
		}
	}
}
