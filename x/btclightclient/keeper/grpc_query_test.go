package keeper_test

import (
	"math/rand"
	"testing"

	"github.com/babylonlabs-io/babylon/testutil/datagen"
	bbn "github.com/babylonlabs-io/babylon/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	keepertest "github.com/babylonlabs-io/babylon/testutil/keeper"
	"github.com/babylonlabs-io/babylon/x/btclightclient/types"
)

func FuzzHashesQuery(f *testing.F) {
	/*
		 Checks:
		 1. If the request is nil, an error is returned
		 2. If the pagination key has not been set,
			`limit` number of hashes are returned and the pagination key
			 has been set to the next hash.
		 3. If the pagination key has been set,
			the `limit` number of hashes after the key are returned.
		 4. End of pagination: the last hashes are returned properly.
		 5. If the pagination key is not a valid hash, an error is returned.

		 Data Generation:
		 - Generate a random chain of headers and insert into storage.
		 - Generate a random `limit` to the query as an integer between 1 and the
		   total number of hashes.
		   Do checks 2-4 by initially querying without a key and then querying
		   with the nextKey attribute.
	*/
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		blcKeeper, ctx := keepertest.BTCLightClientKeeper(t)

		// Test nil request
		resp, err := blcKeeper.Hashes(ctx, nil)
		if resp != nil {
			t.Errorf("Nil input led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Nil input led to a nil error")
		}

		// Test pagination key being invalid
		// We want the key to have a positive length
		bzSz := datagen.RandomIntOtherThan(r, bbn.BTCHeaderHashLen-1, bbn.BTCHeaderHashLen*10) + 1
		key := datagen.GenRandomByteArray(r, bzSz)
		pagination := constructRequestWithKey(r, key)
		hashesRequest := types.NewQueryHashesRequest(pagination)
		resp, err = blcKeeper.Hashes(ctx, hashesRequest)
		if resp != nil {
			t.Errorf("Invalid key led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Invalid key led to a nil error")
		}

		baseHeader, chain := datagen.GenRandBtcChainInsertingInKeeper(
			t,
			r,
			blcKeeper,
			ctx,
			0,
			uint32(datagen.RandomInt(r, 50))+100,
		)

		// Get the headers map
		headersMap := chain.GetHeadersMap()
		headersMap[baseHeader.Hash.String()] = baseHeader.Header.ToBlockHeader()

		// + 1 is necessary to account for the base header
		totalChainLength := chain.ChainLength() + 1
		chainSize := uint32(totalChainLength)
		limit := uint32(r.Int31n(int32(totalChainLength))) + 1
		// Generate a page request with a limit and a nil key
		pagination = constructRequestWithLimit(r, uint64(limit))
		// Generate the initial query
		hashesRequest = types.NewQueryHashesRequest(pagination)
		// Construct a mapping from the hashes found to a boolean value
		// Will be used later to evaluate whether all the hashes were returned
		hashesFound := make(map[string]bool, 0)

		for headersRetrieved := uint32(0); headersRetrieved < chainSize; headersRetrieved += limit {
			resp, err = blcKeeper.Hashes(ctx, hashesRequest)
			if err != nil {
				t.Errorf("Valid request led to an error %s", err)
			}
			if resp == nil {
				t.Fatalf("Valid request led to a nil response")
			}
			// If we are on the last page the elements retrieved should be equal to the remaining ones
			if headersRetrieved+limit >= chainSize && uint32(len(resp.Hashes)) != chainSize-headersRetrieved {
				t.Fatalf("On the last page expected %d elements but got %d", chainSize-headersRetrieved, len(resp.Hashes))
			}
			// Otherwise, the elements retrieved should be equal to the limit
			if headersRetrieved+limit < chainSize && uint32(len(resp.Hashes)) != limit {
				t.Fatalf("On an intermediate page expected %d elements but got %d", limit, len(resp.Hashes))
			}

			for _, hash := range resp.Hashes {
				// Check if the hash was generated by the chain
				if _, ok := headersMap[hash.String()]; !ok {
					t.Fatalf("Hashes returned a hash that was not created")
				}
				hashesFound[hash.String()] = true
			}

			// Construct the next page request
			pagination = constructRequestWithKeyAndLimit(r, resp.Pagination.NextKey, uint64(limit))
			hashesRequest = types.NewQueryHashesRequest(pagination)
		}

		if len(hashesFound) != len(headersMap) {
			t.Errorf("Some hashes were missed. Got %d while %d were expected", len(hashesFound), len(headersMap))
		}
	})
}

func FuzzContainsQuery(f *testing.F) {
	/*
		Checks:
		1. If the request is nil, (nil, error) is returned
		2. The query returns true or false depending on whether the hash exists.

		Data generation:
		- Generate a random chain of headers and insert into storage.
		- Generate a random header but do not insert it into storage.
	*/
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		blcKeeper, ctx := keepertest.BTCLightClientKeeper(t)

		// Test nil input
		resp, err := blcKeeper.Contains(ctx, nil)
		if resp != nil {
			t.Fatalf("Nil input led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Nil input led to a nil error")
		}

		// Generate a random chain of headers and insert it into storage
		_, chain := datagen.GenRandBtcChainInsertingInKeeper(
			t,
			r,
			blcKeeper,
			ctx,
			0,
			uint32(datagen.RandomInt(r, 50))+100,
		)

		// Test with a non-existent header
		query, _ := types.NewQueryContainsRequest(datagen.GenRandomBTCHeaderInfo(r).Hash.MarshalHex())
		resp, err = blcKeeper.Contains(ctx, query)
		if err != nil {
			t.Errorf("Valid input let to an error: %s", err)
		}
		if resp == nil {
			t.Fatalf("Valid input led to nil response")
		}
		if resp.Contains {
			t.Errorf("Non existent header hash led to true result")
		}

		// Test with an existing header
		query, _ = types.NewQueryContainsRequest(chain.GetRandomHeaderInfo(r).Hash.MarshalHex())
		resp, err = blcKeeper.Contains(ctx, query)
		if err != nil {
			t.Errorf("Valid input let to an error: %s", err)
		}
		if resp == nil {
			t.Fatalf("Valid input led to nil response")
		}
		if !resp.Contains {
			t.Errorf("Existent header hash led to false result")
		}
	})
}

func FuzzMainChainQuery(f *testing.F) {
	/*
			 Checks:
			 1. If the request is nil, an error is returned
			 2. If the pagination key is not a valid hash, an error is returned.
		     3. If the pagination key does not correspond to an existing header, an error is returned.
			 4. If the pagination key is not on the main chain, an error is returned.
			 5. If the pagination key has not been set,
				the first `limit` items of the main chain are returned
			 6. If the pagination key has been set, the `limit` items after it are returned.
			 7. End of pagination: the last elements are returned properly and the next_key is set to nil.

			 Data Generation:
			 - Generate a random chain of headers with different PoW and insert them into the headers storage.
	*/
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		blcKeeper, ctx := keepertest.BTCLightClientKeeper(t)

		// Test nil input
		resp, err := blcKeeper.MainChain(ctx, nil)
		if resp != nil {
			t.Errorf("Nil input led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Nil input led to a nil error")
		}

		// Test pagination key being invalid
		// We want the key to have a positive length
		bzSz := datagen.RandomIntOtherThan(r, bbn.BTCHeaderHashLen-1, bbn.BTCHeaderHashLen*10) + 1
		key := datagen.GenRandomByteArray(r, bzSz)
		pagination := constructRequestWithKey(r, key)
		mainchainRequest := types.NewQueryMainChainRequest(pagination)
		resp, err = blcKeeper.MainChain(ctx, mainchainRequest)
		if resp != nil {
			t.Errorf("Invalid key led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Invalid key led to a nil error")
		}

		// Generate a random chain of headers and insert it into storage
		base, chain := datagen.GenRandBtcChainInsertingInKeeper(
			t,
			r,
			blcKeeper,
			ctx,
			0,
			uint32(datagen.RandomInt(r, 50))+100,
		)

		// Check whether the key being set to an element that does not exist leads to an error
		pagination = constructRequestWithKey(r, datagen.GenRandomBTCHeaderInfo(r).Hash.MustMarshal())
		mainchainRequest = types.NewQueryMainChainRequest(pagination)
		resp, err = blcKeeper.MainChain(ctx, mainchainRequest)
		if resp != nil {
			t.Errorf("Key corresponding to header that does not exist led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Key corresponding to a header that does not exist led to a nil error")
		}

		mainchain := make([]*types.BTCHeaderInfo, 0)

		mainchain = append(mainchain, base)
		mainchain = append(mainchain, chain.GetChainInfo()...)
		// we need to reverse the mainchain because the query returns the mainchain from highest to lowest header
		bbn.Reverse(mainchain)
		// // Check whether the key being set to a non-mainchain element leads to an error
		// // Select a random header
		// header := chain.GetRandomHeaderInfo(r)
		// // Get the tip
		// tip := chah.GetTip()

		// Index into the current element of mainchain that we are iterating
		mcIdx := 0
		// Generate a random limit
		mcSize := uint32(len(mainchain))
		limit := uint32(r.Int31n(int32(len(mainchain))) + 1)

		// 50% of the time, do a reverse request
		// Generate a page request with a limit and a nil key
		pagination = constructRequestWithLimit(r, uint64(limit))
		reverse := false
		if datagen.OneInN(r, 2) {
			reverse = true
			pagination.Reverse = true
		}
		// Generate the initial query
		mainchainRequest = types.NewQueryMainChainRequest(pagination)
		for headersRetrieved := uint32(0); headersRetrieved < mcSize; headersRetrieved += limit {
			resp, err = blcKeeper.MainChain(ctx, mainchainRequest)
			if err != nil {
				t.Errorf("Valid request led to an error %s", err)
			}
			if resp == nil {
				t.Fatalf("Valid request led to nil response")
			}
			// If we are on the last page the elements retrieved should be equal to the remaining ones
			if headersRetrieved+limit >= mcSize && uint32(len(resp.Headers)) != mcSize-headersRetrieved {
				t.Fatalf("On the last page expected %d elements but got %d", mcSize-headersRetrieved, len(resp.Headers))
			}
			// Otherwise, the elements retrieved should be equal to the limit
			if headersRetrieved+limit < mcSize && uint32(len(resp.Headers)) != limit {
				t.Fatalf("On an intermediate page expected %d elements but got %d", limit, len(resp.Headers))
			}

			// Iterate through the headers and ensure that they correspond
			// to the current index into the mainchain.
			for i := 0; i < len(resp.Headers); i++ {
				idx := mcIdx
				if reverse {
					idx = len(mainchain) - mcIdx - 1
				}
				if !resp.Headers[i].Eq(mainchain[idx]) {
					t.Errorf("%t", reverse)
					t.Errorf("Response does not match mainchain. Expected %s got %s", mainchain[idx].Hash, resp.Headers[i].HashHex)
				}
				mcIdx += 1
			}

			// Construct the next page request
			pagination = constructRequestWithKeyAndLimit(r, resp.Pagination.NextKey, uint64(limit))
			if reverse {
				pagination.Reverse = true
			}
			mainchainRequest = types.NewQueryMainChainRequest(pagination)
		}
	})
}

func FuzzTipQuery(f *testing.F) {
	/*
		Checks:
		1. If the request is nil, (nil, error) is returned
		2. The query returns the tip BTC header

		Data generation:
		- Generate a random chain of headers and insert into storage
	*/
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		blcKeeper, ctx := keepertest.BTCLightClientKeeper(t)

		// Test nil input
		resp, err := blcKeeper.Tip(ctx, nil)
		if resp != nil {
			t.Errorf("Nil input led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Nil input led to a nil error")
		}

		// Generate a random chain of headers and insert it into storage
		_, chain := datagen.GenRandBtcChainInsertingInKeeper(
			t,
			r,
			blcKeeper,
			ctx,
			0,
			uint32(datagen.RandomInt(r, 50))+100,
		)

		resp, err = blcKeeper.Tip(ctx, types.NewQueryTipRequest())
		if err != nil {
			t.Errorf("valid input led to an error: %s", err)
		}
		if resp == nil {
			t.Fatalf("Valid input led to nil response")
		}
		if !resp.Header.Eq(chain.GetTipInfo()) {
			t.Errorf("Invalid header returned. Expected %s, got %s", chain.GetTipInfo().Hash, resp.Header.HeaderHex)
		}
	})
}

func FuzzBaseHeaderQuery(f *testing.F) {
	/*
		Checks:
		1. If the request is nil, (nil, error) is returned
		2. The query returns the base BTC header

		Data generation:
		- Generate a random chain of headers and insert into storage.
	*/
	datagen.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		blcKeeper, ctx := keepertest.BTCLightClientKeeper(t)

		// Test nil input
		resp, err := blcKeeper.BaseHeader(ctx, nil)
		if resp != nil {
			t.Errorf("Nil input led to a non-nil response")
		}
		if err == nil {
			t.Errorf("Nil input led to a nil error")
		}

		// Generate a random chain of headers and insert it into storage
		base, _ := datagen.GenRandBtcChainInsertingInKeeper(
			t,
			r,
			blcKeeper,
			ctx,
			0,
			uint32(datagen.RandomInt(r, 50))+100,
		)

		resp, err = blcKeeper.BaseHeader(ctx, types.NewQueryBaseHeaderRequest())
		if err != nil {
			t.Errorf("valid input led to an error: %s", err)
		}
		if resp == nil {
			t.Fatalf("Valid input led to nil response")
		}
		if !resp.Header.Eq(base) {
			t.Errorf("Invalid header returned. Expected %s, got %s", base.Hash, resp.Header.HashHex)
		}
	})
}

// Constructors for PageRequest objects
func constructRequestWithKeyAndLimit(r *rand.Rand, key []byte, limit uint64) *query.PageRequest {
	// If limit is 0, set one randomly
	if limit == 0 {
		limit = uint64(r.Int31n(100) + 1) // Use Int31 instead to avoid overflows
	}
	return &query.PageRequest{
		Key:        key,
		Offset:     0, // only offset or key is set
		Limit:      limit,
		CountTotal: false, // only used when offset is used
		Reverse:    false,
	}
}

func constructRequestWithLimit(r *rand.Rand, limit uint64) *query.PageRequest {
	return constructRequestWithKeyAndLimit(r, nil, limit)
}

func constructRequestWithKey(r *rand.Rand, key []byte) *query.PageRequest {
	return constructRequestWithKeyAndLimit(r, key, 0)
}
