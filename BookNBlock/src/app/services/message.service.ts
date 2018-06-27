import { Injectable } from '@angular/core';
import { Logger } from '@nsalaun/ng-logger';
import { BlockchainConnectorFactory } from '../connector/connector.factory';
import { OpenDoorMessage } from '../data/OpenDoorMessage';
import { User } from '../data/user';
import { privateEncrypt } from 'crypto-browserify';
import { Buffer } from 'buffer';
import { UserService } from './user.service';

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

  sendMessage(doorId: string, bookingId: number): Promise<void> {
    const message: OpenDoorMessage = new OpenDoorMessage();
    message.doorId = doorId;
    message.booking = bookingId;
    message.renterPubkey = this.user.publicKey;
    // message.timestamp = privateEncrypt(this.user.privateKey, new Buffer(Date.now().toString())).toString();
    message.timestamp = Date.now();
    return this.factory.get().sendMessage(message);
  }

}
