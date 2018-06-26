import { Component, OnInit, ViewChild } from '@angular/core';
import { Offer } from '../../data/offer';
import * as moment from 'moment';
import { QueryService } from '../../services/query.service';
import { BookingModalComponent } from './booking-modal/booking-modal.component';
import { BookingResult } from './booking-item/booking-item.component';

@Component({
  selector: 'app-booking',
  templateUrl: './booking.component.html',
  styleUrls: ['./booking.component.scss']
})
export class BookingComponent implements OnInit {

  allOffers: Offer[] = [];
  fromDate: string;
  toDate: string;
  locale = 'de';
  selectedDate: any;

  @ViewChild(BookingModalComponent)
  bookModal: BookingModalComponent;


  constructor(private queryService: QueryService) {
    const d = new Date();
    this.fromDate = d.toISOString().substring(0, 10);
    d.setDate(d.getDate() + 10);
    this.toDate = d.toISOString().substring(0, 10);
    this.getAllOffers();
  }

  ngOnInit() {
    this.getAllOffers();
  }

  /**
   * Get all offers.
   */
  getAllOffers(): void {
    const from = new Date(this.fromDate);
    const to = new Date(this.toDate);
    this.queryService.queryAllOffers(from, to).then(result => {
      console.log(result)
      this.allOffers = result
    });
  }

  setSelectedDate(date) {
    this.selectedDate = date;
  }

  canChangeMonthLogic(num, currentDate) {
    currentDate.add(num, 'month');
    const minDate = moment().add(-1, 'month');
    const maxDate = moment().add(1, 'year');

    return currentDate.isBetween(minDate, maxDate);
  }

  isAvailableLogic(dateToCheck: any) {
    if (dateToCheck.isBefore(moment(), 'day')) {
      return false;
    } else {
      return true;
    }
  }

  /**
   * Open the modal dialog to edit the current user.
   */
  openBookModal($event: BookingResult) {
    this.bookModal.isActive($event.offer);
  }

}
