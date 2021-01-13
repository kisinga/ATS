import { Injectable } from "@angular/core";
import { User } from "app/models/user.model";
import { ReplaySubject } from "rxjs";
import { Apollo, gql } from "apollo-angular";
import { UserService } from "./user.service";
import { GetUsersQueryInput } from "app/models/gql/user.query";
import { APIKey } from "app/models/api-key.model";
import { ApiKeyService } from "./api-key.service";

@Injectable({
  providedIn: "root",
})
export class StateService {
  userManagementPage = new ReplaySubject<Number>(1);
  userManagementData = new ReplaySubject<User[]>(1);

  dashboardApiKey = new ReplaySubject<APIKey>(1);
  apikeyloading = new ReplaySubject<boolean>(1);
  constructor(
    private users: UserService,
    private apikeyService: ApiKeyService
  ) {
    this.setAPIKeyLoading(true);
    this.users.getUsers({}).then((r) => {
      this.userManagementData.next(r.data.users.data);
    });

    this.apikeyService.subscribeKey().subscribe((val) => {
      this.apikeyloading.next(false);
      this.dashboardApiKey.next(val.data.apiKeyChanged);
    });
  }
  setAPIKeyLoading(status: boolean) {
    this.apikeyloading.next(status);
  }
}
