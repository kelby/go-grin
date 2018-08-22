package keychain

// #[derive(Clone, Debug)]
type ExtKeychain struct {
  Secp Secp256k1
  Extkey extkey.ExtendedKey
  Key_overrides map[Identifier]SecretKey
  Key_derivation_cache map[Identifier]uint32
}
