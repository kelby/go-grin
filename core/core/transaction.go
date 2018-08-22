package transaction

import "keychain"
import "secp/pedersen"
import "secp"

/// Options for a kernel's structure or use
// #[derive(Serialize, Deserialize)]
type KernelFeatures uint8

const (
  DEFAULT_KERNEL KernelFeatures = iota
  COINBASE_KERNEL
)

/// Options for block validation
// #[derive(Serialize, Deserialize)]
type OutputFeatures uint8

const {
  /// No flags
  DEFAULT_OUTPUT OutputFeatures = iota
  /// Output is a coinbase output, must not be spent until maturity
  COINBASE_OUTPUT
}

type Transaction struct {
  /// List of inputs spent by the transaction.
  Inputs []Input
  /// List of outputs the transaction produces.
  Outputs []Output
  /// List of kernels that make up this transaction (usually a single kernel).
  Kernels []TxKernel
  /// The kernel "offset" k2
  /// excess is k1G after splitting the key k = k1 + k2
  Offset keychain.BlindingFactor
}

func (self Transaction) Eq(tx Transaction) bool {
  return self.Inputs == tx.Inputs && self.Outputs == tx.Outputs && self.Kernels == tx.Kernels && self.Offset == tx.Offset
}

/// A transaction input.
///
/// Primarily a reference to an output being spent by the transaction.
// #[derive(Serialize, Deserialize, Debug, Clone)]
type Input struct {
  /// The features of the output being spent.
  /// We will check maturity for coinbase output.
  Features OutputFeatures
  /// The commit referencing the output being spent.
  Commit pedersen.Commitment
}

/// An output_identifier can be build from either an input _or_ an output and
/// contains everything we need to uniquely identify an output being spent.
/// Needed because it is not sufficient to pass a commitment around.
// #[derive(Serialize, Deserialize, Debug, Clone, PartialEq)]
type OutputIdentifier struct {
  /// Output features (coinbase vs. regular transaction output)
  /// We need to include this when hashing to ensure coinbase maturity can be
  /// enforced.
  Features OutputFeatures
  /// Output commitment
  Commit pedersen.Commitment
}

/// Output for a transaction, defining the new ownership of coins that are being
/// transferred. The commitment is a blinded value for the output while the
/// range proof guarantees the commitment includes a positive value without
/// overflow and the ownership of the private key. The switch commitment hash
/// provides future-proofing against quantum-based attacks, as well as providing
/// wallet implementations with a way to identify their outputs for wallet
/// reconstruction.
// #[derive(Debug, Copy, Clone, Serialize, Deserialize)]
type Output struct {
  /// Options for an output's structure or use
  Features OutputFeatures
  /// The homomorphic commitment representing the output amount
  Commit pedersen.Commitment
  /// A proof that the commitment is in the right range
  Proof pedersen.RangeProof
}
