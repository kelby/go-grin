
/// A proof that a transaction sums to zero. Includes both the transaction's
/// Pedersen commitment and the signature, that guarantees that the commitments
/// amount to zero.
/// The signature signs the fee and the lock_height, which are retained for
/// signature validation.
// #[derive(Serialize, Deserialize, Debug, Clone)]
type TxKernel struct {
  /// Options for a kernel's structure or use
  Features KernelFeatures
  /// Fee originally included in the transaction this proof is for.
  Fee uint64
  /// This kernel is not valid earlier than lock_height blocks
  /// The max lock_height of all *inputs* to this transaction
  Lock_height uint64
  /// Remainder of the sum of all transaction commitments. If the transaction
  /// is well formed, amounts components should sum to zero and the excess
  /// is hence a valid public key.
  Excess pedersen.Commitment
  /// The signature proving the excess is a valid public key, which signs
  /// the transaction fee.
  Excess_sig secp.Signature
}

/// Return the excess commitment for this tx_kernel.
func (self *TxKernel) Excess() pedersen.Commitment {
  self.Excess
}

/// Verify the transaction proof validity. Entails handling the commitment
/// as a public key and checking the signature verifies with the fee as
/// message.
func (self *TxKernel) Verify() -> Result<(), secp::Error> {
  let msg = Message::from_slice(&kernel_sig_msg(self.Fee, self.Lock_height))?;
  let secp = static_secp_instance();
  let secp = secp.lock().unwrap();
  let sig = &self.excess_sig;
  // Verify aggsig directly in libsecp
  let pubkey = &self.excess.to_pubkey(&secp)?;
  if !secp::aggsig::verify_single(&secp, &sig, &msg, None, &pubkey, false) {
    return Err(secp::Error::IncorrectSignature);
  }
  Ok(())
}
