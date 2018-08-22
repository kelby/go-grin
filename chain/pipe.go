/// Contextual information required to process a new block and either reject or
/// accept it.
type BlockContext struct {
  /// The options
  Opts Options
  /// The store
  Store store.ChainStore
  /// The head
  Head Tip
  /// The POW verification function
  // Pow_verifier fn(&BlockHeader, u8) -> bool
  /// MMR sum tree states
  Txhashset txhashset::TxHashSet
  /// Recently processed blocks to avoid double-processing
  Block_hashes_cache Arc<RwLock<VecDeque<Hash>>>
  /// Recently processed headers to avoid double-processing
  Header_hashes_cache Arc<RwLock<VecDeque<Hash>>>
}

func update_header_head(bh: &BlockHeader, ctx: &mut BlockContext, batch: &mut store::Batch, ) (Tip, error) {
  tip := Tip.From_block(bh)

  if tip.Total_difficulty > ctx.Head.Total_difficulty {
    batch.Save_header_head(&tip)

    ctx.Head = tip

    tip
  } else {
    error.New("None")
  }
}
