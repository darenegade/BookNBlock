import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { User } from '../../../data/user';
import { Offer } from '../../../data/offer';
import { BookingResult } from '../../booking/booking-item/booking-item.component';

/**
 * The modal to edit the user information.
 */
@Component({
    selector: 'app-open-door-modal',
    templateUrl: './open-door-modal.component.html',
    styleUrls: ['./open-door-modal.component.scss']
})
export class OpenDoorModalComponent implements OnInit {

    active: boolean;

    @Input()
    bookingResult: BookingResult;

    @Output()
    doorOpen: EventEmitter<BookingResult> = new EventEmitter();

    constructor() { }

    ngOnInit(): void { }

    /**
     * Open the modal.
     */
    isActive() {
        this.active = true;
    }

    /**
     * Close the modal.
     */
    closeModal() {
        this.active = false;
    }

    /**
     * send the offer that has to opened.
     * @param offer the offer for what the door should be opened
     */
    openDoor() {
        this.doorOpen.emit(this.bookingResult);
    }

}
