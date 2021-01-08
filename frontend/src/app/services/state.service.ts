import { Injectable } from "@angular/core";
import { User } from "app/models/user.model";
import { ReplaySubject } from "rxjs";
import { Apollo, gql } from "apollo-angular";
import { UserService } from "./user.service";
import { GetUsersQueryInput } from "app/models/gql/user.query";

@Injectable({
  providedIn: "root",
})
export class StateService {
  userManagementPage = new ReplaySubject<Number>(1);
  userManagementData = new ReplaySubject<User[]>(1);

  dashboardApiKey = new ReplaySubject<string>(1);
  constructor(private users: UserService) {
    this.users.getUsers({}).then((r) => {
      this.userManagementData.next(r.data.users.data);
    });
  }
}
