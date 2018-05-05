import { Injectable } from '@angular/core';
import { Logger } from '@nsalaun/ng-logger';
import { BlockchainConnector } from '../connector/blockchain.connector';

@Injectable()
export class QueryService {

  constructor(
    private connector: BlockchainConnector,
    private log: Logger
  ) { }

}
