import { Apollo, gql } from "apollo-angular";
import { User } from "../user.model";
import { pageInfoFragment } from "./page-info.fragment";
import { PageInfoModel } from "./page-info.model";

export const UsersQuery = gql`
  query getUsers($limit: Int, $afterID: ID) {
    users(limit: $limit, after: $afterID) {
      data {
        ID
        name
        email
      }
      pageInfo {
        ...pageInfoFragment
      }
    }
  }
  ${pageInfoFragment}
`;

export interface GetUsersQueryInput {
  limit?: number;
  after?: string;
}

export interface UsersQueryModel {
  data?: User[];
  pageInfo?: PageInfoModel;
}

export interface UsersQueryResult {
  users?: UsersQueryModel;
}
