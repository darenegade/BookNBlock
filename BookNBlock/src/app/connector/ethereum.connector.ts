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
import { Booking } from '../data/booking';


/**
 * A connector to the Ethereum Blockchain.
 */
@Injectable()
export class EthereumConnector extends BlockchainConnector {

  private web3: any;
  private user: User;
  private contract: Contract;

  constructor(private log: Logger) {
    super();

  }

  init(user: User) {
    this.user = user;

    const provider = new HDWalletProvider(
      `${this.user.passphrase}`,
      `${environment.ethereumAddress}/${this.user.publicKey}`
    );

    this.web3 = new Web3(provider);
    this.web3.shh.setProvider(environment.nodeAddress);
    this.contract = new this.web3.eth.Contract(abi, address);
    this.contract.options.from = this.user.walletId;

    return this;
  }

  async getOffer(id: number): Promise<Offer> {
    this.log.debug(`EthereumConnector.getOffer(${id})`);
    return this.contract.methods.getOffer(id).call().then(o => {
      const offer = this.mapOffer(o, id);
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
            const offer = this.mapOffer(o, id);
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

    return this.contract.methods.insertOffer(
      offer.prize, offer.title, offer.address, offer.nameLandlord,
      offer.description, offer.doorId, offer.fromDate.getTime(), offer.toDate.getTime())
      .send({gas: '7000000'}).then(receipt => {
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
    return this.web3.shh.addPrivateKey(`0x${this.user.privateKey}`).then(id => {
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
  async getBookingsForUser(): Promise<Booking[]> {
    return this.contract.methods.getOwnBookingIDs().call().then(ids => {
      const bookings: Promise<Booking>[] = [];
      ids.map(i => {
        const id = Number.parseInt(i);
        const promise = new Promise<Booking>((resolve, reject) => {
          this.contract.methods.getBooking(id).call().then(booking => {
            const b = new Booking();
            b.id = id;
            b.offerId = booking.offerID;
            b.checkIn = booking.checkIn;
            b.checkOut = booking.checkOut;
            resolve(b);
          }).catch(error => {
            reject(error);
          });
        });
        bookings.push(promise);
      });
      return Promise.all(bookings);
    }).catch(error => {
      return Promise.reject(error);
    });
  }

  private mapOffer(offerFromBC: any, id: number): Offer {
    const offer = new Offer();
    offer.id = id;
    offer.doorId = offerFromBC.door;
    offer.prize = Number.parseFloat(offerFromBC.priceInWei);
    offer.fromDate = new Date(Number.parseInt(offerFromBC.validFrom));
    offer.toDate = new Date(Number.parseInt(offerFromBC.validUntil));
    offer.address = offerFromBC.objectAddress;
    offer.title = offerFromBC.objectName;
    offer.nameLandlord = offerFromBC.ownerName;
    offer.description = offerFromBC.description;
    return offer;
  }
}
