package core

/// A Merkle proof that proves a particular element exists in the MMR.
// #[derive(Serialize, Deserialize, Debug, Eq, PartialEq, Clone, PartialOrd, Ord)]
type MerkleProof struct {
  /// The size of the MMR at the time the proof was created.
  Mmr_size uint64
  /// The sibling path from the leaf up to the final sibling hashing to the
  /// root.
  Path []Hash
}
