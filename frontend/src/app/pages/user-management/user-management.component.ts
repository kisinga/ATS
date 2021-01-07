import { Component, OnDestroy, OnInit } from "@angular/core";
import { Apollo, gql } from "apollo-angular";
import { User } from "app/models/user.model";
import { StateService } from "app/services/state.service";
import { ReplaySubject } from "rxjs";
import { takeUntil } from "rxjs/operators";

@Component({
  templateUrl: "./user-management.component.html",
  styleUrls: ["./user-management.component.scss"],
})
export class UserManagementComponent implements OnInit, OnDestroy {
  users: Array<User>;
  comopnentDestroyed: ReplaySubject<boolean> = new ReplaySubject<boolean>();

  constructor(private state: StateService) {
    this.state.userManagementData
      .pipe(takeUntil(this.comopnentDestroyed))
      .subscribe((k) => {
        this.users = k;
      });
  }

  ngOnInit() {}
  ngOnDestroy(): void {
    this.comopnentDestroyed.next(true);
  }
  generateKey() {}
}
