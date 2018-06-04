import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Offer } from '../../data/offer';

@Component({
  selector: 'app-offer',
  templateUrl: './offer.component.html',
  styleUrls: ['./offer.component.scss']
})
export class OfferComponent implements OnInit {

  offerForm: FormGroup;

  constructor() {
    this.createOfferForm();
  }

  ngOnInit() {
  }

  private createOfferForm() {
    this.offerForm = new FormGroup({
      name: new FormControl('', Validators.required),
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
    console.log(JSON.stringify(this.offerForm.value));
  }

  prepareOffer() {
    const formModel = this.offerForm.value;
    const saveOffer: Offer = {
      // TODO: how we handle this id
      id: 333,
      doorId: '3',
      // isBooked: false,
      prize: formModel.prize,
      fromDate: formModel.fromDate,
      toDate: formModel.toDate,
      // TODO: concat address
      address: '',
      name: formModel.name,
      nameLandlord: formModel.name,
      // TODO: get walletId from user
      walletId: 6,
      description: formModel.title,
      title: formModel.title,
      image: ''
    };
  }

}
