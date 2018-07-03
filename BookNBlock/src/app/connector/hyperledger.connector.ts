import { Injectable } from '@angular/core';
import { BlockchainConnector } from './blockchain.connector';
import { Offer } from '../data/offer';
import { OpenDoorMessage } from '../data/OpenDoorMessage';
import { User } from '../data/user';
import { Booking } from '../data/booking';
declare function require(name: string);

const Fabric_Client = require('fabric-client');
const path = require('path');
const util = require('util');
const os = require('os');
const fs = require('fs');

const certpath = '../build-files/crypto-config/peerOrganizations/org1.bookandblock.com/peers/peer0.org1.bookandblock.com/msp/tlscacerts/tlsca.org1.bookandblock.com-cert.pem';
const NodeRSA = require('node-rsa');
var mqtt = require('mqtt');
var ip = '127.0.0.1';

var broker = mqtt.connect('mqtt://' + ip);
export class MQTTMessage{
  constructor() {
  }
  doorID: String;
  payload: String;
}




@Injectable()
export class HyperledgerConnector extends BlockchainConnector {
private user: User;

  constructor() {
    super();
  }

  init(user: User) {
    this.user = user;
    return this;
  }




  readBlock(fkt, argsn): Promise<Offer> {
    const fabric_client = new Fabric_Client();
    var channel = fabric_client.newChannel('bookchannel');

    let serverCert = fs.readFileSync(path.join(__dirname, certpath));

//fix or pay at least attention for the ssl-target-name-override
//   "ssl-target-name-override": "peer0.org1.bookandblock.com",
    let peer = fabric_client.newPeer('grpcs://localhost:7051',
      {'request-timeout': '5000',
          'pem': Buffer.from(serverCert).toString(),
          "ssl-target-name-override": "peer0.org1.bookandblock.com",
        });

    channel.addPeer(peer);

    var member_user = null;
    var store_path = path.join(__dirname, 'hfc-key-store');
    console.log('Store path:'+store_path);
    var tx_id = null;

// create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
    Fabric_Client.newDefaultKeyValueStore({ path: store_path}).then((state_store) => {
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
        chaincodeId: 'bookandblockcc', fcn: fkt,
        args: argsn
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
        }
      } else {
        console.log("No payloads were returned from query");
      }
      }).catch((err) => {
        console.error('Failed to query successfully :: ' + err);
      });
    return Promise.resolve({});
  }

  
  
  
  
  invoke(fkt, argsn){
    var fabric_client = new Fabric_Client();

// setup the fabric network
    var channel = fabric_client.newChannel('bookchannel');
// original -> var peer = fabric_client.newPeer('grpc://localhost:7051');

    let serverCert = fs.readFileSync(path.join(__dirname, '../build-files/crypto-config/peerOrganizations/org1.bookandblock.com/peers/peer0.org1.bookandblock.com/msp/tlscacerts/tlsca.org1.bookandblock.com-cert.pem'));

//fix or pay at least attention for the ssl-target-name-override
//   "ssl-target-name-override": "peer0.org1.bookandblock.com",
    let peer = fabric_client.newPeer('grpcs://localhost:7051',
    {'request-timeout': '5000', 'pem': Buffer.from(serverCert).toString(),
        "ssl-target-name-override": "peer0.org1.bookandblock.com",});

    let ordererCert = fs.readFileSync(path.join(__dirname, '../build-files/crypto-config/ordererOrganizations/bookandblock.com/orderers/orderer.bookandblock.com/msp/tlscacerts/tlsca.bookandblock.com-cert.pem'));

    channel.addPeer(peer);

    let orderer = fabric_client.newOrderer('grpcs://localhost:7050',
    {'request-timeout': '5000', 'pem': Buffer.from(ordererCert).toString(),
        "ssl-target-name-override": "orderer.bookandblock.com",});

    channel.addOrderer(orderer);

    var member_user = null;
    var store_path = path.join(__dirname, 'hfc-key-store');
    console.log('Store path:'+store_path);
    var tx_id = null;

// create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
    Fabric_Client.newDefaultKeyValueStore({ path: store_path}).then((state_store) => {
  // assign the store to the fabric client
      fabric_client.setStateStore(state_store);
      var crypto_suite = Fabric_Client.newCryptoSuite();
  // use the same location for the state store (where the users' certificate are kept)
  // and the crypto store (where the users' keys are kept)
      var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
      crypto_suite.setCryptoKeyStore(crypto_store);
      fabric_client.setCryptoSuite(crypto_suite);

  // get the enrolled user from persistence, this user will sign all requests
        // the user should be set by the frontend
      return fabric_client.getUserContext('bookandblockgenericuser', true);
  }).then((user_from_store) => {
      if (user_from_store && user_from_store.isEnrolled()) {
          console.log('Successfully loaded user from persistence');
          member_user = user_from_store;
      } else {
          throw new Error('Failed to get user.... run registerUser.js');
      }

  // get a transaction id object based on the current user assigned to fabric client
      tx_id = fabric_client.newTransactionID();
      console.log("Assigning transaction_id: ", tx_id._transaction_id);

  // right now we use arguments to start a request to the Chaincode, see the Chaincode for available functions and their parameters
        // this should be done by the frontend in the future

      var request = {
                //targets: let default to the peer assigned to the client
          chaincodeId: 'bookandblockcc',
          fcn: fkt,
          args: argsn,
          chainId: 'bookchannel',
          txId: tx_id
        };

  // send the transaction proposal to the peers
        return channel.sendTransactionProposal(request);
    }).then((results) => {
        var proposalResponses = results[0];
        var proposal = results[1];
        let isProposalGood = false;
        if (proposalResponses && proposalResponses[0].response &&
            proposalResponses[0].response.status === 200) {
            isProposalGood = true;
            console.log('Transaction proposal was good');
        } else {
            console.error('Transaction proposal was bad');
        }
        if (isProposalGood) {
            console.log(util.format(
      'Successfully sent Proposal and received ProposalResponse: Status - %s, message - "%s"',
            proposalResponses[0].response.status, proposalResponses[0].response.message));

    // build up the request for the orderer to have the transaction committed
            var request = {
            proposalResponses: proposalResponses,
            proposal: proposal
        };

    // set the transaction listener and set a timeout of 30 sec
    // if the transaction did not get committed within the timeout period,
    // report a TIMEOUT status
        var transaction_id_string = tx_id.getTransactionID(); //Get the transaction ID string to be used by the event processing
        var promises = [];

        var sendPromise = channel.sendTransaction(request);
        promises.push(sendPromise); //we want the send transaction first, so that we know where to check status

    // get an eventhub once the fabric client has a user assigned. The user
    // is required bacause the event registration must be signed
        let event_hub = fabric_client.newEventHub();

                // events are not needed yet, therefore it is safe to use grpc instead of grpcs

        event_hub.setPeerAddr('grpc://localhost:7053');

    // using resolve the promise so that result status may be processed
    // under the then clause rather than having the catch clause process
    // the status
        let txPromise = new Promise((resolve, reject) => {
            let handle = setTimeout(() => {
                event_hub.disconnect();
                resolve({event_status : 'TIMEOUT'}); //we could use reject(new Error('Transaction did not complete within 30 seconds'));
            }, 3000);
            event_hub.connect();
            event_hub.registerTxEvent(transaction_id_string, (tx, code) => {
        // this is the callback for transaction event status
        // first some clean up of event listener
            clearTimeout(handle);
            event_hub.unregisterTxEvent(transaction_id_string);
            event_hub.disconnect();

        // now let the application know what happened
            var return_status = {event_status : code, tx_id : transaction_id_string};
            if (code !== 'VALID') {
                console.error('The transaction was invalid, code = ' + code);
                resolve(return_status); // we could use reject(new Error('Problem with the tranaction, event status ::'+code));
            } else {
                  console.log('The transaction has been committed on peer ' + event_hub._ep._endpoint.addr);
                  resolve(return_status);
            }
            }, (err) => {
        //this is the callback if something goes wrong with the event registration or processing
                reject(new Error('There was a problem with the eventhub ::'+err));
            });
        });
        promises.push(txPromise);

        return Promise.all(promises);
    } else {
        console.error('Failed to send Proposal or receive valid response. Response null or status is not 200. exiting...');
        throw new Error('Failed to send Proposal or receive valid response. Response null or status is not 200. exiting...');
    }
    }).then((results) => {
        console.log('Send transaction promise and event listener promise have completed');
  // check the results in the order the promises were added to the promise all list
    if (results && results[0] && results[0].status === 'SUCCESS') {
        console.log('Successfully sent transaction to the orderer.');
    } else {
        console.error('Failed to order the transaction. Error code: ' + response.status);
    }

    if(results && results[1] && results[1].event_status === 'VALID') {
        console.log('Successfully committed the change to the ledger by the peer');
    } else {
        console.log('Transaction failed to be committed to the ledger due to ::'+results[1].event_status);
    }
    }).catch((err) => {
      console.error('Failed to invoke successfully :: ' + err);
    });
    return Promise.resolve();
}

  
  
  
  
  
  

  
  
  getOffer(id: number): Promise<Offer>{
    const offer = this.readBlock('getOffer', [id]);
    return Promise.resolve(offer);
  }
  
  insertOffer(offer: Offer): Promise<number>{
    const argsn = [offer.id,'true', offer.prize, offer.fromDate, offer.toDate, offer.title, offer.nameLandlord, '', offer.walletId ];
    this.invoke('insertOffer', argsn);
    return Promise.resolve(0);
  }
  
  rentOffer(offerId: number, checkIn?: Date, checkOut?: Date): Promise<void>{
    const argsn = [offerId, this.user.publicKey];
    this.invoke('rentAnOffer', argsn);
    return Promise.resolve();
  }

  
  
  sendMessage(message: OpenDoorMessage): Promise<void> {
    const msg = new MQTTMessage();
    msg.doorID = message.doorId;
    const str = message.timestamp + ', ' + message.renterPubkey
    const key = new NodeRSA(this.user.privateKey);
    msg.payload = key.encryptPrivate(str, 'base64');
    const jsonString= JSON.stringify(msg);
    broker.publish('door', jsonString);
    return Promise.resolve();
    // throw new Error('Method not implemented.');
  }

  
  getAllOffers(from: Date, to: Date): Promise<Offer[]> {
    throw new Error('Method not implemented.');
  }
  authenticateUser(user: any): Promise<boolean> {
    throw new Error('Method not implemented.');
  }


  getBookingsForUser(): Promise<Booking[]> {
    throw new Error('Method not implemented.');
  }

}
