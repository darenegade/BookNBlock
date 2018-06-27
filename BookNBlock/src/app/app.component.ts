import { Component, OnInit } from '@angular/core';
import { User } from './data/user';
import { AuthService } from './auth/auth.service';
import { Router } from '@angular/router';
import { LoginService } from './ui/login/login.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

  currentUser: User;
  loggedIn: boolean;

  constructor(
    private authService: AuthService,
    private router: Router,
    private loginService: LoginService) {
    this.updateLoginState();
  }

  ngOnInit() {
    if (this.loginService.isUserLoggedIn()) {
      this.loggedIn = true;
    } else {
      this.updateLoginState();
    }
  }

  /**
   * Logout.
   */
  logout() {
    this.authService.logout();
    this.router.navigate(['']);
  }

  /**
   * Update the current login state.
   */
  private updateLoginState() {
    this.loginService.currentLoginState.subscribe(state => this.loggedIn = state);
  }

}
