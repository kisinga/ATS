import { Injectable } from "@angular/core";
import { User } from "app/models/user.model";
import { ReplaySubject } from "rxjs";
import { Apollo, gql } from "apollo-angular";

@Injectable({
  providedIn: "root",
})
export class StateService {
  userManagementPage = new ReplaySubject<Number>(1);
  userManagementData = new ReplaySubject<Array<User>>(1);

  dashboardApiKey = new ReplaySubject<string>(1);
  constructor(private apollo: Apollo) {
    this.apollo
      .watchQuery<{ users: Array<User> }>({
        query: gql`
          {
            users {
              ID
              email
            }
          }
        `,
      })
      .valueChanges.subscribe((result) => {
        this.userManagementData.next(result.data.users);
      });
  }
}
