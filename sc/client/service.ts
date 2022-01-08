// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "./wasmlib"
import * as events from "./events"

export interface GetOwnerResult {
	owner: wasmlib.AgentID;
}

export class ScService extends wasmlib.Service {

  constructor(client: BasicClient, chainId: string) {
    super(client, chainId, 0x210c291a);
  }

	public async init(owner: wasmlib.AgentID): Promise<void> {
		const args: wasmlib.Argument[] = [
				{ key: 'owner', value: owner, },
		];
    	await this.postRequest(0x1f44d644, args);
	}

	public async setOwner(owner: wasmlib.AgentID): Promise<void> {
		const args: wasmlib.Argument[] = [
				{ key: 'owner', value: owner, },
		];
    	await this.postRequest(0x2a15fe7b, args);
	}

	public async getOwner(): Promise<GetOwnerResult> {
		const args: wasmlib.Argument[] = [
		];
		const response = await this.callView(0x137107a6, args);
        let result: GetOwnerResult = {};

		let owner = response['owner'];
		result.owner = '';
		if (owner) {
			result.owner = owner.toString(owner);
		}

		return result;
	}
}
