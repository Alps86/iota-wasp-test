import * as iota from "@iota/client-wasm/web";

let config = {
  // Either set a default seed, or set it to undefined to generate a new seed each page load
  seed: undefined,
  waspWebSocketUrl: 'ws://127.0.0.1:9090/chain/%chainId/ws',
  waspApiUrl: 'http://127.0.0.1:9090',
  goshimmerApiUrl: 'http://127.0.0.1:8080',
  chainId: 'pG9BvsC7h1tYPtqQityhH1qjCaN8A65mGWoxDQgSSbRt',
  googleAnalytics: undefined,
  contractName: 'fairroulette'
};

export class Client {
  public chainId: string;
  constructor(chainId: string) {
    this.chainId = chainId;
  }

  public increment() {

    console.log("ooppokpokpokpkpo")

  }


}
