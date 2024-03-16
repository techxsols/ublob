package indexer

import (
	"context"
	"log"
	"math/big"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (i *Indexer) getBlobs(height int64) ([]ethtypes.Log, error) {
	blockNumber := big.NewInt(height)
	block, err := i.client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatalf("Failed to get block: %v", err)
	}
	txs := block.Transactions()


	for _, tx := range txs {
		if tx.Type() != 0x03 {// blob tx 
			continue
		}
		blob := tx.BlobTxSidecar()
		log.Print(len(blob.Blobs))
	}
	return []ethtypes.Log{},nil
}
