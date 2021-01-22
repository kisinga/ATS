import { Injectable } from "@angular/core";
import { ApolloQueryResult, FetchPolicy } from "@apollo/client/core";
import { Apollo } from "apollo-angular";
import {
  DisableMeterMutation,
  EnableMeterMutation,
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
    inputs: GetMetersQueryInput,
    fetchPolicy: FetchPolicy
  ): Promise<ApolloQueryResult<MetersQueryResult>> {
    return this.apollo
      .query<MetersQueryResult>({
        query: MetersQuery,
        variables: { inputs },
        fetchPolicy,
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

  disableMeter(meterNumber: string) {
    return this.apollo.mutate({
      mutation: DisableMeterMutation,
      variables: { meterNumber },
    });
  }

  enableMeter(meterNumber: string) {
    return this.apollo.mutate({
      mutation: EnableMeterMutation,
      variables: { meterNumber },
    });
  }
}
