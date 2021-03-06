import { Component, OnInit, EventEmitter, Input, Output } from '@angular/core';
import { Offer } from '../../../data/offer';
import { TransactionService } from '../../../services/transaction.service';
import { AlertService } from '../../../services/alert.service';
import { NgxSpinnerService } from 'ngx-spinner';

@Component({
  selector: 'app-booking-modal',
  templateUrl: './booking-modal.component.html',
  styleUrls: ['./booking-modal.component.scss']
})
export class BookingModalComponent implements OnInit {

  active: boolean;
  currentOffer: Offer;
  fromDate: string;
  toDate: string;

  @Output()
  bookingRequest: EventEmitter<any> = new EventEmitter();

  constructor(
    private spinner: NgxSpinnerService
  ) { }

  ngOnInit(): void {
  }

  /**
   * Open the modal.
   */
  isActive(offer: Offer) {
    this.currentOffer = offer;
    this.fromDate = this.currentOffer.fromDate.toISOString().substring(0, 10);
    this.toDate = this.currentOffer.toDate.toISOString().substring(0, 10);
    this.active = true;
  }

  /**
   * Close the modal.
   */
  closeModal() {
    this.spinner.hide();
    this.active = false;
  }

  /**
   * Submit changes to parent component.
   */
  submitChanges() {
    this.spinner.show();
    this.bookingRequest.emit({ offer: this.currentOffer, fromDate: this.fromDate, toDate: this.toDate });
  }
}
