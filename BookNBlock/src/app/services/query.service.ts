import { Injectable } from '@angular/core';
import { Logger } from '@nsalaun/ng-logger';
import { Offer } from '../data/offer';
import { BlockchainConnectorFactory } from '../connector/connector.factory';
import { Booking } from '../data/booking';

@Injectable()
export class QueryService {

  constructor(
    private factory: BlockchainConnectorFactory,
    private log: Logger
  ) { }

  queryAllOffers(from: Date, to: Date): Promise<Offer[]> {
    return this.factory.get().getAllOffers(from, to);
  }

  queryOffer(offerId: number): Promise<Offer> {
    return this.factory.get().getOffer(offerId);
  }

  queryBookingsForUser(): Promise<Booking[]> {
    return this.factory.get().getBookingsForUser();
  }

}
