import { LitElement, html, css } from 'lit';
import { property } from 'lit/decorators.js';
//import { Client } from './client'
import * as iota from "@iota/client-wasm/web";




const logo = new URL('../../assets/open-wc-logo.svg', import.meta.url).href;

export class IotaFrontend extends LitElement {
  @property({ type: String }) title = 'My app';

  @property({type: Number}) count = 0;

  static styles = css`
    :host {
      min-height: 100vh;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: flex-start;
      font-size: calc(10px + 2vmin);
      color: #1a2b42;
      max-width: 960px;
      margin: 0 auto;
      text-align: center;
      background-color: var(--iota-frontend-background-color);
    }

    main {
      flex-grow: 1;
    }

    .logo {
      margin-top: 36px;
      animation: app-logo-spin infinite 20s linear;
    }

    @keyframes app-logo-spin {
      from {
        transform: rotate(0deg);
      }
      to {
        transform: rotate(360deg);
      }
    }

    .app-footer {
      font-size: calc(12px + 0.5vmin);
      align-items: center;
    }

    .app-footer a {
      margin-left: 5px;
    }
  `;


  firstUpdated() {
  //  this.client = new Client("puyBJadUhSx5Ew2KsXmBYA7Vp1SecjSX7bQENXSY7Exg")

   // this.client.increment()
  }

  private _increment(e: Event) {
    this.count++;

    console.log("dfguhjiok")

    iota.init().then(() => {

      async function main() {
        // Get the nodeinfo
        

        let iota_client =  await new iota.ClientBuilder().node("https://iota-node2.alps86.de").build();
        console.log("Nodeinfo: ", await iota_client.getInfo())
      }
      main()
    });
  }

  render() {
    return html`
      <main>
        <div class="logo"><img alt="open-wc logo" src=${logo} /></div>
        <h1>${this.title}</h1>

        <p><button @click="${this._increment}">Click Me!</button></p>
        <p>Click count: ${this.count}</p>

        <p>Edit <code>src/IotaFrontend.ts</code> and save to reload.</p>
        <a
          class="app-link"
          href="https://open-wc.org/guides/developing-components/code-examples"
          target="_blank"
          rel="noopener noreferrer"
        >
          Code examples
        </a>
      </main>

      <p class="app-footer">
        🚽 Made with love by
        <a
          target="_blank"
          rel="noopener noreferrer"
          href="https://github.com/open-wc"
          >open-wc</a
        >.
      </p>
    `;
  }
}
