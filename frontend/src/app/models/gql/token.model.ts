import { Apollo, gql } from "apollo-angular";
import { Token } from "../token.model";
import { pageInfoFragment } from "./page-info.fragment";
import { PageInfoModel } from "./page-info.model";

export const TokensQuery = gql`
  query getTokens(
    $limit: Int = 10
    $beforeOrAfter: ID
    $reversed: Boolean
    $meterNumber: String
  ) {
    getTokens(
      limit: $limit
      beforeOrAfter: $beforeOrAfter
      reversed: $reversed
      meterNumber: $meterNumber
    ) {
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
  beforeOrAfter?: string;
  reversed?: boolean;
  meterNumber?: string;
}

export interface TokensQueryModel {
  data?: Token[];
  pageInfo?: PageInfoModel;
}

export interface TokensQueryResult {
  getTokens?: TokensQueryModel;
}
