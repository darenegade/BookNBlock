import { Injectable } from '@angular/core';

@Injectable()
export abstract class BlockchainConnector {

  abstract readBlock(): Promise<any>;

  abstract writeBlock(block: any): Promise<void>;

  abstract sendMessage(message: any): Promise<void>;

}
