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
    { provide: User, useValue: {walletId: 12345, privateKey: 'private', publicKey: 'public', ethereum: true} as User },
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
