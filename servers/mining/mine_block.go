/// Serializer that outputs the pre-pow part of the header,
/// including the nonce (last 8 bytes) that can be sent off
/// to the miner to mutate at will
type HeaderPrePowWriter struct {
  Pre_pow []byte
}
