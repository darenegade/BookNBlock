import { BlockchainConnector } from './blockchain.connector';
import { Logger } from '@nsalaun/ng-logger';
import { environment } from '../../environments/environment';
import { MockConnector } from './mock.connector';
import { AppComponent } from '../app.component';
import { EthereumConnector } from './ethereum.connector';
import { HyperledgerConnector } from './hyperledger.connector';
import { User } from '../data/user';
import { Injectable } from '@angular/core';

@Injectable()
export class BlockchainConnectorFactory {

  constructor(private log: Logger, private user: User,
    private mock: MockConnector, private ethereum: EthereumConnector,
    private hyperledger: HyperledgerConnector) {}

  get(): BlockchainConnector {
    if (environment.mock) {
      this.log.debug('Use mock connector');
      return this.mock;
    } else if (this.user.ethereum) {
      this.log.debug('Use ethereum connector');
      return this.ethereum;
    } else {
      this.log.debug('Use hyperledger connector');
      return this.hyperledger;
    }
  }
}
