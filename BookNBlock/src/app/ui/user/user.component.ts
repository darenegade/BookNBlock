import { Component, OnInit, ViewChild } from '@angular/core';
import { User } from '../../data/user';
import { UserService } from '../../services/user.service';
import { ModalComponent } from './modal/modal.component';
import { BlockchainConnectorFactory } from '../../connector/connector.factory';
import { QueryService } from '../../services/query.service';
import { Offer } from '../../data/offer';
import { OpenDoorModalComponent } from './openDoorModal/open-door-modal.component';
import { MessageService } from '../../services/message.service';

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

  offers: Offer[];

  offerToBook: Offer;

  @ViewChild(OpenDoorModalComponent)
  doorModal: OpenDoorModalComponent;

  @ViewChild(ModalComponent)
  editModal: ModalComponent;

  constructor(
    private userService: UserService,
    private queryService: QueryService,
    private messageService: MessageService,
  ) {

  }

  ngOnInit() {
    this.user = this.userService.getCurrentLoginUser();
    this.getOfferForUser();
  }

  /**
   * Open the modal dialog to edit the current user.
   */
  openEditModal() {
    this.editModal.isActive();
  }

  openDoorModal(offer: Offer) {
    this.offerToBook = offer;
    this.doorModal.isActive();
  }

  /**
   * Update the user information.
   */
  updatedUser(updatedUser: User) {
    this.userService.update(updatedUser).subscribe();
    this.getOfferForUser();
  }

  openDoorForOffer(offer: Offer) {
    this.messageService.sendMessage(Number(offer.doorId));
  }

  private getOfferForUser() {
    this.queryService.queryOffersForUser().then(
      result => this.offers = result);
  }

}
