import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {User} from "../../../data/user";

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.scss']
})
export class SignInComponent implements OnInit {

  user: User;

  @Output()
  onLogin: EventEmitter<User> = new EventEmitter();

  constructor() {
    this.user = new User();
  }

  ngOnInit() {
  }

  login() {
    alert('hello' + this.user.userName);
    this.onLogin.emit(this.user);
  }

}
