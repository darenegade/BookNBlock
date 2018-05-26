import { Component, OnInit } from '@angular/core';
import { Offer } from '../../data/offer';

@Component({
  selector: 'app-booking',
  templateUrl: './booking.component.html',
  styleUrls: ['./booking.component.scss']
})
export class BookingComponent implements OnInit {

  allOffers: Offer[] = [];

  offer: Offer = {
    id: 4,
    doorId: 2222,
    isBooked: false,
    prize: 400,
    fromDate: new Date(2018, 4, 17),
    toDate: new Date(2018, 4, 21),
    address: 'Marienplatz 5, 89069 München',
    name: 'Max Mustermann',
    walletId: 38383,
    description: 'Wunderschöne Wohnung in bester Lage',
    title: 'Wohnen mitten in München',
    image: null
  };

  offer2: Offer = {
    id: 5,
    doorId: 4677,
    isBooked: false,
    prize: 200,
    fromDate: new Date(2018, 7, 1),
    toDate: new Date(2018, 7, 14),
    address: 'Hochkalterstraße 5, 81547 München',
    name: 'Anna Mustermann',
    walletId: 646466,
    description: 'Wohnen im Grünen, gleich in der Nähe der Isar',
    title: 'Vermiete meine Wohnung wegen Urlaub',
    image: null
  };

  offer3: Offer = {
    id: 5,
    doorId: 4677,
    isBooked: false,
    prize: 200,
    fromDate: new Date(2018, 7, 1),
    toDate: new Date(2018, 7, 14),
    address: 'Hochkalterstraße 5, 81547 München',
    name: 'Anna Mustermann',
    walletId: 646466,
    description: 'Wohnen im Grünen, gleich in der Nähe der Isar',
    title: 'Vermiete meine Wohnung wegen Urlaub',
    image: null
  };

  constructor() {
    this.allOffers.push(this.offer);
    this.allOffers.push(this.offer2);
    this.allOffers.push(this.offer3);
  }

  ngOnInit() {
    this.getAllOffers();
  }

  getAllOffers(): Offer[] {
    return this.allOffers;
  }

}
