package evmstore

import (
	"github.com/Fantom-foundation/lachesis-base/utils/cachescale"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type (
	// StoreCacheConfig is a config for the db.
	StoreCacheConfig struct {
		// Cache size for Receipts (size in bytes).
		ReceiptsSize uint
		// Cache size for Receipts (number of blocks).
		ReceiptsBlocks int
		// Cache size for TxPositions.
		TxPositions int
		// Cache size for EVM database.
		EvmDatabase int
		// Cache size for EvmBlock (number of blocks).
		EvmBlocksNum int
		// Cache size for EvmBlock (size in bytes).
		EvmBlocksSize uint
	}
	// StoreConfig is a config for store db.
	StoreConfig struct {
		Cache StoreCacheConfig
	}
)

// DefaultStoreConfig for product.
func DefaultStoreConfig(scale cachescale.Func) StoreConfig {
	return StoreConfig{
		StoreCacheConfig{
			ReceiptsSize:   scale.U(4 * opt.MiB),
			ReceiptsBlocks: scale.I(4000),
			TxPositions:    scale.I(20000),
			EvmDatabase:    scale.I(16 * opt.MiB),
			EvmBlocksNum:   scale.I(5000),
			EvmBlocksSize:  scale.U(6 * opt.MiB),
		},
	}
}

// LiteStoreConfig is for tests or inmemory.
func LiteStoreConfig() StoreConfig {
	return StoreConfig{
		StoreCacheConfig{
			ReceiptsSize:   3 * 1024,
			ReceiptsBlocks: 100,
			TxPositions:    500,
			EvmBlocksNum:   100,
			EvmBlocksSize:  3 * 1024,
		},
	}
}
