// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package sc

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	IdxParamOwner = 0

	IdxResultOwner = 1

	IdxStateOwner = 2
)

const keyMapLen = 3

var keyMap = [keyMapLen]wasmlib.Key{
	ParamOwner,
	ResultOwner,
	StateOwner,
}

var idxMap [keyMapLen]wasmlib.Key32
