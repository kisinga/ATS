import { Apollo, gql } from "apollo-angular";
import { Token } from "../token.model";
import { pageInfoFragment } from "./page-info.fragment";
import { PageInfoModel } from "./page-info.model";

export const TokensQuery = gql`
  query getTokens($limit: Int, $afterID: ID, $meterNumber: String) {
    getTokens(limit: $limit, after: $afterID, meterNumber: $meterNumber) {
      data {
        meterNumber
        tokenString
        ID
        status
        apiKey
      }
      pageInfo {
        ...pageInfoFragment
      }
    }
  }
  ${pageInfoFragment}
`;

export interface GetTokensQueryInput {
  limit?: number;
  after?: string;
}

export interface TokensQueryModel {
  data?: Token[];
  pageInfo?: PageInfoModel;
}

export interface TokensQueryResult {
  getTokens?: TokensQueryModel;
}
