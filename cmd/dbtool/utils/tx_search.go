package utils

import (
	"fmt"

	"github.com/tendermint/tendermint/rpc/core"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	rpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func DoTxSearch(query string, full bool, pages, perPage int) ([]*coretypes.ResultTx, error) {
	results := []*coretypes.ResultTx{}
	page := 1
	keepGoing := true
	for keepGoing {
		fmt.Printf("Getting transactions (page %d, perPage %d)...\n", page, perPage)
		res, err := core.TxSearch(
			&rpctypes.Context{},
			query,
			false,
			&page,
			&perPage,
			"asc",
		)
		if err != nil {
			return nil, err
		}
		results = append(results, res.Txs...)
		fmt.Printf("results: %d | total: %d\n", len(results), res.TotalCount)
		if len(res.Txs) < perPage ||
			(page == pages && !full) {
			keepGoing = false
		} else {
			page++
		}
	}
	return results, nil
}