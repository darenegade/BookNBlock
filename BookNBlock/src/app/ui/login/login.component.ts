import {
  Component,
  OnInit
} from '@angular/core';
import { User } from '../../data/user';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from '../../auth/auth.service';
import { UserService } from '../../services/user.service';
import { AlertService } from '../../services/alert.service';
import { LoginService } from './login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  signIn: boolean;
  loggedIn: boolean;
  returnUrl: string;

  constructor(private route: ActivatedRoute,
    private router: Router,
    private authenticationService: AuthService,
    private userService: UserService,
    private alertService: AlertService,
    private loginService: LoginService) { }

  ngOnInit() {
    this.signIn = true;
    // reset login status
    this.authenticationService.logout();
    this.loginService.currentLoginState.subscribe(state => this.loggedIn = state);
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';
  }

  /**
   * Login.
   * @param user user that requests login
   */
  login(user: User): void {
    this.authenticationService.login(user)
      .subscribe(data => {
        this.router.navigate([this.returnUrl]);
        this.loginService.setIsLoggedIn(true);
      }, error => {
        this.alertService.error(error);
      });
  }

  /**
   * Register a new user.
   * @param user that requests register
   */
  register(user: User): void {
    this.userService.create(user)
      .subscribe(data => {
        this.alertService.success('Registration successful', true);
        this.router.navigate(['/login']);
        this.signIn = true;
      }, error => {
        this.alertService.error(error);
      });
  }

  /**
   * Clear the alerts.
   */
  clearAlerts(): void {
    this.alertService.clear();
  }

}
