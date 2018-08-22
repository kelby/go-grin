package core

type Difficulty struct {
  Num uint64
}

type Difficulty interface {
  /// Difficulty of zero, which is invalid (no target can be
  /// calculated from it) but very useful as a start for additions.
  func Zero() Difficulty

  /// Difficulty of one, which is the minumum difficulty
  /// (when the hash equals the max target)
  func One() Difficulty

  /// Convert a `u32` into a `Difficulty`
  func From_num(num uint64) Difficulty

  /// Computes the difficulty from a hash. Divides the maximum target by the
  /// provided hash and applies the Cuckoo sizeshift adjustment factor (see
  /// https://lists.launchpad.net/mimblewimble/msg00494.html).
  func From_hash_and_shift(h &Hash, shift uint8) Difficulty

  /// Converts the difficulty into a u64
  func (self *Difficulty)To_num() uint64
}
