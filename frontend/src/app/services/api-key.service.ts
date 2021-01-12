import { Injectable } from "@angular/core";
import { FetchResult } from "@apollo/client/core";
import { Apollo } from "apollo-angular";
import {
  APIKeSubscription,
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
      query: APIKeSubscription,
    });
  }
}
