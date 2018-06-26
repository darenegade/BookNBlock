import { Component, OnInit, ViewChild } from '@angular/core';
import { User } from '../../data/user';
import { UserService } from '../../services/user.service';
import { ModalComponent } from './modal/modal.component';
import { BlockchainConnectorFactory } from '../../connector/connector.factory';

/**
 * The user management component.
 */
@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.scss']
})
export class UserComponent implements OnInit {

  user: User;

  @ViewChild(ModalComponent)
  editModal: ModalComponent;

  constructor(
    private userService: UserService,
    private blockchainFactory: BlockchainConnectorFactory
  ) { }

  ngOnInit() {
    this.user = this.userService.getCurrentLoginUser();
    console.log(this.user);
  }

  /**
   * Open the modal dialog to edit the current user.
   */
  openEditModal() {
    this.editModal.isActive();
  }

  /**
   * Update the user information.
   */
  updatedUser(updatedUser: User) {
    console.log('UpdatedUser', updatedUser);
    this.userService.update(updatedUser).subscribe();
  }

  private getOffersForUser() {

  }

}
