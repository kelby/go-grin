package keychain

const SECRET_KEY_SIZE = 32

// Size of an identifier in bytes
const IDENTIFIER_SIZE = 10

type BlindingFactor [SECRET_KEY_SIZE]uint8

// #[derive(Clone, PartialEq, Eq, Ord, Hash, PartialOrd)]
type Identifier [IDENTIFIER_SIZE]uint8
