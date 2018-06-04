import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgLoggerModule, Logger } from '@nsalaun/ng-logger';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
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
import { SignInComponent } from './ui/login/sign-in/sign-in.component';
import { SignUpComponent } from './ui/login/sign-up/sign-up.component';
import { HomeComponent } from './ui/home/home.component';
import { MockConnector } from './connector/mock.connector';

const privateKey = `diet asthma equip loan jealous twist divorce cloth gym ramp stomach noise`;
const publicKey = 'VhXic4UDRfv5w86p2hq7';
const walletId = '0xADF900e582b34EC29DF534e32db6250cf9529FB9';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    BookingComponent,
    OfferComponent,
    SignInComponent,
    SignUpComponent,
    HomeComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    ReactiveFormsModule,
    NgLoggerModule.forRoot(environment.loglevel),
    RouterModule.forRoot(routes)
  ],
  providers: [
    { provide: User, useValue: {walletId: walletId, privateKey: privateKey, publicKey: publicKey, ethereum: true} as User },
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
