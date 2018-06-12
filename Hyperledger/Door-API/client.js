var forge = require('node-forge');
var crypto = require('crypto');

const fs = require('fs');
const path = require('path');

var WebSocket= require('rpc-websockets').Client

var ws = new WebSocket('ws://localhost:9999')

// looad the private key from the user store (wallet)
let pkey = fs.readFileSync(path.join(__dirname + '/../hfc-key-store/', 'private-key-priv'), 'utf8');

var privateKey = pkey;
var signer = crypto.createSign('sha256');
signer.update('sesamoeffnedich');
var sign = signer.sign(privateKey,'base64');

process.argv[4]=sign

ws.on('open', function() {
//param 1 = offer id param 2 = türid etc.
  ws.call('openDoor', process.argv.slice(2)).then(function(result) {
  })
  ws.close()
})
