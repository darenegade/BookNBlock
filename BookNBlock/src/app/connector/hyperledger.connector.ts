import { Injectable } from '@angular/core';
import { BlockchainConnector } from './blockchain.connector';

@Injectable()
export class HyperledgerConnector extends BlockchainConnector {

  constructor() {
    super();
  }

  readBlock(): Promise<any> {
    return Promise.resolve({});
  }

  writeBlock(block: any): Promise<void> {
    return Promise.resolve();
  }

  sendMessage(message: any): Promise<void> {
    return Promise.resolve();
  }
}
