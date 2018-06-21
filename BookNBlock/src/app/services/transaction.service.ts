import { Injectable } from '@angular/core';
import { Logger } from '@nsalaun/ng-logger';
import { User } from '../data/user';
import { Offer } from '../data/offer';
import { BlockchainConnectorFactory } from '../connector/connector.factory';
import {UserService} from './user.service';

@Injectable()
export class TransactionService {

  user: User;

  constructor(
    private factory: BlockchainConnectorFactory,
    private log: Logger,
    private userService: UserService
  ) {
    this.user = this.userService.getCurrentLoginUser();
  }

  insertOffer(doorId: number, prize: number, fromDate: Date, toDate: Date,
    address: string, name: string, description: string, image?: any): Promise<void> {
    const offer: Offer = {
      id: undefined,
      doorId: doorId,
      isBooked: false,
      prize: prize,
      fromDate: fromDate,
      toDate: toDate,
      address: address,
      name: name,
      description: description,
      walletId: this.user.walletId,
      image: image
    } as Offer;
    return this.factory.get().insertOffer(offer);
  }

  rentOffer(offerId: number, checkIn: Date, checkOut: Date): Promise<boolean> {
    return this.factory.get().rentOffer(offerId, checkIn, checkOut);
  }

}
