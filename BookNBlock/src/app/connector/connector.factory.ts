import {BlockchainConnector} from './blockchain.connector';
import {Logger} from '@nsalaun/ng-logger';
import {environment} from '../../environments/environment';
import {MockConnector} from './mock.connector';
import {AppComponent} from '../app.component';
import {EthereumConnector} from './ethereum.connector';
import {HyperledgerConnector} from './hyperledger.connector';
import {User} from '../data/user';
import {Injectable} from '@angular/core';
import {UserService} from '../services/user.service';

@Injectable()
export class BlockchainConnectorFactory {

  user: User;

  constructor(private log: Logger,
              private userService: UserService,
              private mock: MockConnector,
              private ethereum: EthereumConnector,
              private hyperledger: HyperledgerConnector) {
    this.user = this.userService.getCurrentLoginUser();
  }

  get(): BlockchainConnector {
    if (environment.mock) {
      this.log.debug('Use mock connector');
      return this.mock.init(this.user);
    } else if (this.user.ethereum) {
      this.log.debug('Use ethereum connector');
      return this.ethereum.init(this.user);
    } else {
      this.log.debug('Use hyperledger connector');
      return this.hyperledger.init(this.user);
    }
  }
}
