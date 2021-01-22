import { Injectable } from "@angular/core";
import { ApolloQueryResult } from "@apollo/client/core";
import { Apollo } from "apollo-angular";
import {
  GetMetersQueryInput,
  MetersQuery,
  MetersQueryResult,
  NewMeterInput,
  NewMeterMutation,
} from "app/models/gql/meter.model";
import { Meter } from "app/models/meter.model";

@Injectable({
  providedIn: "root",
})
export class MeterService {
  constructor(private apollo: Apollo) {}
  getMeters(
    inputs: GetMetersQueryInput
  ): Promise<ApolloQueryResult<MetersQueryResult>> {
    return this.apollo
      .query<MetersQueryResult>({
        query: MetersQuery,
        variables: { inputs },
        fetchPolicy: "cache-first",
      })
      .toPromise();
  }
  createMeter(input: NewMeterInput) {
    return this.apollo
      .mutate<Meter>({
        mutation: NewMeterMutation,
        variables: { input },
      })
      .toPromise();
  }
}
