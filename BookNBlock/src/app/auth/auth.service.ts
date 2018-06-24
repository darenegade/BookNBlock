import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import { User } from '../data/user';


/**
 * Authentication service.
 */
@Injectable()
export class AuthService {

  constructor(private http: HttpClient) { }

  /**
   * Login a user.
   *
   * @param {User} user the user to login
   * @returns {Observable<User>}
   */
  login(user: User): Observable<User> {
    return this.http
      .post<User>('/api/authenticate', user)
      .map(loginUser => {
        if (loginUser && loginUser.token) {
          localStorage.setItem('currentUser', JSON.stringify(loginUser));
        }
        return loginUser;
      }
      );
  }

  /**
   * Logout by removing currentUser from local storage.
   */
  logout() {
    // reomove the user form local storage to log out user
    localStorage.removeItem('currentUser');
  }

}
