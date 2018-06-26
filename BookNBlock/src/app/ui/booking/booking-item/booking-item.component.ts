import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { Offer } from '../../../data/offer';
import { Booking } from '../../../data/booking';
import { QueryService } from '../../../services/query.service';

export class BookingResult {
    offer: Offer;
    booking: Booking;
}

@Component({
    selector: 'app-booking-item',
    templateUrl: './booking-item.component.html',
    styleUrls: ['./booking-item.component.scss']
})
export class BookingItemComponent implements OnInit {

    @Input()
    offer: Offer;

    @Input()
    booking: Booking;

    @Input()
    buttonText: string;

    @Output()
    confirmButton: EventEmitter<BookingResult> = new EventEmitter();

    constructor(private queryService: QueryService) { }

    ngOnInit(): void {
        if (!this.offer) {
            this.queryService.queryOffer(this.booking.offerId).then(result => {
                this.offer = result;
            });
        }
    }

    clickOnButton() {
        const result = new BookingResult();
        result.offer = this.offer;
        result.booking = this.booking;
        this.confirmButton.emit(result);
    }
}
