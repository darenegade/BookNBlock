import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { Offer } from '../../../data/offer';

@Component({
    selector: 'app-booking-item',
    templateUrl: './booking-item.component.html',
    styleUrls: ['./booking-item.component.scss']
})
export class BookingItemComponent implements OnInit {

    @Input()
    offer: Offer;

    @Input()
    buttonText: string;

    @Output()
    confirmButton: EventEmitter<Offer> = new EventEmitter();

    constructor() { }

    ngOnInit(): void { }

    clickOnButton(offer: Offer) {
        this.confirmButton.emit(offer);
    }
}
