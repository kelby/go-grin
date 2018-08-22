package core

import "target"
/// A Cuckoo Cycle proof of work, consisting of the shift to get the graph
/// size (i.e. 31 for Cuckoo31 with a 2^31 or 1<<31 graph size) and the nonces
/// of the graph solution. While being expressed as u64 for simplicity, each
/// nonce is strictly less than half the cycle size (i.e. <2^30 for Cuckoo 31).
///
/// The hash of the `Proof` is the hash of its packed nonces when serializing
/// them at their exact bit size. The resulting bit sequence is padded to be
/// byte-aligned.
///
// #[derive(Clone, PartialOrd, PartialEq)]
type Proof struct {
  /// Power of 2 used for the size of the cuckoo graph
  Cuckoo_sizeshift uint8
  /// The nonces
  Nonces []uint64
}

type Proof interface {
  New(in_nonces []uint64) Proof
  Zero(proof_size usize) Proof
  Random(proof_size usize) Proof
  To_difficulty() target.Difficulty
  Proof_size() usize
}
