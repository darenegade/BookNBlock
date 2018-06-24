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
import { UserService } from '../services/user.service';

/**
 * A connector to the Ethereum Blockchain.
 */
@Injectable()
export class EthereumConnector extends BlockchainConnector {

  private web3: any;
  private user: User;
  private contract: Contract;

  constructor(private log: Logger, private userService: UserService) {
    super();
    this.user = this.userService.getCurrentLoginUser();
    const provider = new HDWalletProvider(
      `${this.user.passphrase}`,
      `${environment.ethereumAddress}/${this.user.publicKey}`
    );

    this.web3 = new Web3(provider);
    this.contract = new this.web3.eth.Contract(abi, address);
    this.contract.options.from = this.user.walletId;
  }

  async getOffer(id: number): Promise<Offer> {
    this.log.debug(`EthereumConnector.getOffer(${id})`);
    return this.contract.methods.getOffer(id).call().then(o => {
      const offer = new Offer();
      offer.id = id;
      offer.doorId = o.door;
      offer.prize = o.priceInWei;
      offer.fromDate = o.validFrom;
      offer.toDate = o.validUntil;
      offer.address = o.objectAddress;
      offer.title = o.objectName;
      offer.nameLandlord = o.ownerName;

      return Promise.resolve(offer);
    }).catch(error => {
      return Promise.reject(error);
    });
  }

  async getAllOffers(from: Date, to: Date): Promise<Offer[]> {
    this.log.debug(`EthereumConnector.getAllOffers()`);
    return this.contract.methods.getFreeOfferIDs(from.getTime(), to.getTime()).call().then(ids => {
      const promises: Promise<Offer>[] = [];
      ids.map(i => {
        const id = Number.parseInt(i);
        const p = new Promise<Offer>((resolve, reject) => {
          this.contract.methods.getOffer(id).call().then(o => {
            const offer = new Offer();
            offer.id = id;
            offer.doorId = o.door;
            offer.prize = o.priceInWei;
            offer.fromDate = o.validFrom;
            offer.toDate = o.validUntil;
            offer.address = o.objectAddress;
            offer.title = o.objectName;
            offer.nameLandlord = o.ownerName;

            resolve(offer);
          }).catch(error => {
            reject(error);
          });
        });
        promises.push(p);
      });
      return Promise.all(promises);
    }).catch(error => {
      return Promise.reject(error);
    });
  }

  searchOffer(criterion: any): Promise<Offer[]> {
    throw new Error('Method not implemented.');
  }

  async insertOffer(offer: Offer): Promise<number> {
    this.log.debug('EthereumConnector.insertOffer()');
    return this.contract.methods.insertOffer(offer.prize, offer.title, offer.address, offer.nameLandlord,
      offer.description, offer.doorId, offer.fromDate.getTime(), offer.toDate.getTime()).send().then(receipt => {
        return 0;
      }).catch(error => {
        return Promise.reject(error);
      });
  }

  rentOffer(offerId: number, checkIn?: Date, checkOut?: Date): Promise<void> {
    this.log.debug(`EthereumConnector.rentOffer()`);
    return this.getOffer(offerId).then(offer => {
      return this.contract.methods.rentAnOffer(offerId, checkIn.getTime(), checkOut.getTime()).send({ value: offer.prize })
        .then(receipt => {
          return Promise.resolve();
        })
        .catch(error => {
          return Promise.reject(error);
        });
    });
  }

  sendMessage(message: OpenDoorMessage): Promise<void> {
    return this.web3.shh.addPrivateKey(this.user.privateKey).then(id => {
      return this.web3.shh.post({
        sig: id, // signs using the private key ID
        pubKey: message.doorId,
        ttl: 10,
        topic: '0x426f6f6b',
        payload: this.web3.utils.asciiToHex(JSON.stringify(message)),
        powTime: 3,
        powTarget: 0.5
      }).then(h => {
        console.log(`Message with hash ${h} was successfuly sent`);
        return Promise.resolve();
      }).catch(error => {
        return Promise.reject(error);
      });
    });
  }

  authenticateUser(user: any): Promise<boolean> {
    throw new Error('Method not implemented.');
  }


  /**
   * Get all offers where the current logged in User has booked.
   * @returns {Promise<Offer[]>} promise that conatins array of all offers
   */
  getBookedOfferForUser(): Promise<Offer[]> {
    return this.contract.methods.getOwnBookingIDs().call().then(offers => {
      return Promise.resolve(offers);
    }).catch(error => {
      return Promise.reject(error);
    });
  }

}
