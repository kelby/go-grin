package core

type Hash [32]byte

/// A hash consisting of all zeroes, used as a sentinel. No known preimage.
ZERO_HASH := Hash{}
