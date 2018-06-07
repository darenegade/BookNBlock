# Using whisper

- Start the boot node with the whisperbootnode executeable.
- Start a node for the client with the whispernode executeable.

## Sending OpenDoorMessages
Connect with web3, put your private key on the node and start posting.

```js
const RENTER_PRIVATE_KEY =
  `0xc81803c093599fb7a4bba473ed74a2209fa53fcb59a7a7bcc6e8226157e87cb5`;
const DOOR_PRIVATE_KEY =
  `0x6ecd6756d5e9d9df44be83b82d99b17983ae5ce9d0f2de9dcd68c80197aafc4a`;
const DOOR_PUBLIC_KEY =
` 0x04f0f871df7b11b3a186210ef251d10837ccfb757de9d8669225bcf73632853def72ae7680f8acdfa1ac94345017d2b4c185275a1ea2f7bbe03e939146ba355889`;

this.web3 = new Web3('http://localhost:8545');
this.web3.shh.addPrivateKey(RENTER_PRIVATE_KEY).then(
  id => this.renterIdentity = id
);
...
web3.shh.post({
	 sig: this.renterIdentity, // signs using the private key ID
	 pubKey: DOOR_PUBLIC_KEY,
	 ttl: 10,
	 topic: "0x426f6f6b",
	 payload: web3.utils.asciiToHex(JSON.stringify(message)),
	 powTime: 3,
	 powTarget: 0.5
}).then(h => {
	 console.log(`Message with hash ${h} was successfuly sent`)
}).catch(err => {
	 throw new Error(err)
});
````