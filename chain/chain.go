package chain

import "core"

// #[derive(Debug, Clone)]
type Orphan struct {
  Block Block
  Opts Options
  Added Instant
}

type OrphanBlockPool struct {
  // blocks indexed by their hash
  Orphans map[Hash]Orphan
  // additional index of height -> hash
  // so we can efficiently identify a child block (ex-orphan) after processing a block
  Height_Idx map[u64][]Hash
}

/// Facade to the blockchain block processing pipeline and storage. Provides
/// the current view of the TxHashSet according to the chain state. Also
/// maintains locking for the pipeline to avoid conflicting processing.
type Chain struct {
  Db_Root string
  Store store.ChainStore
  Adapter ChainAdapter

  Head Tip
  Orphans OrphanBlockPool
  Txhashset txhashset.TxHashSet
  // Recently processed blocks to avoid double-processing
  Block_hashes_cache Hash
  // Recently processed headers to avoid double-processing
  Header_hashes_cache Hash

  // POW verification function
  // pow_verifier fn(&BlockHeader, u8) -> bool,
}


type Chain interface {
  /// Initializes the blockchain and returns a new Chain instance. Does a
  /// check on the current chain head to make sure it exists and creates one
  /// based on the genesis block if necessary.
  Init(db_root String, db_env lmdb.Environment, adapter ChainAdapter, genesis Block) (Chain, error)

  /// Processes a single block, then checks for orphans, processing
  /// those as well if they're found
  Process_block(b Block, opts Options) (Tip, Block), error

  type Result interface {}

  /// Attempt to add a new block to the chain. Returns the new chain tip if it
  /// has been added to the longest chain, None if it's added to an (as of
  /// now) orphan chain.
  pub fn process_block_no_orphans(
    &self,
    b: Block,
    opts: Options,
  ) -> Result<(Option<Tip>, Option<Block>), Error> {
    let head = self.store.head()?;
    let bhash = b.hash();
    let mut ctx = self.ctx_from_head(head, opts)?;

    let res = pipe::process_block(&b, &mut ctx);

    let add_to_hash_cache = || {
      // only add to hash cache below if block is definitively accepted
      // or rejected
      let mut cache = self.block_hashes_cache.write().unwrap();
      cache.push_front(bhash);
      cache.truncate(HASHES_CACHE_SIZE);
    };

    match res {
      Ok(Some(ref tip)) => {
        // block got accepted and extended the head, updating our head
        let chain_head = self.head.clone();
        {
          let mut head = chain_head.lock().unwrap();
          *head = tip.clone();
        }
        add_to_hash_cache();

        // notifying other parts of the system of the update
        self.adapter.block_accepted(&b, opts);

        Ok((Some(tip.clone()), Some(b)))
      }
      Ok(None) => {
        add_to_hash_cache();

        // block got accepted but we did not extend the head
        // so its on a fork (or is the start of a new fork)
        // broadcast the block out so everyone knows about the fork
          // broadcast the block
        self.adapter.block_accepted(&b, opts);

        Ok((None, Some(b)))
      }
      Err(e) => {
        match e.kind() {
          ErrorKind::Orphan => {
            let block_hash = b.hash();
            let orphan = Orphan {
              block: b,
              opts: opts,
              added: Instant::now(),
            };

            // In the case of a fork - it is possible to have multiple blocks
            // that are children of a given block.
            // We do not handle this currently for orphans (future enhancement?).
            // We just assume "last one wins" for now.
            &self.orphans.add(orphan);

            debug!(
              LOGGER,
              "process_block: orphan: {:?}, # orphans {}",
              block_hash,
              self.orphans.len(),
            );
            Err(ErrorKind::Orphan.into())
          }
          ErrorKind::Unfit(ref msg) => {
            debug!(
              LOGGER,
              "Block {} at {} is unfit at this time: {}",
              b.hash(),
              b.header.height,
              msg
            );
            Err(ErrorKind::Unfit(msg.clone()).into())
          }
          _ => {
            info!(
              LOGGER,
              "Rejected block {} at {}: {:?}",
              b.hash(),
              b.header.height,
              e
            );
            add_to_hash_cache();
            Err(ErrorKind::Other(format!("{:?}", e).to_owned()).into())
          }
        }
      }
    }
  }

  /// Process a block header received during "header first" propagation.
  pub fn process_block_header(&self, bh: &BlockHeader, opts: Options) -> Result<(), Error>

  /// Attempt to add a new header to the header chain.
  /// This is only ever used during sync and uses sync_head.
  pub fn sync_block_header(&self, bh: &BlockHeader, opts: Options) -> Result<Option<Tip>, Error>

  fn ctx_from_head<a>(&self, head: Tip, opts: Options) -> Result<pipe::BlockContext, Error>

  /// For the given commitment find the unspent output and return the
  /// associated Return an error if the output does not exist or has been
  /// spent. This querying is done in a way that is consistent with the
  /// current chain state, specifically the current winning (valid, most
  /// work) fork.
  func (self *Chain) Is_unspent(output_ref: &OutputIdentifier) (Hash, error)

  fn next_block_height(&self) -> Result<u64, Error>

  /// Validate a vec of "raw" transactions against the current chain state.
  /// Specifying a "pre_tx" if we need to adjust the state, for example when
  /// validating the txs in the stempool we adjust the state based on the
  /// txpool.
  pub fn validate_raw_txs(
    &self,
    txs: Vec<Transaction>,
    pre_tx: Option<Transaction>,
  ) -> Result<Vec<Transaction>, Error>

  /// Verify we are not attempting to spend a coinbase output
  /// that has not yet sufficiently matured.
  func (self *Chain) Verify_coinbase_maturity(tx: &Transaction) error

  /// Validate the current chain state.
  pub fn validate(&self, skip_rproofs: bool) -> Result<(), Error>

  /// Sets the txhashset roots on a brand new block by applying the block on
  /// the current txhashset state.
  pub fn set_txhashset_roots(&self, b: &mut Block, is_fork: bool) -> Result<(), Error>

  /// Return a pre-built Merkle proof for the given commitment from the store.
  pub fn get_merkle_proof(
    &self,
    output: &OutputIdentifier,
    block_header: &BlockHeader,
  ) -> Result<MerkleProof, Error>

  /// Return a merkle proof valid for the current output pmmr state at the
  /// given pos
  pub fn get_merkle_proof_for_pos(&self, commit: Commitment) -> Result<MerkleProof, String>

  /// Returns current txhashset roots
  pub fn get_txhashset_roots(&self) -> (Hash, Hash, Hash)

  /// Provides a reading view into the current txhashset state as well as
  /// the required indexes for a consumer to rewind to a consistent state
  /// at the provided block hash.
  pub fn txhashset_read(&self, h: Hash) -> Result<(u64, u64, File), Error>

  /// Writes a reading view on a txhashset state that's been provided to us.
  /// If we're willing to accept that new state, the data stream will be
  /// read as a zip file, unzipped and the resulting state files should be
  /// rewound to the provided indexes.
  pub fn txhashset_write<T>(
    &self,
    h: Hash,
    txhashset_data: File,
    status: &T,
  ) -> Result<(), Error>
  where
    T: TxHashsetWriteStatus

  /// Triggers chain compaction, cleaning up some unnecessary historical
  /// information. We introduce a chain depth called horizon, which is
  /// typically in the range of a couple days. Before that horizon, this
  /// method will:
  ///
  /// * compact the MMRs data files and flushing the corresponding remove logs
  /// * delete old records from the k/v store (older blocks, indexes, etc.)
  ///
  /// This operation can be resource intensive and takes some time to execute.
  /// Meanwhile, the chain will not be able to accept new blocks. It should
  /// therefore be called judiciously.
  pub fn compact(&self) -> Result<(), Error>

  /// returns the last n nodes inserted into the output sum tree
  pub fn get_last_n_output(&self, distance: u64) -> Vec<(Hash, OutputIdentifier)>

  /// as above, for rangeproofs
  pub fn get_last_n_rangeproof(&self, distance: u64) -> Vec<(Hash, RangeProof)>

  /// as above, for kernels
  pub fn get_last_n_kernel(&self, distance: u64) -> Vec<(Hash, TxKernel)>

  /// outputs by insertion index
  pub fn unspent_outputs_by_insertion_index(
    &self,
    start_index: u64,
    max: u64,
  ) -> Result<(u64, u64, Vec<Output>), Error>

  /// Total difficulty at the head of the chain
  pub fn total_difficulty(&self) -> Difficulty

  /// Orphans pool size
  pub fn orphans_len(&self) -> usize

  /// Total difficulty at the head of the header chain
  pub fn total_header_difficulty(&self) -> Result<Difficulty, Error>

  /// Reset header_head and sync_head to head of current body chain
  pub fn reset_head(&self) -> Result<(), Error>

  /// Get the tip that's also the head of the chain
  pub fn head(&self) -> Result<Tip, Error>

  /// Block header for the chain head
  pub fn head_header(&self) -> Result<BlockHeader, Error>

  /// Gets a block header by hash
  pub fn get_block(&self, h: &Hash) -> Result<Block, Error>

  /// Gets a block header by hash
  pub fn get_block_header(&self, h: &Hash) -> Result<BlockHeader, Error>

  /// Gets the block header at the provided height
  pub fn get_header_by_height(&self, height: u64) -> Result<BlockHeader, Error>

  /// Verifies the given block header is actually on the current chain.
  /// Checks the header_by_height index to verify the header is where we say
  /// it is
  pub fn is_on_current_chain(&self, header: &BlockHeader) -> Result<(), Error>

  /// Get the tip of the current "sync" header chain.
  /// This may be significantly different to current header chain.
  pub fn get_sync_head(&self) -> Result<Tip, Error>

  /// Get the tip of the header chain.
  pub fn get_header_head(&self) -> Result<Tip, Error>

  /// Builds an iterator on blocks starting from the current chain head and
  /// running backward. Specialized to return information pertaining to block
  /// difficulty calculation (timestamp and previous difficulties).
  pub fn difficulty_iter(&self) -> store::DifficultyIter

  /// Check whether we have a block without reading it
  Block_exists(&self, h: Hash) -> Result<bool, Error>
}

/// Verify that the tx has a lock_height that is less than or equal to
/// the height of the next block.
func (self *Chain) Verify_tx_lock_height(tx *Transaction) error {
  height := self.Next_block_height()

  if tx.Lock_height() <= height {
    return nil
  } else {
    errors.New("ErrorKind TxLockHeight")
  }
}
