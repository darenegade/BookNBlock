import { Component, OnInit, EventEmitter, Input } from '@angular/core';
import { Offer } from '../../../data/offer';
import { TransactionService } from '../../../services/transaction.service';
import { AlertService } from '../../../services/alert.service';

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

  constructor(private transactionService: TransactionService, private alert: AlertService) { }

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
    this.active = false;
  }

  /**
   * Submit changes to parent component.
   */
  submitChanges() {
    this.transactionService.rentOffer(this.currentOffer.id, new Date(this.fromDate), new Date(this.toDate)).then(() => {
      this.alert.success('Zimmer erfolgreich gebucht.');
      this.closeModal();
    }).catch(err => {
      this.alert.error('Leider konnte das Zimmer nicht gebucht werden.');
      console.error(err);
    });
  }
}
