/// An edge in the Cuckoo graph, simply references two u64 nodes.
// #[derive(Debug, Copy, Clone, PartialEq, PartialOrd, Eq, Ord, Hash)]
type Edge struct {
  U uint64
  V uint64
}

/// Cuckoo cycle context
type Cuckoo struct {
  Mask uint64
  Size uint64
  V [4]uint64
}

/// Miner for the Cuckoo Cycle algorithm. While the verifier will work for
/// graph sizes up to a u64, the miner is limited to u32 to be more memory
/// compact (so shift <= 32). Non-optimized for now and and so mostly used for
/// tests, being impractical with sizes greater than 2^22.
type Miner struct {
  Easiness uint64
  Proof_size usize
  Cuckoo Cuckoo
  Graph []uint32
  Sizeshift uint8
}
