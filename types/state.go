package types

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

// ActiveState contains fields of current state of beacon chain,
// it changes every block.
type ActiveState struct {
	Height            uint64                   // Height of the current block.
	Randao            common.Hash              // Randao beacon state from global's point of view.
	FfgVoterBitmask   []byte                   // FfgVoterBitmask records the validators that voted for this epoch as bitfield.
	BalanceDeltas     []uint                   // BalanceDeltas is the deltas to validator balances.
	PartialCrosslinks []PartialCrosslinkRecord // PartialCrosslinks records data about crosslinks in progress.
	TotalSkipCount    uint64                   // TotalSkipCount records total number of skips to determine minimal time stamp.
}

// PartialCrosslinkRecord contains information about cross links
// that are being put together during this epoch.
type PartialCrosslinkRecord struct {
	ShardID        uint16      // ShardID is the shard crosslink being made for.
	ShardBlockHash common.Hash // ShardBlockHash is the hash of the block.
	VoterBitmask   []byte      // VoterBitmask determines which of the eligible voters are voting for it.
}

// CrystallizedState contains fields of every epoch state,
// it changes every epoch.
type CrystallizedState struct {
	ActiveValidators       []ValidatorRecord // ActiveValidators is the list of active validators.
	QueuedValidators       []ValidatorRecord // QueuedValidators is the list of joined but not yet inducted validators.
	ExitedValidators       []ValidatorRecord // ExitedValidators is the list of removed validators pending withdrawal.
	CurrentShuffling       []uint16          // CurrentShuffling is hhe permutation of validators used to determine who cross-links what shard in this epoch.
	CurrentEpoch           uint64            // CurrentEpoch is the current epoch.
	LastJustifiedEpoch     uint64            // LastJustifiedEpoch is the last justified epoch.
	LastFinalizedEpoch     uint64            // LastFinalizedEpoch is the last finalized epoch.
	Dynasty                uint64            // Dynasty is the current dynasty.
	NextShard              uint16            // NextShard is the next shard that cross-linking assignment will start from.
	CurrentCheckpoint      common.Hash       // CurrentCheckpoint is the current FFG checkpoint.
	CrosslinkRecords       []CrosslinkRecord // CrosslinkRecords records about the most recent crosslink for each shard.
	TotalDeposits          uint              // TotalDeposits is the Total balance of deposits.
	CrosslinkSeed          common.Hash       // CrosslinkSeed is used to select the committees for each shard.
	CrosslinkSeedLastReset uint64            // CrosslinkSeedLastReset is the last epoch the crosslink seed was reset.
}

// ValidatorRecord contains information about a validator
type ValidatorRecord struct {
	PubKey           ecdsa.PublicKey // PubKey is the validator's public key.
	ReturnShard      uint16          // ReturnShard is the shard balance will be sent to after withdrawal.
	ReturnAddress    common.Address  // ReturnAddress is the address balance will be sent to after withdrawal.
	RandaoCommitment common.Hash     // RandaoCommitment is validator's current RANDAO beacon commitment.
	Balance          uint64          // Balance is validator's current balance.
	SwitchDynasty    uint64          // SwitchDynasty is the dynasty where the validator can (be inducted | be removed | withdraw their balance).
}

// CrosslinkRecord contains the fields of last fully formed
// crosslink to be submitted into the chain.
type CrosslinkRecord struct {
	Epoch uint64      // Epoch records the epoch the crosslink was submitted in.
	Hash  common.Hash // Hash is the block hash.
}
