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
        disabled
        createdBy {
          name
          email
        }
      }
      pageInfo {
        ...pageInfoFragment
      }
    }
  }
  ${pageInfoFragment}
`;

export const NewUserMutation = gql`
  mutation createUser($input: NewUser!) {
    createUser(input: $input) {
      ID
      name
      email
      disabled
      createdBy {
        name
        email
      }
    }
  }
`;

// export const userDataFragment = gql`

// `;
export const DisableUserMutation = gql`
  mutation disableUser($email: String!) {
    disableUser(email: $email) {
      ID
      name
      email
      disabled
      createdBy {
        name
        email
      }
    }
  }
`;

export const EnableUserMutation = gql`
  mutation enableUser($email: String!) {
    enableUser(email: $email) {
      ID
      name
      email
      disabled
      createdBy {
        name
        email
      }
    }
  }
`;

export interface GetUsersQueryInput {
  limit?: number;
  after?: string;
}
export interface NewUserInput {
  name: string;
  email: string;
}

export interface UsersQueryModel {
  data?: User[];
  pageInfo?: PageInfoModel;
}

export interface UsersQueryResult {
  users?: UsersQueryModel;
}
