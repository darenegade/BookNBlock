import { Routes } from '@angular/router';
import { LoginComponent } from './ui/login/login.component';
import { OfferComponent } from './ui/offer/offer.component';
import { BookingComponent } from './ui/booking/booking.component';

export const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'offer', component: OfferComponent },
  { path: 'booking', component: BookingComponent }
];
