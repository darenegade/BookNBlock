import { Component, OnInit } from '@angular/core';
import {AlertService} from '../../services/alert.service';
import {Alert, AlertType} from './alert';

@Component({
  selector: 'app-alert',
  templateUrl: './alert.component.html',
  styleUrls: ['./alert.component.scss']
})
export class AlertComponent implements OnInit {

  alerts: Alert[] = [];

  constructor(private alertService: AlertService) { }

  ngOnInit() {
    this.alertService.getAlert().subscribe((alert: Alert) => {
      if (!alert) {
        // clear alerts when an empty alert is received
        this.alerts = [];
        return;
      }

      // add alert to array
      this.alerts.push(alert);
    });
  }

  removeAlert(alert: Alert) {
    this.alerts = this.alerts.filter(x => x !== alert);
  }

  cssClass(alert: Alert) {
    if (!alert) {
      return;
    }

    // return css class based on alert type
    switch (alert.type) {
      case AlertType.Success:
        return 'notification is-success';
      case AlertType.Error:
        return 'notification is-danger';
      case AlertType.Info:
        return 'notification is-info';
      case AlertType.Warning:
        return 'notification is-warning';
    }
  }
}
