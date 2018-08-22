package wallet

type WalletConfig struct {
  // Right now the decision to run or not a wallet is based on the command.
  // This may change in the near-future.
  // pub enable_wallet: bool,

  // The api interface/ip_address that this api server (i.e. this wallet) will run
  // by default this is 127.0.0.1 (and will not accept connections from external clients)
  Api_listen_interface string
  // The port this wallet will run on
  Api_listen_port uint16
  // The api address of a running server node against which transaction inputs
  // will be checked during send
  Check_node_api_http_addr string
  // The directory in which wallet files are stored
  Data_file_dir string
}

type WalletSeed [32]uint8

