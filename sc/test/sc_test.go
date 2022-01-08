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
}
