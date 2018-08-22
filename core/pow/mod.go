/// Validates the proof of work of a given header, and that the proof of work
/// satisfies the requirements of the header.
func Verify_size(bh &BlockHeader, cuckoo_sz uint8) bool

/// Mines a genesis block using the internal miner
func Mine_genesis_block() (Block, error)

/// Runs a proof of work computation over the provided block using the provided
/// Mining Worker, until the required difficulty target is reached. May take a
/// while for a low target...
func Pow_size(bh &BlockHeader, diff Difficulty, proof_size usize, sz uint8) error
