// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
	"test/sc/go/sc"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, sc.ScName, sc.OnLoad)
	require.NoError(t, ctx.ContractExists(sc.ScName))

	for i := 1; i < 10; i++ {
		cb := sc.ScFuncs.CallIncrement(ctx)
		cb.Func.TransferIotas(1).Post()

		require.NoError(t, ctx.Err)
	}

	cb2 := sc.ScFuncs.GetCounter(ctx)
	cb2.Func.Call()

	require.EqualValues(t, 10, cb2.Results.Counter().Value())
}
