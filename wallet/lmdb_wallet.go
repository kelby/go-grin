package wallet

import "store"

type LMDBBackend struct {
  Db store.Store
  Config WalletConfig
  /// passphrase: TODO better ways of dealing with this other than storing
  Passphrase string
  /// Keychain
  Keychain Keychain
  /// client
  Client WalletClient
}
