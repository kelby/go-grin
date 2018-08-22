package wallet

type FileBatch struct {
  /// List of outputs
  Outputs map[string]OutputData
  /// Wallet Details
  Details WalletDetails
  /// Data file path
  Data_file_path string
  /// Details file path
  Details_file_path string
  /// lock file path
  Lock_file_path string
}

/// Wallet information tracking all our outputs. Based on HD derivation and
/// avoids storing any key data, only storing output amounts and child index.
// #[derive(Debug, Clone)]
type FileWallet struct {
  /// Keychain
  Keychain Keychain
  /// Client implementation
  Client WalletClient
  /// Configuration
  Config WalletConfig
  /// passphrase: TODO better ways of dealing with this other than storing
  Passphrase: string
  /// List of outputs
  Outputs map[string]OutputData
  /// Details
  Details WalletDetails
  /// Data file path
  Data_file_path string
  /// Backup file path
  Backup_file_path string
  /// lock file path
  Lock_file_path string
  /// details file path
  Details_file_path string
  /// Details backup file path
  Details_bak_path string
}
