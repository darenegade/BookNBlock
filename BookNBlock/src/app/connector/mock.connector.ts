import { Injectable } from '@angular/core';
import { BlockchainConnector } from './blockchain.connector';
import { Offer } from '../data/offer';
import { OpenDoorMessage } from '../data/OpenDoorMessage';
import { Logger } from '@nsalaun/ng-logger';

/**
 * A connector for testing.
 */
@Injectable()
export class MockConnector extends BlockchainConnector {

  private offers: Offer[] = [
    { id: 1, doorId: 1, isBooked: false, prize: 100, fromDate: new Date(2018, 0, 1), toDate: new Date(2018, 0, 31),
      address: 'World Disney', name: 'Mickey Mouse', description: 'Mickey\'s house', walletId: 10000, title: 'Wohnen in Mickey\'s house' },
    { id: 2, doorId: 1, isBooked: true, prize: 100, fromDate: new Date(2018, 8, 1), toDate: new Date(2018, 8, 10),
      address: 'World Disney', name: 'Mickey Mouse', description: 'Mickey\'s house', walletId: 10000, title: 'Wohnen in Mickey\'s house' },
    { id: 3, doorId: 1, isBooked: false, prize: 100, fromDate: new Date(2018, 11, 24), toDate: new Date(2018, 11, 30),
      address: 'Entenhausen', name: 'Donald Duck', description: 'Donalds\'s house', walletId: 20000, title: 'Wohnen in Mickey\'s house' }
  ];

  constructor(private log: Logger) {
    super();
  }

  getOffer(id: number): Promise<Offer> {
    this.log.debug(`MockConnector.getOffer(${id})`);
    for (const offer of this.offers) {
      if (offer.id === id) {
        return Promise.resolve(offer);
      }
    }
    return Promise.resolve(undefined);
  }

  getAllOffers(): Promise<Offer[]> {
    this.log.debug(`MockConnector.getAllOffers()`);
    return Promise.resolve(this.offers.filter(offer => !offer.isBooked));
  }

  searchOffer(criterion: any): Promise<Offer[]> {
    this.log.debug(`MockConnector.searchOffer(${JSON.stringify(criterion)})`);
    throw new Error('Method not implemented.');
  }

  insertOffer(offer: Offer): Promise<void> {
    this.log.debug(`MockConnector.insertOffer(${JSON.stringify(offer)})`);
    this.offers.push(offer);
    return Promise.resolve();
  }

  rentOffer(offerId: number, checkIn?: Date, checkOut?: Date): Promise<boolean> {
    this.log.debug(`MockConnector.rentOffer(${offerId})`);
    return Promise.resolve(this.offers.find(offer => offer.id === offerId) !== undefined);
  }

  sendMessage(message: OpenDoorMessage): Promise<void> {
    this.log.debug(`MockConnector.sendMessage(${JSON.stringify(message)})`);
    return Promise.resolve();
  }

  authenticateUser(user: any): Promise<boolean> {
    this.log.debug(`MockConnector.authenticateUser(${JSON.stringify(user)})`);
    return Promise.resolve(true);
  }
}
