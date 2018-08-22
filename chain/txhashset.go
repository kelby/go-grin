package chain

import "store"

type PMMRHandle struct {
	Backend  PMMRBackend
	Last_pos uint64
}

type TxHashSet struct {
	Output_pmmr_h PMMRHandle
	Rproof_pmmr_h PMMRHandle
	Kernel_pmmr_h PMMRHandle

	// chain store used as index of commitments to MMR positions
	Commit_index store.ChainStore
}

type TxHashSet interface {
	/// Open an existing or new set of backends for the TxHashSet
	Open(root_dir string, commit_index store.ChainStore, header BlockHeader) (TxHashSet, error)

	/// Check if an output is unspent.
	/// We look in the index to find the output MMR pos.
	/// Then we check the entry in the output MMR and confirm the hash matches.
	Is_unspent(output_id *OutputIdentifier) (Hash, error)

	/// returns the last N nodes inserted into the tree (i.e. the 'bottom'
	/// nodes at level 0
	/// TODO: These need to return the actual data from the flat-files instead
	/// of hashes now
	Last_n_output(distance uint64) map[Hash]OutputIdentifier

	/// as above, for range proofs
	Last_n_rangeproof(distance uint64) map[Hash]RangeProof

	/// as above, for kernels
	Last_n_kernel(distance uint64) map[Hash]TxKernel

	/// returns outputs from the given insertion (leaf) index up to the
	/// specified limit. Also returns the last index actually populated
	Outputs_by_insertion_index(start_index uint64, max_count uint64) (uint64, []OutputIdentifier)

	/// highest output insertion index available
	Highest_output_insertion_index() uint64

	/// As above, for rangeproofs
	Rangeproofs_by_insertion_index(start_index uint64, max_count uint64) (uint64, []RangeProof)

	/// Get sum tree roots
	/// TODO: Return data instead of hashes
	Roots() (Hash, Hash, Hash)

	/// build a new merkle proof for the given position
	Merkle_proof(commit Commitment) (MerkleProof, string)

	/// Compact the MMR data files and flush the rm logs
	Compact() error
}

func (self *TxHashSet) Is_unspent(output_id *OutputIdentifier) (Hash, error) {
		pos := self.Commit_index.Get_output_pos(&output_id.Commit)

		output_pmmr := PMMR::at(self.Output_pmmr_h.Backend, self.Output_pmmr_h.Last_pos)

		hash, err := output_pmmr.Get_hash(pos)

		if err != nil {
				errors.New(ErrorKind::OutputNotFound.into())
			} else {
				hash, err := output_id.hash_with_index(pos - 1)

				if err != nil {
					errors.New(ErrorKind::TxHashSetErr(format!("txhashset hash mismatch")).Into())
				} else {
					hash
				}
			}
}
