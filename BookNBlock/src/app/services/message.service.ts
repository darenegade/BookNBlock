import { Injectable } from '@angular/core';
import { Logger } from '@nsalaun/ng-logger';
import { BlockchainConnectorFactory } from '../connector/connector.factory';
import { OpenDoorMessage } from '../data/OpenDoorMessage';
import { User } from '../data/user';
import { privateEncrypt } from 'crypto-browserify';
import { Buffer } from 'buffer';
import {UserService} from './user.service';

@Injectable()
export class MessageService {

  user: User;

  constructor(
    private factory: BlockchainConnectorFactory,
    private userService: UserService,
    private log: Logger
  ) {
    this.user = this.userService.getCurrentLoginUser();
  }

  sendMessage(doorId: number): Promise<void> {
    const message: OpenDoorMessage = {
      doorId: doorId,
      renterId: privateEncrypt(this.user.privateKey, new Buffer(String(this.user.walletId))).toString(),
      renterPK: this.user.publicKey,
      timestemp: privateEncrypt(this.user.privateKey, new Buffer(Date.now().toString())).toString()
    } as OpenDoorMessage;
    return this.factory.get().sendMessage(message);
  }

}
