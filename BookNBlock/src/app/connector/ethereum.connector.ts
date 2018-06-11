import { Injectable } from '@angular/core';
import { BlockchainConnector } from './blockchain.connector';
import { Offer } from '../data/offer';
import { OpenDoorMessage } from '../data/OpenDoorMessage';
import { Logger } from '@nsalaun/ng-logger';
const Web3 = require('web3');
const HDWalletProvider = require('truffle-hdwallet-provider');
import { abi, address } from './LockContract';
import { environment } from '../../environments/environment';
import { Contract, BatchRequest } from 'web3/types';
import { User } from '../data/user';

/**
 * A connector to the Ethereum Blockchain.
 */
@Injectable()
export class EthereumConnector extends BlockchainConnector {

  private web3: any;
  private contract: Contract;

  constructor(private log: Logger, private user: User) {
    super();

    const provider = new HDWalletProvider(
      `${this.user.privateKey}`,
      `${environment.ethereumAddress}/${this.user.publicKey}`
    );

    this.web3 = new Web3(provider);
    this.contract = new this.web3.eth.Contract(abi, address);
    this.contract.options.from = this.user.walletId;
  }

  async getOffer(id: number): Promise<Offer> {
    this.log.debug(`EthereumConnector.getOffer(${id})`);
    return this.contract.methods.getOfferIDs().call().then(ids => {
      if (ids.map(i => Number.parseInt(i)).indexOf(id) >= 0) {
        return this.contract.methods.getOffer(id).call().then(o => {
          const offer = new Offer();
          offer.id = id;
          offer.doorId = o.door;
          offer.prize  = o.priceInWei;
          offer.fromDate = o.validFrom;
          offer.toDate = o.validUntil;
          offer.address = o.objectAddress;
          offer.name = o.objectName;
          offer.nameLandlord = o.ownerName;

          return Promise.resolve(offer);
        });
      } else {
        return Promise.reject(`No offer with id ${id}`);
      }
    }).catch(error => {
      return Promise.reject(error);
    });
  }

  async getAllOffers(from: Date, to: Date): Promise<Offer[]> {
    this.log.debug(`EthereumConnector.getAllOffers()`);
    return this.contract.methods.getFreeOfferIDs(from.getTime(), to.getTime()).call().then(ids => {
      return ids.map(i => {
        const id = Number.parseInt(i);
        return this.contract.methods.getOffer(id).call().then(o => {
          const offer = new Offer();
          offer.id = id;
          offer.doorId = o.door;
          offer.prize  = o.priceInWei;
          offer.fromDate = o.validFrom;
          offer.toDate = o.validUntil;
          offer.address = o.objectAddress;
          offer.name = o.objectName;
          offer.nameLandlord = o.ownerName;

          return (offer);
        }).catch(error => {
          return Promise.reject(error);
        });
      });
    }).catch(error => {
      return Promise.reject(error);
    });
  }

  searchOffer(criterion: any): Promise<Offer[]> {
    throw new Error('Method not implemented.');
  }

  async insertOffer(offer: Offer): Promise<number> {
    this.log.debug('EthereumConnector.insertOffer()');
    return this.contract.methods.insertOffer(offer.prize, offer.name, offer.address, offer.nameLandlord,
      offer.description, offer.doorId, offer.fromDate.getTime(), offer.toDate.getTime()).send().then(receipt => {
        return 0;
      }).catch(error => {
        return Promise.reject(error);
      });
  }

  rentOffer(offerId: number, checkIn?: Date, checkOut?: Date): Promise<void> {
    this.log.debug(`EthereumConnector.rentOffer()`);
    return this.contract.methods.rentAnOffer(offerId, checkIn.getTime(), checkOut.getTime()).send()
      .then(receipt => Promise.resolve())
      .catch(error => Promise.reject(error));
  }

  sendMessage(message: OpenDoorMessage): Promise<void> {
    throw new Error('Method not implemented.');
  }

  authenticateUser(user: any): Promise<boolean> {
    throw new Error('Method not implemented.');
  }
}
