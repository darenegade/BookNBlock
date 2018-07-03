import { Component, OnInit, ViewChild } from '@angular/core';
import { User } from '../../data/user';
import { UserService } from '../../services/user.service';
import { ModalComponent } from './modal/modal.component';
import { BlockchainConnectorFactory } from '../../connector/connector.factory';
import { QueryService } from '../../services/query.service';
import { Offer } from '../../data/offer';
import { OpenDoorModalComponent } from './openDoorModal/open-door-modal.component';
import { MessageService } from '../../services/message.service';
import { Booking } from '../../data/booking';
import { BookingResult } from '../booking/booking-item/booking-item.component';
import { AlertService } from '../../services/alert.service';

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

  bookings: Booking[];

  offers: Offer[];

  bookingResult: BookingResult;

  @ViewChild(OpenDoorModalComponent)
  doorModal: OpenDoorModalComponent;

  @ViewChild(ModalComponent)
  editModal: ModalComponent;

  constructor(
    private userService: UserService,
    private queryService: QueryService,
    private messageService: MessageService,
    private alert: AlertService
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

  openDoorModal(bookingResult: BookingResult) {
    this.bookingResult = bookingResult;
    this.doorModal.isActive();
  }

  /**
   * Update the user information.
   */
  updatedUser(updatedUser: User) {
    this.userService.update(updatedUser).subscribe();
    this.getOfferForUser();
  }

  openDoorForOffer($event: BookingResult) {
    this.messageService.sendMessage($event.offer.doorId, $event.booking.id).then(() => {
      this.alert.info('Tür geöffnet');
      this.doorModal.closeModal();
    }).catch(err => {
      this.alert.error(err);
      this.doorModal.closeModal();
    })
  }

  private getOfferForUser() {
    this.queryService.queryBookingsForUser().then(
      result => {
        console.log(result);
        this.bookings = result;
      });
  }

}
