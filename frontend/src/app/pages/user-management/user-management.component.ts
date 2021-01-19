import { Component, OnDestroy, OnInit } from "@angular/core";
import { NbDialogService } from "@nebular/theme";
import { Apollo, gql } from "apollo-angular";
import { User } from "app/models/user.model";
import { StateService } from "app/pages/shared/services/state.service";
import { ReplaySubject } from "rxjs";
import { takeUntil } from "rxjs/operators";
import { DateFromObjectIdPipe } from "../shared/pipes/date-from-object-id.pipe";
import { ApiKeyService } from "../shared/services/api-key.service";
import { UserService } from "../shared/services/user.service";
import { NewUserComponent } from "./dialogs/new-user/new-user.component";

@Component({
  templateUrl: "./user-management.component.html",
  styleUrls: ["./user-management.component.scss"],
})
export class UserManagementComponent implements OnInit, OnDestroy {
  users: Array<User>;
  comopnentDestroyed: ReplaySubject<boolean> = new ReplaySubject<boolean>();
  displayedColumns: string[] = ["email", "name", "createdby", "date", "delete"];
  loadingUser = "";
  constructor(
    private state: StateService,
    private dialogService: NbDialogService,
    private userService: UserService
  ) {
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
  openNewUserModal() {
    this.dialogService.open(NewUserComponent).onClose.subscribe((user) => {});
  }
  disableUser(email: string) {
    this.loadingUser = email;
    this.userService
      .disableUser(email)
      .toPromise()
      .then((t) => {
        if (this.loadingUser === email) {
          this.loadingUser = "";
        }
      });
  }
  enableUser(email: string) {
    this.loadingUser = email;
    this.userService
      .enableUser(email)
      .toPromise()
      .then((t) => {
        // console.log(t);
        if (this.loadingUser === email) {
          this.loadingUser = "";
        }
      });
  }
  generateKey() {}
}
