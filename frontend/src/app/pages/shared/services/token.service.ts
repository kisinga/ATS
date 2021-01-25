import { Injectable } from "@angular/core";
import { ApolloQueryResult, FetchPolicy } from "@apollo/client/core";
import { Apollo } from "apollo-angular";
import {
  GetTokensQueryInput,
  TokensQuery,
  TokensQueryModel,
  TokensQueryResult,
} from "app/models/gql/token.model";

@Injectable({
  providedIn: "root",
})
export class TokenService {
  constructor(private apollo: Apollo) {}
  getTokens(
    inputs: GetTokensQueryInput,
    fetchPolicy: FetchPolicy
  ): Promise<ApolloQueryResult<TokensQueryResult>> {
    return this.apollo
      .query<TokensQueryResult>({
        query: TokensQuery,
        variables: { inputs },
        fetchPolicy,
      })
      .toPromise();
  }
}
