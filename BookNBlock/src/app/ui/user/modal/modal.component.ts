import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { User } from '../../../data/user';

/**
 * The modal to edit the user information.
 */
@Component({
    selector: 'app-modal',
    templateUrl: './modal.component.html',
    styleUrls: ['./modal.component.scss']
})
export class ModalComponent implements OnInit {

    active: boolean;

    @Input()
    currentUser: User;

    @Output()
    close: EventEmitter<User> = new EventEmitter<User>();

    constructor() { }

    ngOnInit(): void {
        console.log(this.currentUser);
    }

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
     * Submit changes to parent component.
     */
    submitChanges() {
        console.log('CurrentUser', this.currentUser);
        this.closeModal();
        this.close.emit(this.currentUser);
    }
}
