package core

import "keychain"

/// Implemented by types that hold inputs and outputs (and kernels)
/// containing Pedersen commitments.
/// Handles the collection of the commitments as well as their
/// summing, taking potential explicit overages of fees into account.
type Committed interface {
  /// Gather the kernel excesses and sum them.
  Sum_kernel_excesses(offset *keychain.BlindingFactor) (Commitment, Commitment), error

  /// Gathers commitments and sum them.
  Sum_commitments(overage int64) (Commitment, error)

  /// Vector of input commitments to verify.
  Inputs_committed []Commitment

  /// Vector of output commitments to verify.
  Outputs_committed []Commitment

  /// Vector of kernel excesses to verify.
  Kernels_committed []Commitment

  /// Verify the sum of the kernel excesses equals the
  /// sum of the outputs, taking into account both
  /// the kernel_offset and overage.
  Verify_kernel_sums(overage int64, kernel_offset keychain.BlindingFactor) (Commitment, Commitment), error
}
