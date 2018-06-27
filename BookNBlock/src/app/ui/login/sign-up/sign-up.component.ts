import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { User } from '../../../data/user';
import { AlertService } from '../../../services/alert.service';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent implements OnInit {

  newUser: User;
  confirmPassword: string;

  @Output()
  registerEvent: EventEmitter<User> = new EventEmitter();

  constructor(private alertSerive: AlertService) {
    this.newUser = new User();
  }

  ngOnInit() {
  }

  register() {
    if (this.newUser.password === this.confirmPassword) {
      this.newUser.userName = this.newUser.email;
      this.registerEvent.emit(this.newUser);
    } else {
      this.alertSerive.warn('Passwords does not match!');
    }
  }

}
