import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { UserService } from '../../services/user.service';
import { isNullOrUndefined } from 'util';

@Injectable()
export class LoginService {

  private loggedInSource = new BehaviorSubject(false);
  currentLoginState = this.loggedInSource.asObservable();

  constructor(private userService: UserService) { }

  /**
   * Update the current loggin state for the user.
   * @param value true if the user is logged in otherwise false
   */
  setIsLoggedIn(value: boolean): void {
    this.loggedInSource.next(value);
  }

  /**
   * Check if a user is currently logged in.
   * @returns ture if a user is logged in otherwise false.
   */
  isUserLoggedIn(): boolean {
    return isNullOrUndefined(this.userService.getCurrentLoginUser()) ? false : true;
  }
}
