import gql from 'graphql-tag';

export const pageInfoFragment = gql`
  fragment pageInfoFragment on PageInfo {
    startCursor
    endCursor
    hasNextPage
  }
`;
