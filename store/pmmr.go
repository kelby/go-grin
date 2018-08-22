package store

/// PMMR persistent backend implementation. Relies on multiple facilities to
/// handle writing, reading and pruning.
///
/// * A main storage file appends Hash instances as they come.
/// This AppendOnlyFile is also backed by a mmap for reads.
/// * An in-memory backend buffers the latest batch of writes to ensure the
/// PMMR can always read recent values even if they haven't been flushed to
/// disk yet.
/// * A leaf_set tracks unpruned (unremoved) leaf positions in the MMR..
/// * A prune_list tracks the positions of pruned (and compacted) roots in the
/// MMR.
type  PMMRBackend struct {
  Data_dir string
  Prunable bool
  Hash_file AppendOnlyFile
  Data_file AppendOnlyFile
  Leaf_set LeafSet
  Prune_list PruneList
  marker marker.PhantomData<PMMRable>
}
