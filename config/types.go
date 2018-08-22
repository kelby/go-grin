package config

import "servers"
import "wallet"

/// Going to hold all of the various configuration types
/// separately for now, then put them together as a single
/// ServerConfig object afterwards. This is to flatten
/// out the configuration file into logical sections,
/// as they tend to be quite nested in the code
/// Most structs optional, as they may or may not
/// be needed depending on what's being run
// #[derive(Clone, Debug, Serialize, Deserialize)]
type GlobalConfig struct {
  /// Keep track of the file we've read
  Config_file_path PathBuf
  /// keep track of whether we're using
  /// a config file or just the defaults
  /// for each member
  Using_config_file bool
  /// Global member config
  Members ConfigMembers
}

/// Keeping an 'inner' structure here, as the top
/// level GlobalConfigContainer options might want to keep
/// internal state that we don't necessarily
/// want serialised or deserialised
// #[derive(Clone, Debug, Serialize, Deserialize)]
type ConfigMembers struct {
  /// Server config
  // #[serde(default)]
  Server servers.ServerConfig
  /// Mining config
  Mining_server servers.StratumServerConfig
  /// Logging config
  Logging LoggingConfig

  /// Wallet config. May eventually need to be moved to its own thing. Or not.
  /// Depends on whether we end up starting the wallet in its own process but
  /// with the same lifecycle as the server.
  // #[serde(default)]
  Wallet wallet.WalletConfig
}
