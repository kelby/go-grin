package chain

import "store"

const (
  STORE_SUBPATH string = "chain"
  BLOCK_HEADER_PREFIX byte = byte('h')
  BLOCK_PREFIX byte = byte('b')
  HEAD_PREFIX byte = byte('H')
  HEADER_HEAD_PREFIX byte = byte('I')
  SYNC_HEAD_PREFIX byte = byte('s')
  HEADER_HEIGHT_PREFIX byte = byte('8')
  COMMIT_POS_PREFIX byte = byte('c')
  BLOCK_INPUT_BITMAP_PREFIX byte = byte('B')
)

/// All chain-related database operations
type ChainStore struct {
  Db store.Store
  Header_cache map[Hash]BlockHeader
  Block_input_bitmap_cache map[Hash][]uint8
}

type ChainStore interface {
  Head() (Tip, error)
  Head_header() (BlockHeader, error)
  Get_header_head() (Tip, error)
  Get_sync_head() (Tip, error)
  Get_block(h *Hash) (Block, error)
  Block_exists(h *Hash) (bool, error)
  Get_block_header(h *Hash) (BlockHeader, error)
  Is_on_current_chain(header *BlockHeader) error
  Get_header_by_height(height uint64) (BlockHeader, error)
  Get_output_pos(commit *Commitment) (uint64, error)
  Build_block_input_bitmap(block *Block) (Bitmap, error)
  Build_and_cache_block_input_bitmap(block *Block) (Bitmap, error)
  Get_block_input_bitmap(bh *Hash) ((bool, Bitmap), error)
  Batch() (Batch, error)

  get_block_input_bitmap_db(, bh *Hash) ((bool, Bitmap), error)
}

  func (self *ChainStore) Head() (Tip, error) {
    self.Db.Get_ser(&vec![HEAD_PREFIX]), "HEAD"
  }

func (self *ChainStore) Is_on_current_chain(header *BlockHeader) error {
  head := self.Head()

  if header.Height > head.Height {
    errors.New("header.height > head.height")
  }

  header_at_height := self.Get_header_by_height(header.Height)
  if header.Hash() == header_at_height.Hash() {
    // ...
  } else {
    errors.New("header.hash == header_at_height.hash")
  }
}
