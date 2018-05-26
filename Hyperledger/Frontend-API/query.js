//'use strict';
// as it should be modular maybe we should not use strict, this decision should be made by a javascript developer

/*
* Copyright IBM Corp All Rights Reserved
*
* SPDX-License-Identifier: Apache-2.0
*/
/*
 * Chaincode query
 */

var Fabric_Client = require('fabric-client');
var path = require('path');
var util = require('util');
var os = require('os');
var fs = require('fs');

module.exports = {
   
query: function () {

var fabric_client = new Fabric_Client();
var channel = fabric_client.newChannel('bookchannel');

let serverCert = fs.readFileSync(path.join(__dirname, '../build-files/crypto-config/peerOrganizations/org1.bookandblock.com/peers/peer0.org1.bookandblock.com/msp/tlscacerts/tlsca.org1.bookandblock.com-cert.pem'));

//fix or pay at least attention for the ssl-target-name-override
//   "ssl-target-name-override": "peer0.org1.bookandblock.com",
let peer = fabric_client.newPeer(
        'grpcs://localhost:7051',
        {
	'request-timeout': '5000',
          'pem': Buffer.from(serverCert).toString(),
                "ssl-target-name-override": "peer0.org1.bookandblock.com",
        }
);

channel.addPeer(peer);

var member_user = null;
var store_path = path.join(__dirname, 'hfc-key-store');
console.log('Store path:'+store_path);
var tx_id = null;

// create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
Fabric_Client.newDefaultKeyValueStore({ path: store_path
}).then((state_store) => {
	// assign the store to the fabric client
	fabric_client.setStateStore(state_store);
	var crypto_suite = Fabric_Client.newCryptoSuite();
	// use the same location for the state store (where the users' certificate are kept)
	// and the crypto store (where the users' keys are kept)
	var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
	crypto_suite.setCryptoKeyStore(crypto_store);
	fabric_client.setCryptoSuite(crypto_suite);

	// get the enrolled user from persistence, this user will sign all requests
        // this should be done by the frontend
	return fabric_client.getUserContext('bookandblockgenericuser', true);
}).then((user_from_store) => {
	if (user_from_store && user_from_store.isEnrolled()) {
		console.log('Successfully loaded user from persistence');

//mabye fix this ugly array structure
var arr = Object.values(arguments[0]);
console.log(arr[0][0]); 
console.log(arr[0].slice(1))


		member_user = user_from_store;
	} else {
		throw new Error('Failed to get user.... run registerUser.js');
	}

	// right now we use arguments to start a request to the Chaincode, see the Chaincode for available functions and their parameters
        // this should be done by the frontend in the future
	
	        const request = {
                //targets : --- letting this default to the peers assigned to the channel
                chaincodeId: 'bookandblockcc',
                fcn: arr[0][0],
		args: arr[0].slice(1, arr[0].length-1)
        };



	// send the query proposal to the peer
	return channel.queryByChaincode(request);
}).then((query_responses) => {
	console.log("Query has completed, checking results");
	// query_responses could have more than one  results if there multiple peers were used as targets
	if (query_responses && query_responses.length == 1) {
		if (query_responses[0] instanceof Error) {
			console.error("error from query = ", query_responses[0]);
		} else {
			console.log("Response is ", query_responses[0].toString());





/***********************************************************************
************************************************************************
************************************************************************

This code is intended for the Door-API as it validates the sent signed message
As the door reads from the blockchain we placed the function here for testing purposes,
this has to be changed.

*/



var resVal = query_responses[0].toString();
                        name = resVal;
                        
var parsed = JSON.parse(resVal);                        
var valueParsed = parsed['tenantPk']                      
                        
var forge = require('node-forge');
var crypto = require('crypto');

const fs = require('fs');
const path = require('path');

var arr = Object.values(arguments[0]);
                        
var decoded = Buffer.from(valueParsed, 'base64') 
                       
var verifier = crypto.createVerify('sha256');

// some message both parties client and server agreed on this could be extended with a ttl, nonce etc. etc.
verifier.update('sesamoeffnedich');
var ver = verifier.verify(decoded, arr[0].slice(arr[0].length-1)[0],'base64');

// true or false if message was signed correctly
console.log(ver);



/***********************************************************************
************************************************************************
************************************************************************

*/
                    
       
		}
	} else {
		console.log("No payloads were returned from query");
	}
}).catch((err) => {
	console.error('Failed to query successfully :: ' + err);
});

}
      
};
