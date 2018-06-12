import { TestBed, inject } from '@angular/core/testing';

import { EthereumConnector } from './ethereum.connector';

describe('EthereumConnector', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [EthereumConnector]
    });
  });

  it('should be created', inject([EthereumConnector], (service: EthereumConnector) => {
    expect(service).toBeTruthy();
  }));
});
