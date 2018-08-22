package core

import "fmt"
import "time"

type BlockHeader struct {
	/// Version of the block
	Version uint16
	/// Height of this block since the genesis block (height 0)
	Height uint64
	/// Hash of the block previous to this in the chain.
	Previous string
	/// Timestamp at which the block was built.
	Timestamp time.Time
	/// Total accumulated difficulty since genesis block
	Total_difficulty string
	/// Merklish root of all the commitments in the TxHashSet
	Output_root string
	/// Merklish root of all range proofs in the TxHashSet
	Range_proof_root string
	/// Merklish root of all transaction kernels in the TxHashSet
	Kernel_root string
	/// Total accumulated sum of kernel offsets since genesis block.
	/// We can derive the kernel offset sum for *this* block from
	/// the total kernel offset of the previous block header.
	Total_kernel_offset string
	/// Total accumulated sum of kernel commitments since genesis block.
	/// Should always equal the UTXO commitment sum minus supply.
	Total_kernel_sum string
	/// Total size of the output MMR after applying this block
	Output_mmr_size uint64
	/// Total size of the kernel MMR after applying this block
	Kernel_mmr_size uint64
	/// Nonce increment used to mine this block.
	Nonce uint64
	/// Proof of work data.
	Pow string
}

func Default() BlockHeader {
	return BlockHeader{
		Version:             1,
		Height:              0,
		Previous:            "ZERO_HASH",
		Timestamp:           time.Now(),
		Total_difficulty:    "Difficulty::one()",
		Output_root:         "ZERO_HASH",
		Range_proof_root:    "ZERO_HASH",
		Kernel_root:         "ZERO_HASH",
		Total_kernel_offset: "BlindingFactor::zero()",
		Total_kernel_sum:    "Commitment::from_vec(vec![0; 33])",
		Output_mmr_size:     0,
		Kernel_mmr_size:     0,
		Nonce:               0,
		Pow:                 "Proof::zero(proof_size),"}
}

/// Total kernel offset for the chain state up to and including this block.
func (self *BlockHeader) Total_kernel_offset() -> BlindingFactor {
	self.Total_kernel_offset
}

/// A block as expressed in the MimbleWimble protocol. The reward is
/// non-explicit, assumed to be deducible from block height (similar to
/// bitcoin's schedule) and expressed as a global transaction fee (added v.H),
/// additive to the total of fees ever collected.
// #[derive(Debug, Clone)]
type Block struct {
	/// The header with metadata and commitments to the rest of the data
	Header BlockHeader
	/// List of transaction inputs
	Inputs []Input
	/// List of transaction outputs
	Outputs []Output
	/// List of kernels with associated proofs (note these are offset from
	/// tx_kernels)
	Kernels []TxKernel
}

/// Sum of all fees (inputs less outputs) in the block
func (self *Block) Total_fees() uint64 {
	self.Kernels.iter().map(|p| p.fee).sum()

	total_fees := 0

	for _, kernel := range self.Kernels {
		total_fees += kernel.Fee
	}

	return total_fees
}
