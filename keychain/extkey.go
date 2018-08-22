package keychain

// #[derive(Debug, Clone)]
type ChildKey struct {
  /// Child number of the key (n derivations)
  N_child uint32
  /// Root key id
  Root_key_id Identifier
  /// Key id
  Key_id Identifier
  /// The private key
  Key SecretKey
}

/// An ExtendedKey is a secret key which can be used to derive new
/// secret keys to blind the commitment of a transaction output.
/// To be usable, a secret key should have an amount assigned to it,
/// but when the key is derived, the amount is not known and must be
/// given.
// #[derive(Debug, Clone)]
type ExtendedKey struct {
  /// Child number of the extended key
  N_child uint32
  /// Root key id
  Root_key_id Identifier
  /// Key id
  Key_id Identifier
  /// The secret key
  Key SecretKey
  /// The chain code for the key derivation chain
  Chain_code [32]uint8
}
