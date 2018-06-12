import { TestBed, inject } from '@angular/core/testing';

import { HyperledgerConnector } from './hyperledger.connector';

describe('HyperledgerConnector', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [HyperledgerConnector]
    });
  });

  it('should be created', inject([HyperledgerConnector], (service: HyperledgerConnector) => {
    expect(service).toBeTruthy();
  }));
});
