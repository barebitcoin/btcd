// Copyright (c) 2014-2020 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func main() {
	ctx := context.Background()

	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:                "localhost:8332",
		User:                "yourrpcuser",
		Pass:                "yourrpcpass",
		DisableConnectOnNew: true,
		HTTPPostMode:        true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:          true, // Bitcoin core does not provide TLS by default
	}
	batchClient, err := rpcclient.NewBatch(ctx, connCfg)
	defer batchClient.Shutdown(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// batch mode requires async requests
	blockCount := batchClient.GetBlockCountAsync(ctx)
	block1 := batchClient.GetBlockHashAsync(ctx, 1)
	batchClient.GetBlockHashAsync(ctx, 2)
	batchClient.GetBlockHashAsync(ctx, 3)
	block4 := batchClient.GetBlockHashAsync(ctx, 4)
	difficulty := batchClient.GetDifficultyAsync(ctx)

	// sends all queued batch requests
	batchClient.Send(ctx)

	fmt.Println(blockCount.Receive())
	fmt.Println(block1.Receive())
	fmt.Println(block4.Receive())
	fmt.Println(difficulty.Receive())
}
