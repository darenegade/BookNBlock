import { Routes } from '@angular/router';
import { LoginComponent } from './ui/login/login.component';
import { OfferComponent } from './ui/offer/offer.component';
import { BookingComponent } from './ui/booking/booking.component';
import {HomeComponent} from "./ui/home/home.component";

export const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'offer', component: OfferComponent },
  { path: 'booking', component: BookingComponent }
];
