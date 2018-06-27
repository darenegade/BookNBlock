import { LitElement, html } from './node_modules/@polymer/lit-element';
import '@material/mwc-button'
import '@material/mwc-button'
import './web3'

export class WhisperSender extends LitElement {

    static get properties() {
        return {
            web3: Object,
            identity: Object
        }
    }

    static get topic() {
        return
    }

    constructor() {
        super();
        this.web3 = new Web3('http://localhost:8545');

        console.log(this.web3);
        console.log(this.web3.shh)
        this.web3.shh.generateSymKeyFromPassword("asdf")
            .then(id => this.identity = id)
            .catch(err =>  {
                console.log(err)
            });
        /*this.web3.shh.subscribe('messages', {
            symKeyID: this.identity,
            ttl: 20,
            topics: ["0x426f6f6b"],
            minPow: 0.8,
        }, function (error, message, subscription) {
            console.log(error + message);
        })*/
    }

    _render({ identity, messages }) {
        return html`
        <div>identity: ${identity}</div>
        
        <label for="to">to</label>
        <input id="to" type="text">
        
        <label for="message">messsage</label>
        <input id="message" type="text">
        <mwc-button on-click="${() => this.sendMessage()}" id="submit">Send</mvc-button>
    `;
    }

    sendMessage() {
        const message = {
            message: this.shadowRoot.getElementById('message').value,
            to: 'TürID'
        }
        const to = this.shadowRoot.getElementById('to').value;

        console.log(message);

        const web3 = this.web3;
        const shh = this.web3.shh;
        web3.shh.post({
            symKeyID: this.identity, // encrypts using the sym key ID
            ttl: 10,
            topic: "0x426f6f6b",
            payload: web3.utils.asciiToHex(message),
            powTime: 3,
            powTarget: 0.5
        }).then(h => {
            console.log(`Message with hash ${h} was successfuly sent`)
        }).catch(err => {
            throw new Error(err)
        });

    }
}

customElements.define('whisper-sender', WhisperSender);
