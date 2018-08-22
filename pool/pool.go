package pool

type Pool struct {
  /// Entries in the pool (tx + info + timer) in simple insertion order.
  Entries []PoolEntry
  /// The blockchain
  Blockchain chain.BlockChain
  Name string
}

func New(chain chain.BlockChain, name string) Pool {
  return Pool {
    entries: [],
    blockchain: chain.Clone(),
    name: name
  }
}
