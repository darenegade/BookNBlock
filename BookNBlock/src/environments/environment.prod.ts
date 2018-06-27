import { Level } from '@nsalaun/ng-logger';

export const environment = {
  production: true,
  mock: false,
  loglevel: Level.INFO,
  nodeAddress: 'http://localhost:9945',
  ethereumAddress: 'http://localhost:8545'
};
