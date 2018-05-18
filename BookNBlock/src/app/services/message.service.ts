import { Injectable } from '@angular/core';
import { Logger } from '@nsalaun/ng-logger';
import { BlockchainConnectorFactory } from '../connector/connector.factory';
import { OpenDoorMessage } from '../data/OpenDoorMessage';
import { User } from '../data/user';
import * as crypto from 'crypto';

@Injectable()
export class MessageService {

  constructor(
    private factory: BlockchainConnectorFactory,
    private user: User,
    private log: Logger
  ) { }

  sendMessage(doorId: number): Promise<void> {
    const message: OpenDoorMessage = {
      doorId: doorId,
      renterId: crypto.privateEncrypt(this.user.privateKey, new Buffer(this.user.walletId)).toString(),
      renterPK: this.user.publicKey,
      timestemp: crypto.privateEncrypt(this.user.privateKey, new Buffer(Date.now())).toString()
    } as OpenDoorMessage;
    return this.factory.get().sendMessage(message);
  }

}
