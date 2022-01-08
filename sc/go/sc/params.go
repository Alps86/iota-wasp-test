// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package sc

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ImmutableInitParams struct {
	id int32
}

func (s ImmutableInitParams) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamOwner])
}

type MutableInitParams struct {
	id int32
}

func (s MutableInitParams) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamOwner])
}

type ImmutableSetOwnerParams struct {
	id int32
}

func (s ImmutableSetOwnerParams) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamOwner])
}

type MutableSetOwnerParams struct {
	id int32
}

func (s MutableSetOwnerParams) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamOwner])
}
