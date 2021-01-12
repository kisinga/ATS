import { Component, OnDestroy, OnInit } from "@angular/core";
import { APIKey } from "app/models/api-key.model";
import { ApiKeyService } from "app/services/api-key.service";
import { StateService } from "app/services/state.service";
import { ReplaySubject } from "rxjs";
import { takeUntil } from "rxjs/operators";

@Component({
  selector: "ngx-dashboard",
  templateUrl: "./dashboard.component.html",
})
export class DashboardComponent implements OnInit, OnDestroy {
  apiKey: APIKey;
  comopnentDestroyed: ReplaySubject<boolean> = new ReplaySubject<boolean>();
  loading: true;

  constructor(
    private state: StateService,
    private apikeyService: ApiKeyService
  ) {
    this.state.dashboardApiKey
      .pipe(takeUntil(this.comopnentDestroyed))
      .subscribe((k) => {
        this.apiKey = k;
      });
  }

  ngOnInit() {}
  ngOnDestroy(): void {
    this.comopnentDestroyed.next(true);
  }
  generateKey() {
    this.apikeyService
      .generate()
      .toPromise()
      .then((r) => {
        console.log(r);
      });
  }
}
