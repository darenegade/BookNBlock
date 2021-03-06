import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { User } from '../../data/user';
import { UserService } from '../../services/user.service';
import { TransactionService } from '../../services/transaction.service';
import { AlertService } from '../../services/alert.service';
import { Logger } from '@nsalaun/ng-logger';
import { NgxSpinnerService } from 'ngx-spinner';

@Component({
  selector: 'app-offer',
  templateUrl: './offer.component.html',
  styleUrls: ['./offer.component.scss']
})
export class OfferComponent implements OnInit {

  offerForm: FormGroup;
  user: User;

  constructor(private userService: UserService,
    private transactionService: TransactionService,
    private spinner: NgxSpinnerService,
    private alert: AlertService,
    private log: Logger) {
    this.createOfferForm();
  }

  ngOnInit() {
  }

  private createOfferForm() {
    this.offerForm = new FormGroup({
      doorId: new FormControl('', Validators.required),
      nameLandlord: new FormControl('', Validators.required),
      title: new FormControl('', [Validators.required, Validators.minLength(10)]),
      description: new FormControl('', Validators.required),
      date: new FormControl('', Validators.required),
      toDate: new FormControl('', Validators.required),
      prize: new FormControl('', Validators.required),
      street: new FormControl('', Validators.required),
      number: new FormControl('', Validators.required),
      city: new FormControl('', Validators.required),
      zip: new FormControl('', Validators.required),
    });
  }

  onSubmit() {
    this.spinner.show();
    const formModel = this.offerForm.value;
    const address = `${formModel.street} ${formModel.number}
    ${formModel.zip} ${formModel.city}`;
    this.transactionService.insertOffer(formModel.doorId, formModel.prize, new Date(formModel.date),
      new Date(formModel.toDate), address, formModel.title,
      formModel.nameLandlord, formModel.description).then(offer => {
        this.alert.success('Zimmer erfolgreich angelegt.');
        this.spinner.hide();
      }).catch(err => {
        this.spinner.hide();
        this.log.error(err);
        this.alert.error('Zimmer konnte nicht angelegt werden.');
      });
  }
}
