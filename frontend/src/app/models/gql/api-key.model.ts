import gql from "graphql-tag";
import { APIKey } from "../api-key.model";

export const APIKeySubscription = gql`
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
