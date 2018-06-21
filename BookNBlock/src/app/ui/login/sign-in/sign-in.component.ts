import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { User } from '../../../data/user';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.scss']
})
export class SignInComponent implements OnInit {

  user: User;

  @Output()
  loginEvent: EventEmitter<User> = new EventEmitter();

  constructor() {
    this.user = new User();
  }

  ngOnInit() {
  }

  login() {
    this.loginEvent.emit(this.user);
  }

}
