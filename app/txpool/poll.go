package txpool

import (
	"context"
	"log"
	"time"

	"github.com/itzmeanjan/harmony/app/config"
	"github.com/itzmeanjan/harmony/app/data"
)

// PollTxPoolContent - Poll current content of Ethereum Mempool periodically & do further processing
// processing with data received back
//
// @note Further processing part is yet to be implemented
func PollTxPoolContent(ctx context.Context, rpc *data.Resource, comm chan struct{}) {

	for {

		var result map[string]map[string]map[string]*data.MemPoolTx

		if err := rpc.RPCClient.CallContext(ctx, &result, "txpool_content"); err != nil {

			log.Printf("[!] Failed to fetch mempool content : %s\n", err.Error())

			// Letting supervisor know, pool polling go routine is dying
			// it must take care of spawning another one to continue functioning
			close(comm)
			break

		}

		// Process current tx pool content

		// Sleep for desired amount of time & get to work again
		<-time.After(time.Duration(config.GetMemPoolPollingPeriod()) * time.Millisecond)

	}

}