package core

/// Prunable Merkle Mountain Range implementation. All positions within the tree
/// start at 1 as they're postorder tree traversal positions rather than array
/// indices.
///
/// Heavily relies on navigation operations within a binary tree. In particular,
/// all the implementation needs to keep track of the MMR structure is how far
/// we are in the sequence of nodes making up the MMR.
type struct PMMR {
  /// The last position in the PMMR
  Last_pos uint64
  Backend Backend
  // only needed to parameterise Backend
  // marker marker.PhantomData
}
