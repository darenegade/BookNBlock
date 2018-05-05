import { Injectable } from '@angular/core';
import { BlockchainConnector } from '../connector/blockchain.connector';
import { Logger } from '@nsalaun/ng-logger';

@Injectable()
export class TransactionService {

  constructor(
    private connector: BlockchainConnector,
    private log: Logger
  ) { }

}
