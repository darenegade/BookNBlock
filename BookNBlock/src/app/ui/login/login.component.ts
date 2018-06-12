import {
  Component,
  OnInit
} from '@angular/core';
import {User} from "../../data/user";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  signIn: boolean;

  constructor() {}

  ngOnInit() {
    this.signIn = true;
  }

  login(user: User) {
    alert('Hallo ' + user.userName );
  }

  register(user: User) {
    alert('register' + user.userName);
  }

}
