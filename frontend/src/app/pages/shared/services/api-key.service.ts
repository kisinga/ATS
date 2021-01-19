import { Injectable } from "@angular/core";
import { ApolloQueryResult, FetchResult } from "@apollo/client/core";
import { Apollo, QueryRef } from "apollo-angular";
import { APIKey } from "app/models/api-key.model";
import {
  APIKeyGenerate,
  APIKeyGenerateResult,
  APIKeyQuery,
  APIKeyQueryResult,
  APIKeySubscription,
  APIKeySubscriptionQuery,
  APIKeySubscriptionResult,
} from "app/models/gql/api-key.model";
import { Observable } from "rxjs";

@Injectable({
  providedIn: "root",
})
export class ApiKeyService {
  currentWatcher: QueryRef<APIKeyQueryResult>;

  constructor(private apollo: Apollo) {
    this.currentWatcher = this.apollo.watchQuery<APIKeyQueryResult>({
      query: APIKeyQuery,
    });
  }

  current(): Observable<ApolloQueryResult<APIKeyQueryResult>> {
    return this.currentWatcher.valueChanges;
  }

  getCurrent() {
    return this.apollo.query({
      query: APIKeyQuery,
    });
  }

  generate(): Observable<FetchResult<APIKeyGenerateResult>> {
    return this.apollo.mutate<APIKeyGenerateResult>({
      mutation: APIKeyGenerate,
    });
  }
  refetch() {
    return this.currentWatcher.refetch();
  }
}
