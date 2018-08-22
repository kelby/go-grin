package api

/// The state of the current fork tip
// #[derive(Serialize, Deserialize, Debug, Clone)]
struct Tip struct {
  /// Height of the tip (max height of the fork)
  Height uint64
  // Last block pushed to the fork
  Last_block_pushed string
  // Block previous to last
  Prev_block_to_last string
  // Total difficulty accumulated on that fork
  Total_difficulty uint64
}

/// Status page containing different server information
// #[derive(Serialize, Deserialize, Debug, Clone)]
type Status struct {
  // The protocol version
  Protocol_version uint32
  // The user user agent
  User_agent string
  // The current number of connections
  Connections uint32
  // The state of the current fork Tip
  Tip Tip
}

/// TxHashSet
// #[derive(Serialize, Deserialize, Debug, Clone)]
type TxHashSet struct {
  /// Output Root Hash
  Output_root_hash string
  // Rangeproof root hash
  Range_proof_root_hash string
  // Kernel set root hash
  Kernel_root_hash string
}

/// Wrapper around a list of txhashset nodes, so it can be
/// presented properly via json
// #[derive(Serialize, Deserialize, Debug, Clone)]
type TxHashSetNode struct {
  // The hash
  Hash string
}

// #[derive(Debug, Serialize, Deserialize, Clone)]
type Output struct {
  /// The output commitment representing the amount
  Commit PrintableCommitment
}

// #[derive(Debug, Clone)]
type PrintableCommitment struct {
  Commit pedersen.Commitment
}

// As above, except formatted a bit better for human viewing
// #[derive(Debug, Clone)]
type OutputPrintable struct {
  /// The type of output Coinbase|Transaction
  Output_type OutputType
  /// The homomorphic commitment representing the output's amount
  /// (as hex string)
  Commit pedersen.Commitment
  /// Whether the output has been spent
  Spent bool
  /// Rangeproof (as hex string)
  Proof string
  /// Rangeproof hash (as hex string)
  Proof_hash string

  Merkle_proof MerkleProof
}

// Printable representation of a block
// #[derive(Debug, Serialize, Deserialize, Clone)]
type TxKernelPrintable struct {
  Features string
  Fee uint64
  Lock_Height uint64
  Excess string
  Excess_Sig string
}

// Just the information required for wallet reconstruction
// #[derive(Debug, Serialize, Deserialize, Clone)]
type BlockHeaderInfo struct {
  // Hash
  Hash string
  /// Height of this block since the genesis block (height 0)
  Height uint64
  /// Hash of the block previous to this in the chain.
  Previous string
}

// #[derive(Debug, Serialize, Deserialize, Clone)]
type BlockHeaderPrintable struct {
  // Hash
  Hash string
  /// Version of the block
  Version uint16
  /// Height of this block since the genesis block (height 0)
  Height uint64
  /// Hash of the block previous to this in the chain.
  Previous string
  /// rfc3339 timestamp at which the block was built.
  Timestamp string
  /// Merklish root of all the commitments in the TxHashSet
  Output_root string
  /// Merklish root of all range proofs in the TxHashSet
  Range_proof_root string
  /// Merklish root of all transaction kernels in the TxHashSet
  Kernel_root string
  /// Nonce increment used to mine this block.
  Nonce uint64
  /// Size of the cuckoo graph
  Cuckoo_size uint8
  Cuckoo_solution []uint64
  /// Total accumulated difficulty since genesis block
  Total_difficulty uint64
  /// Total kernel offset since genesis block
  Total_kernel_offset string
}

// Printable representation of a block
// #[derive(Debug, Serialize, Deserialize, Clone)]
type BlockPrintable struct {
  /// The block header
  Header BlockHeaderPrintable
  // Input transactions
  Inputs []string
  /// A printable version of the outputs
  Outputs []OutputPrintable
  /// A printable version of the transaction kernels
  Kernels []TxKernelPrintable
}

// #[derive(Debug, Serialize, Deserialize, Clone)]
type CompactBlockPrintable struct {
  /// The block header
  Header BlockHeaderPrintable
  /// Full outputs, specifically coinbase output(s)
  Out_full []OutputPrintable
  /// Full kernels, specifically coinbase kernel(s)
  Kern_full []TxKernelPrintable
  /// Kernels (hex short_ids)
  Kern_ids []string
}

// For wallet reconstruction, include the header info along with the
// transactions in the block
// #[derive(Debug, Serialize, Deserialize, Clone)]
type BlockOutputs struct {
  /// The block header
  Header BlockHeaderInfo
  /// A printable version of the outputs
  Outputs []OutputPrintable
}

// For traversing all outputs in the UTXO set
// transactions in the block
// #[derive(Debug, Serialize, Deserialize, Clone)]
type OutputListing struct {
  /// The last available output index
  Highest_index uint64
  /// The last insertion index retrieved
  Last_retrieved_index uint64
  /// A printable version of the outputs
  Outputs []OutputPrintable
}

// #[derive(Serialize, Deserialize)]
type PoolInfo struct {
  /// Size of the pool
  Pool_size usize
}
