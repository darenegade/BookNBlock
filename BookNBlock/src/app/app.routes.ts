import { Routes } from '@angular/router';
import { LoginComponent } from './ui/login/login.component';
import { OfferComponent } from './ui/offer/offer.component';
import { BookingComponent } from './ui/booking/booking.component';
import {HomeComponent} from './ui/home/home.component';
import {AuthGuard} from './auth/auth.guard';
import {UserComponent} from './ui/user/user.component';

export const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },

  { path: 'login', component: LoginComponent },
  { path: 'userManagement', component: UserComponent, canActivate: [AuthGuard]},
  { path: 'home', component: HomeComponent, canActivate: [AuthGuard]},
  { path: 'offer', component: OfferComponent, canActivate: [AuthGuard] },
  { path: 'booking', component: BookingComponent, canActivate: [AuthGuard] },

  // otherwise redirect to home
  { path: '**', redirectTo: '' }
];
