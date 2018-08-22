type Miner struct {
  Config StratumServerConfig
  Chain chain.Chain
  Tx_pool pool.TransactionPool
  Stop AtomicBool

  // Just to hold the port we're on, so this miner can be identified
  // while watching debug output
  Debug_output_id string
}
