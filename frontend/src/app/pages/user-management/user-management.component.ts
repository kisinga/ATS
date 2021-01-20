import { Component, OnDestroy, OnInit } from "@angular/core";
import { NbDialogService } from "@nebular/theme";
import { Apollo, gql } from "apollo-angular";
import { GetUsersQueryInput } from "app/models/gql/user.query";
import { User } from "app/models/user.model";
import { StateService } from "app/pages/shared/services/state.service";
import { ReplaySubject } from "rxjs";
import { NewUserInput } from "app/models/gql/user.query";
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
  loading: Boolean = false;
  constructor(
    private state: StateService,
    private dialogService: NbDialogService,
    private userService: UserService
  ) {
    this.getUsers({});
  }

  ngOnInit() {}
  ngOnDestroy(): void {
    this.comopnentDestroyed.next(true);
  }
  openNewUserModal() {
    this.dialogService.open(NewUserComponent);
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
        this.getUsers({});
      });
  }
  getUsers(params: GetUsersQueryInput) {
    this.loading = true;
    this.userService.getUsers(params).then((r) => {
      this.users = r.data.users.data;
      this.loading = false;
    });
  }
  enableUser(email: string) {
    this.loadingUser = email;
    this.userService
      .enableUser(email)
      .toPromise()
      .then((t) => {
        if (this.loadingUser === email) {
          this.loadingUser = "";
        }
        this.getUsers({});
      });
  }
  generateKey() {}
}
