import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {User} from "../../../data/user";

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent implements OnInit {

  newUser: User;
  confirmPassword: string;

  @Output()
  onRegister: EventEmitter<User> = new EventEmitter();

  constructor() {
    this.newUser = new User();
  }

  ngOnInit() {
  }

  register() {
    if(this.newUser.password === this.confirmPassword) {
      this.newUser.userName = this.newUser.email;
      // TODO: generate WalletId
      this.onRegister.emit(this.newUser);
    }
  }

}
