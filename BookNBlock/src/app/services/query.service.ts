import { Injectable } from '@angular/core';
import { Logger } from '@nsalaun/ng-logger';
import { Offer } from '../data/offer';
import { BlockchainConnectorFactory } from '../connector/connector.factory';

@Injectable()
export class QueryService {

  constructor(
    private factory: BlockchainConnectorFactory,
    private log: Logger
  ) { }

  queryAllOffers(): Promise<Offer[]> {
    return this.factory.get().getAllOffers();
  }

  queryOffer(offerId: number): Promise<Offer> {
    return this.factory.get().getOffer(offerId);
  }

}
