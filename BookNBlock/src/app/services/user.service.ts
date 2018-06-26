import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from '../data/user';
import { Observable } from 'rxjs/Observable';


@Injectable()
export class UserService {
  constructor(private http: HttpClient) { }

  /**
   * Get the current logged in user.
   */
  getCurrentLoginUser(): User {
    const user = localStorage.getItem('currentUser');
    if (user) {
      return JSON.parse(user).user;
    }
    return undefined;
  }

  /**
   * Get all users.
   */
  getAll(): Observable<any> {
    return this.http.get<User[]>('/api/users');
  }

  /**
   * Get an user by ID.
   * @param id the id of the user
   */
  getById(id: number): Observable<any> {
    return this.http.get('/api/users/' + id);
  }

  /**
   * Create a new user.
   * @param user the user to create
   */
  create(user: User): Observable<any> {
    return this.http.post('/api/users', user);
  }

  /**
   * Update an user
   * @param user the updated user
   */
  update(user: User): Observable<any> {
    return this.http.put('/api/users/' + user.id, user);
  }

  /**
   * Delete an user by ID.
   * @param id the ID of the user to delete
   */
  delete(id: number): Observable<any> {
    return this.http.delete('/api/users/' + id);
  }
}
