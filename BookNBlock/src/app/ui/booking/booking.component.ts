import { Component, OnInit } from '@angular/core';
import { Offer } from '../../data/offer';
import {MockConnector} from '../../connector/mock.connector';

@Component({
  selector: 'app-booking',
  templateUrl: './booking.component.html',
  styleUrls: ['./booking.component.scss']
})
export class BookingComponent implements OnInit {

  allOffers: Offer[] = [];

  constructor(private mockConnecotr: MockConnector) {
  }

  ngOnInit() {
    this.getAllOffers();
  }

  /**
   * Get all offers.
   */
  getAllOffers(): void {
    this.mockConnecotr.getAllOffers().then(result => this.allOffers = result);
  }

}
