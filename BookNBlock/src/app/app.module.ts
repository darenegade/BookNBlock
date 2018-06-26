import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgLoggerModule, Logger } from '@nsalaun/ng-logger';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
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
import { SignInComponent } from './ui/login/sign-in/sign-in.component';
import { SignUpComponent } from './ui/login/sign-up/sign-up.component';
import { HomeComponent } from './ui/home/home.component';
import { MockConnector } from './connector/mock.connector';
import { AuthService } from './auth/auth.service';
import { AuthGuard } from './auth/auth.guard';
import { UserService } from './services/user.service';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { JwtInterceptor } from './helpers/jwt.interceptor';
import { fakeBackendProvider } from './helpers/fake-backend';
import { AlertService } from './services/alert.service';
import { AlertComponent } from './ui/alert/alert.component';
import { UserComponent } from './ui/user/user.component';
import { LoginService } from './ui/login/login.service';
import { QuestionableBooleanPipe } from './ui/shared/questionableBoolean.pipe';
import { ModalComponent } from './ui/user/modal/modal.component';
import { DatepickerComponent } from './ui/booking/datepicker/datepicker.component';
import { BookingModalComponent } from './ui/booking/booking-modal/booking-modal.component';


const passphrase = `diet asthma equip loan jealous twist divorce cloth gym ramp stomach noise`;
const publicKey = 'VhXic4UDRfv5w86p2hq7';
const privateKey = 'f3813f7438c5cb4ce4ad706b0e7a0196e786e5432c6e59763f50ff3aefa26323';
const walletId = '0xADF900e582b34EC29DF534e32db6250cf9529FB9';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    BookingComponent,
    OfferComponent,
    SignInComponent,
    SignUpComponent,
    HomeComponent,
    AlertComponent,
    UserComponent,
    ModalComponent,
    QuestionableBooleanPipe,
    DatepickerComponent,
    BookingModalComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    NgLoggerModule.forRoot(environment.loglevel),
    RouterModule.forRoot(routes)
  ],
  providers: [
    BlockchainConnectorFactory,
    MockConnector,
    HyperledgerConnector,
    EthereumConnector,
    MessageService,
    QueryService,
    TransactionService,
    AuthService,
    AuthGuard,
    UserService,
    AlertService,
    LoginService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: JwtInterceptor,
      multi: true
    },

    // provider used to create fake backend
    fakeBackendProvider,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
