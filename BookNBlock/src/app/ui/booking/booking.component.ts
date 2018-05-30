import { Component, OnInit } from '@angular/core';
import { Offer } from '../../data/offer';
import {MockConnector} from '../../connector/mock.connector';
import {QueryService} from '../../services/query.service';

@Component({
  selector: 'app-booking',
  templateUrl: './booking.component.html',
  styleUrls: ['./booking.component.scss']
})
export class BookingComponent implements OnInit {

  allOffers: Offer[] = [];

  constructor(private queryService: QueryService) {
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

}
