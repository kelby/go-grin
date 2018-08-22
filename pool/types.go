const (
  /// Dandelion relay timer
  DANDELION_RELAY_SECS uint64 = 600

  /// Dandelion embargo timer
  DANDELION_EMBARGO_SECS uint64 = 180

  /// Dandelion patience timer
  DANDELION_PATIENCE_SECS uint64 = 10

  /// Dandelion stem probability (stem 90% of the time, fluff 10%).
  DANDELION_STEM_PROBABILITY usize = 90
)

/// Configuration for "Dandelion".
/// Note: shared between p2p and pool.
// #[derive(Debug, Clone, Serialize, Deserialize)]
type DandelionConfig struct {
  /// Choose new Dandelion relay peer every n secs.
  // #[serde = "default_dandelion_relay_secs"]
  Relay_secs uint64
  /// Dandelion embargo, fluff and broadcast tx if not seen on network before
  /// embargo expires.
  // #[serde = "default_dandelion_embargo_secs"]
  Embargo_secs uint64
  /// Dandelion patience timer, fluff/stem processing runs every n secs.
  /// Tx aggregation happens on stem txs received within this window.
  // #[serde = "default_dandelion_patience_secs"]
  Patience_secs uint64
  /// Dandelion stem probability (stem 90% of the time, fluff 10% etc.)
  // #[serde = "default_dandelion_stem_probability"]
  Stem_probability usize
}

/// Transaction pool configuration
// #[derive(Clone, Debug, Serialize, Deserialize)]
type PoolConfig struct {
  /// Base fee for a transaction to be accepted by the pool. The transaction
  /// weight is computed from its number of inputs, outputs and kernels and
  /// multiplied by the base fee to compare to the actual transaction fee.
  // #[serde = "default_accept_fee_base"]
  Accept_fee_base uint64

  /// Maximum capacity of the pool in number of transactions
  // #[serde = "default_max_pool_size"]
  Max_pool_size usize
}

/// Represents a single entry in the pool.
/// A single (possibly aggregated) transaction.
// #[derive(Clone, Debug)]
type PoolEntry struct {
  /// The state of the pool entry.
  State PoolEntryState
  /// Info on where this tx originated from.
  Src TxSource
  /// Timestamp of when this tx was originally added to the pool.
  Tx_at Timespec
  /// The transaction itself.
  Tx Transaction
}

/// Placeholder: the data representing where we heard about a tx from.
///
/// Used to make decisions based on transaction acceptance priority from
/// various sources. For example, a node may want to bypass pool size
/// restrictions when accepting a transaction from a local wallet.
///
/// Most likely this will evolve to contain some sort of network identifier,
/// once we get a better sense of what transaction building might look like.
// #[derive(Clone, Debug)]
type TxSource struct {
  /// Human-readable name used for logging and errors.
  Debug_name string
  /// Unique identifier used to distinguish this peer from others.
  Identifier string
}

type PoolError int

/// Possible errors when interacting with the transaction pool.
// #[derive(Debug)]
const (
  /// An invalid pool entry caused by underlying tx validation error
  InvalidTx PoolError = iota
  /// Attempt to add a transaction to the pool with lock_height
  /// greater than height of current block
  ImmatureTransaction
  /// Attempt to spend a coinbase output before it has sufficiently matured.
  ImmatureCoinbase
  /// Problem propagating a stem tx to the next Dandelion relay node.
  DandelionError
  /// Transaction pool is over capacity, can't accept more transactions
  OverCapacity
  /// Transaction fee is too low given its weight
  LowFeeTransaction
  /// Attempt to add a duplicate output to the pool.
  DuplicateCommitment
  /// Other kinds of error (not yet pulled out into meaningful errors).
  Other
)

/// Dummy adapter used as a placeholder for real implementations
// #[allow(dead_code)]
type NoopAdapter struct {}
