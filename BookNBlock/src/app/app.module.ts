import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgLoggerModule, Logger } from '@nsalaun/ng-logger';
import { FormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { environment } from '../environments/environment';
import { EthereumConnector } from './connector/ethereum.connector';
import { HyperledgerConnector } from './connector/hyperledger.connector';
import { MessageService } from './services/message.service';
import { QueryService } from './services/query.service';
import { TransactionService } from './services/transaction.service';
import { RouterModule } from '@angular/router';
import { routes } from './app.routes';
import { LoginComponent } from './ui/login/login.component';
import { BookingComponent } from './ui/booking/booking.component';
import { OfferComponent } from './ui/offer/offer.component';
import { BlockchainConnectorFactory } from './connector/connector.factory';
import { User } from './data/user';
import { MockConnector } from './connector/mock.connector';

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCqGKukO1De7zhZj6+H0qtjTkVxwTCpvKe4eCZ0FPqri0cb2JZfXJ/DgYSF6vUp
wmJG8wVQZKjeGcjDOL5UlsuusFncCzWBQ7RKNUSesmQRMSGkVb1/3j+skZ6UtW+5u09lHNsj6tQ5
1s1SPrCBkedbNf0Tp0GbMJDyR4e9T04ZZwIDAQABAoGAFijko56+qGyN8M0RVyaRAXz++xTqHBLh
3tx4VgMtrQ+WEgCjhoTwo23KMBAuJGSYnRmoBZM3lMfTKevIkAidPExvYCdm5dYq3XToLkkLv5L2
pIIVOFMDG+KESnAFV7l2c+cnzRMW0+b6f8mR1CJzZuxVLL6Q02fvLi55/mbSYxECQQDeAw6fiIQX
GukBI4eMZZt4nscy2o12KyYner3VpoeE+Np2q+Z3pvAMd/aNzQ/W9WaI+NRfcxUJrmfPwIGm63il
AkEAxCL5HQb2bQr4ByorcMWm/hEP2MZzROV73yF41hPsRC9m66KrheO9HPTJuo3/9s5p+sqGxOlF
L0NDt4SkosjgGwJAFklyR1uZ/wPJjj611cdBcztlPdqoxssQGnh85BzCj/u3WqBpE2vjvyyvyI5k
X6zk7S0ljKtt2jny2+00VsBerQJBAJGC1Mg5Oydo5NwD6BiROrPxGo2bpTbu/fhrT8ebHkTz2epl
U9VQQSQzY1oZMVX8i1m5WUTLPz2yLJIBQVdXqhMCQBGoiuSoSjafUhV7i1cEGpb88h5NBYZzWXGZ
37sJ5QsW+sJyoNde3xH8vdXhzU7eT82D6X/scw9RZz+/6rCJ4p0=
-----END RSA PRIVATE KEY-----`;
const publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqGKukO1De7zhZj6+H0qtjTkVxwTCpvKe4eCZ0
FPqri0cb2JZfXJ/DgYSF6vUpwmJG8wVQZKjeGcjDOL5UlsuusFncCzWBQ7RKNUSesmQRMSGkVb1/
3j+skZ6UtW+5u09lHNsj6tQ51s1SPrCBkedbNf0Tp0GbMJDyR4e9T04ZZwIDAQAB
-----END PUBLIC KEY-----`;

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    BookingComponent,
    OfferComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    NgLoggerModule.forRoot(environment.loglevel),
    RouterModule.forRoot(routes)
  ],
  providers: [
    { provide: User, useValue: {walletId: 12345, privateKey: privateKey, publicKey: publicKey, ethereum: true} as User },
    BlockchainConnectorFactory,
    MockConnector,
    HyperledgerConnector,
    EthereumConnector,
    MessageService,
    QueryService,
    TransactionService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
