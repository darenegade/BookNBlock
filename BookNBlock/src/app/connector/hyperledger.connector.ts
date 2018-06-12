import { Injectable } from '@angular/core';
import { BlockchainConnector } from './blockchain.connector';
import { Offer } from '../data/offer';
import { OpenDoorMessage } from '../data/OpenDoorMessage';

/**
 * A connector to the Hyperledger Blockchain.
 */
@Injectable()
export class HyperledgerConnector extends BlockchainConnector {

  constructor() {
    super();
  }

  getOffer(id: number): Promise<Offer> {
    throw new Error('Method not implemented.');
  }

  getAllOffers(): Promise<Offer[]> {
    throw new Error('Method not implemented.');
  }

  searchOffer(criterion: any): Promise<Offer[]> {
    throw new Error('Method not implemented.');
  }

  insertOffer(offer: Offer): Promise<void> {
    throw new Error('Method not implemented.');
  }

  rentOffer(offerId: number, checkIn?: Date, checkOut?: Date): Promise<boolean> {
    throw new Error('Method not implemented.');
  }

  sendMessage(message: OpenDoorMessage): Promise<void> {
    throw new Error('Method not implemented.');
  }

  authenticateUser(user: any): Promise<boolean> {
    throw new Error('Method not implemented.');
  }
}
