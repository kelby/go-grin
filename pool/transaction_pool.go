package pool

/// Transaction pool implementation.
type TransactionPool struct {
  /// Pool Config
  Config PoolConfig

  /// Our transaction pool.
  Txpool Pool
  /// Our Dandelion "stempool".
  Stempool Pool

  /// The blockchain
  Blockchain BlockChain
  /// The pool adapter
  Adapter PoolAdapter
}

/// Get the total size of the pool.
/// Note: we only consider the txpool here as stempool is under embargo.
func (self *TransactionPool) Total_size() uint {
  return len(self.Txpool)
}

func (self *TransactionPool)add_to_stempool(entry PoolEntry) PoolError {
  // Add tx to stempool (passing in all txs from txpool to validate against).
  self.stempool
    .add_to_pool(entry.clone(), self.txpool.all_transactions())?;

  // Note: we do not notify the adapter here,
  // we let the dandelion monitor handle this.
  Ok(())
}
