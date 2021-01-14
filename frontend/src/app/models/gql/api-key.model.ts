import { Apollo, gql } from "apollo-angular";
import { APIKey } from "../api-key.model";
import { Subscription } from "apollo-angular";
import { Injectable } from "@angular/core";

@Injectable({
  providedIn: "root",
})
export class APIKeySubscription extends Subscription {
  document = gql`
    subscription apiKeyChanged {
      apiKeyChanged {
        ID
        createdBy {
          email
          name
        }
      }
    }
  `;
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

export interface APIKeySubscriptionResult {
  apiKeyChanged?: APIKey;
}

export interface APIKeyGenerateResult {
  generateAPIKey?: APIKey;
}
