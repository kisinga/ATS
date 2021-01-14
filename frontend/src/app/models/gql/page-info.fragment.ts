import { Apollo, gql } from "apollo-angular";

export const pageInfoFragment = gql`
  fragment pageInfoFragment on PageInfo {
    startCursor
    endCursor
    hasNextPage
  }
`;
