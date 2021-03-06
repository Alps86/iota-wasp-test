// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "./wasmlib"
import * as service from "./service"


export interface EventIncrement {
  timestamp: wasmlib.Int32;
  counter: wasmlib.Int64;
}

export interface scEvents {
	sc_increment: (event: EventIncrement) => void;
}

export function handleVmMessage(message: string[]): void {
    const messageHandlers: wasmlib.MessageHandlers = {
		'sc.increment': (index) => {
			const evt: EventIncrement = {
				timestamp: Number(message[++index]),
				counter: BigInt(message[++index]),
			};
			this.emitter.emit('sc_increment', evt);
		},
    };

    const topicIndex = 3;
    const topic = message[topicIndex];

    if (typeof messageHandlers[topic] != 'undefined') {
      messageHandlers[topic](topicIndex);
    }
}
