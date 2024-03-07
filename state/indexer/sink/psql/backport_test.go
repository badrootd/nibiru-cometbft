package psql

import (
	"github.com/badrootd/nibiru-cometbft/state/indexer"
	"github.com/badrootd/nibiru-cometbft/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
