import gql from "graphql-tag";
import { APIKey } from "../api-key.model";

export const APIKeSubscription = gql`
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

export interface APIKeySubscriptionResult {
  apiKeyChanged?: APIKey;
}
