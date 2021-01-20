import { Injectable } from "@angular/core";
import { Apollo } from "apollo-angular";

@Injectable({
  providedIn: "root",
})
export class MeterService {
  constructor(private apollo: Apollo) {}
  // getMeters(
  //   inputs: GetUsersQueryInput
  // ): Promise<ApolloQueryResult<UsersQueryResult>> {
  //   return this.apollo
  //     .query<UsersQueryResult>({
  //       query: UsersQuery,
  //       variables: { inputs },
  //       fetchPolicy: "cache-first",
  //     })
  //     .toPromise();
  // }
  // createMeter(input: NewUserInput) {
  //   return this.apollo
  //     .mutate<User>({
  //       mutation: NewUserMutation,
  //       variables: { input },
  //     })
  //     .toPromise();
  // }
}
