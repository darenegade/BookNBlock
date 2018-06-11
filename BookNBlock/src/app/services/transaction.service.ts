import { Injectable } from '@angular/core';
import { Logger } from '@nsalaun/ng-logger';
import { User } from '../data/user';
import { Offer } from '../data/offer';
import { BlockchainConnectorFactory } from '../connector/connector.factory';

@Injectable()
export class TransactionService {

  constructor(
    private factory: BlockchainConnectorFactory,
    private user: User,
    private log: Logger
  ) { }

  insertOffer(doorId: string, prize: number, fromDate: Date, toDate: Date,
    address: string, name: string, nameLandlord: string, description: string, image?: any): Promise<Offer> {
    const offer: Offer = {
      id: undefined,
      doorId: doorId,
      prize: prize,
      fromDate: fromDate,
      toDate: toDate,
      address: address,
      name: name,
      nameLandlord: nameLandlord,
      description: description,
      walletId: this.user.walletId,
      image: image,
      title: ''
    } as Offer;
    return this.factory.get().insertOffer(offer).then(id => {
      offer.id = id;
      return Promise.resolve(offer);
    });
  }

  rentOffer(offerId: number, checkIn: Date, checkOut: Date): Promise<void> {
    return this.factory.get().rentOffer(offerId, checkIn, checkOut);
  }

}
