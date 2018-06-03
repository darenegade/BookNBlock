import { Injectable } from '@angular/core';
import { BlockchainConnector } from './blockchain.connector';
import { Offer } from '../data/offer';
import { OpenDoorMessage } from '../data/OpenDoorMessage';
import { Logger } from '@nsalaun/ng-logger';
const Web3 = require('web3');
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

    this.web3 = new Web3(environment.ethereumAddress);
    this.contract = new this.web3.eth.Contract(abi, address);
    this.contract.options.from = this.user.publicKey;
  }

  async getOffer(id: number): Promise<Offer> {
    this.log.debug(`EthereumConnector.getOffer(${id})`);
    const ids: number[] = await this.contract.methods.getOfferIDs().call();
    this.log.debug(ids);
    if (ids.indexOf(id) >= 0) {
      const o = await this.contract.methods.getOffer(id).call();
      const offer = new Offer();
      offer.id = o.index;
      offer.doorId = o.door;
      offer.prize  = o.prizeInWei;
      offer.fromDate = o.validFrom;
      offer.toDate = o.validUntil;
      offer.address = o.objectAddress;
      offer.name = o.objectName;
      offer.nameLandlord = o.ownerName;
      offer.walletId = o.owner;

      return Promise.resolve(offer);
    } else {
      return Promise.reject(`No offer with id ${id}`);
    }
  }

  async getAllOffers(): Promise<Offer[]> {
    this.log.debug(`EthereumConnector.getAllOffers()`);
    const indeces = await this.contract.methods.getOpenOfferIndeces().call();
    const promises: Promise<Offer>[] =  indeces.map(async i => {
      const p = new Promise<Offer>(async (resolve, reject) => {
        const o = await this.contract.methods.getOffer(i).call();
        const offer = new Offer();
        offer.id = o.index;
        offer.doorId = o.door;
        offer.prize  = o.prizeInWei;
        offer.fromDate = o.validFrom;
        offer.toDate = o.validUntil;
        offer.address = o.objectAddress;
        offer.name = o.objectName;
        offer.nameLandlord = o.ownerName;
        offer.walletId = o.owner;

        resolve(offer);
      });

      promises.push(p);
    });

    return Promise.all(promises);
  }

  searchOffer(criterion: any): Promise<Offer[]> {
    throw new Error('Method not implemented.');
  }

  async insertOffer(offer: Offer): Promise<void> {
    this.log.debug('EthereumConnector.insertOffer()');
    const receipt = await this.contract.methods.insertOffer(offer.prize, offer.name, offer.address, offer.nameLandlord,
      offer.description, offer.doorId, offer.fromDate.getTime(), offer.toDate.getTime()).send();
    this.log.debug(receipt);
  }

  rentOffer(offerId: number, checkIn?: Date, checkOut?: Date): Promise<boolean> {
    return this.contract.methods.rentAnOffer(offerId, checkIn.getTime(), checkOut.getTime()).send({from: this.user.publicKey})
      .then(receipt => Promise.resolve(true))
      .catch(error => Promise.reject(error));
  }

  sendMessage(message: OpenDoorMessage): Promise<void> {
    throw new Error('Method not implemented.');
  }

  authenticateUser(user: any): Promise<boolean> {
    throw new Error('Method not implemented.');
  }
}
