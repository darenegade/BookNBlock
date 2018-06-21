import { Component, OnInit } from '@angular/core';
import { Offer } from '../../data/offer';
import * as moment from 'moment';
import { QueryService } from '../../services/query.service';

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


  constructor(private queryService: QueryService) {
    this.fromDate = new Date().toDateString();
    this.toDate = new Date().toDateString();
  }

  ngOnInit() {
    this.getAllOffers();
  }

  /**
   * Get all offers.
   */
  getAllOffers(): void {
    this.queryService.queryAllOffers().then(result => this.allOffers = result);
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

}
