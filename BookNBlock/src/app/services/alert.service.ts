import { Injectable } from '@angular/core';
import { Router, NavigationStart } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { Subject } from 'rxjs/Subject';
import { Alert, AlertType } from '../ui/alert/alert';

@Injectable()
export class AlertService {
  private subject = new Subject<Alert>();
  private keepAfterRouteChange = false;

  constructor(private router: Router) {
    // clear alert messages on route change unless 'keepAfterRouteChange' flag is true
    router.events.subscribe(event => {
      if (event instanceof NavigationStart) {
        if (this.keepAfterRouteChange) {
          // only keep for a single route change
          this.keepAfterRouteChange = false;
        } else {
          // clear alert messages
          this.clear();
        }
      }
    });
  }

  getAlert(): Observable<any> {
    return this.subject.asObservable();
  }

  /**
   * Open an SUCESS alter.
   * @param message the massage to display in alert
   * @param keepAfterRouteChange idecator if the alter should be kept after
   *                             the route changes (set to false by default)
   */
  success(message: string, keepAfterRouteChange = false) {
    this.alert(AlertType.Success, message, keepAfterRouteChange);
  }

  /**
  * Open an ERROR alter.
  * @param message the massage to display in alert
  * @param keepAfterRouteChange idecator if the alter should be kept after
  *                             the route changes (set to false by default)
  */
  error(message: string, keepAfterRouteChange = false) {
    this.alert(AlertType.Error, message, keepAfterRouteChange);
  }

  /**
  * Open an INFO alter.
  * @param message the massage to display in alert
  * @param keepAfterRouteChange idecator if the alter should be kept after
  *                             the route changes (set to false by default)
  */
  info(message: string, keepAfterRouteChange = false) {
    this.alert(AlertType.Info, message, keepAfterRouteChange);
  }

  /**
  * Open an WARNING alter.
  * @param message the massage to display in alert
  * @param keepAfterRouteChange idecator if the alter should be kept after
  *                             the route changes (set to false by default)
  */
  warn(message: string, keepAfterRouteChange = false) {
    this.alert(AlertType.Warning, message, keepAfterRouteChange);
  }

  /**
  * Open an generic alter.
  * @param type the alter type (success, error, info, warn)
  * @param message the massage to display in alert
  * @param keepAfterRouteChange idecator if the alter should be kept after
  *                             the route changes (set to false by default)
  */
  alert(type: AlertType, message: string, keepAfterRouteChange = false) {
    this.keepAfterRouteChange = keepAfterRouteChange;
    this.subject.next(<Alert>{ type: type, message: message });
  }

  /**
   * Clear the alter.
   */
  clear() {
    // clear alerts
    this.subject.next();
  }
}
