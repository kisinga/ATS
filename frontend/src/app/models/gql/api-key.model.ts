import {gql, Subscription} from "apollo-angular";
import {APIKey} from "../api-key.model";
import {Injectable} from "@angular/core";

@Injectable({
  providedIn: "root",
})
export class APIKeySubscription extends Subscription {
  document = APIKeySubscriptionQuery;
}

export const APIKeyGenerate = gql`
  mutation generateAPIKey {
    generateAPIKey {
      ID
      createdBy {
        email
      }
    }
  }
`;

export const APIKeySubscriptionQuery = gql`
  subscription apiKeyChanged {
    apiKeyChanged {
      ID
      createdBy {
        email
      }
    }
  }
`;
export const APIKeyQuery = gql`
  query currentAPIKey {
    currentAPIKey {
      ID
      createdBy {
        email
      }
    }
  }
`;

export interface APIKeySubscriptionResult {
  generateAPIKey?: APIKey;
}

export interface APIKeyGenerateResult {
  generateAPIKey?: APIKey;
}

export interface APIKeyQueryResult {
  currentAPIKey?: APIKey;
}
