// package main

package chain

import "fmt"

import "github.com/kelby/go-grin/core"

/// Options for block validation
// type Options uint32

// const (
// 	/// No flags
//   NONE = Options(0b00000000)
//   /// Runs without checking the Proof of Work, mostly to make testing easier.
//   SKIP_POW = Options(0b00000001)
// /// Adds block while in syncing mode.
//   SYNC = Options(0b00000010)
// 	/// Block validation on a block we mined ourselves
//   MINE = Options(0b00000100)
// )

type Options uint32

const (
	GET    Options = 0
	POST   Options = 1
	DELETE Options = 2
	PUT    Options = 4
)

// Binary
// const (
//     MASK          = 0b11110
//     DEFAULT_COLOR = 0b00000
//     BOLD          = 0b00001
//     UNDERLINE     = 0b00010
//     FLASHING_TEXT = 0b00100
//     NO_CHANGE     = 0b01000
// )

// type SearchRequest uint32
// var (
//   SearchRequestUNIVERSAL SearchRequest = 0b00000001 // UNIVERSAL
//   SearchRequestWEB       SearchRequest = 1 // WEB
//   SearchRequestIMAGES    SearchRequest = 2 // IMAGES
//   SearchRequestLOCAL     SearchRequest = 3 // LOCAL
//   SearchRequestNEWS      SearchRequest = 4 // NEWS
//   SearchRequestPRODUCTS  SearchRequest = 5 // PRODUCTS
//   SearchRequestVIDEO     SearchRequest = 6 // VIDEO
// )

// func main() {
// 	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}

// 	// x := Options.GET
// 	fmt.Println("%s", b)
// }

type TxHashSetRoots struct {
  /// Output root
  Output_Root core.Hash
  /// Range Proof root
  Rproof_Root core.Hash
  /// Kernel root
  Kernel_Root core.Hash
}


/// The tip of a fork. A handle to the fork ancestry from its leaf in the
/// blockchain tree. References the max height and the latest and previous
/// blocks
/// for convenience and the total difficulty.
// #[derive(Serialize, Deserialize, Debug, Clone, PartialEq)]
type Tip struct {
  /// Height of the tip (max height of the fork)
  Height uint64
  /// Last block pushed to the fork
  Last_block_h core.Hash
  /// Block previous to last
  Prev_block_h core.Hash
  /// Total difficulty accumulated on that fork
  Total_difficulty core.Difficulty
}

type Tip interface {
  /// Creates a new tip at height zero and the provided genesis hash.
  New(gbh Hash) Tip

  /// Append a new block to this tip, returning a new updated tip.
  From_block(bh *BlockHeader) Tip
}
