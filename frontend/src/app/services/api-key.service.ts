import { Injectable } from "@angular/core";
import { ApolloQueryResult, FetchResult } from "@apollo/client/core";
import { Apollo } from "apollo-angular";
import { APIKey } from "app/models/api-key.model";
import {
  APIKeyGenerate,
  APIKeyGenerateResult,
  APIKeySubscription,
  APIKeySubscriptionResult,
} from "app/models/gql/api-key.model";
import { Observable } from "rxjs";

@Injectable({
  providedIn: "root",
})
export class ApiKeyService {
  constructor(private apollo: Apollo) {}
  subscribeKey(): Observable<FetchResult<APIKeySubscriptionResult>> {
    return this.apollo.subscribe<APIKeySubscriptionResult>({
      query: APIKeySubscription,
    });
  }

  generate(): Observable<FetchResult<APIKeyGenerateResult>> {
    return this.apollo.mutate<APIKeyGenerateResult>({
      mutation: APIKeyGenerate,
    });
  }
}
