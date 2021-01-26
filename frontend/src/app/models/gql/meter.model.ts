import {gql} from "apollo-angular";
import {Meter} from "../meter.model";
import {pageInfoFragment} from "./page-info.fragment";
import {PageInfoModel} from "./page-info.model";

export const MetersQuery = gql`
  query getMeters($limit: Int, $afterID: ID) {
    meters(limit: $limit, after: $afterID) {
      data {
        meterNumber
        ID
        updatedBy {
          name
          email
        }
        active
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

export const NewMeterMutation = gql`
  mutation createMeter($input: NewMeter!) {
    createMeter(input: $input) {
      meterNumber
      ID
      updatedBy {
        name
        email
      }
      active
      createdBy {
        name
        email
      }
    }
  }
`;

// export const userDataFragment = gql`

// `;
export const DisableMeterMutation = gql`
  mutation disableMeter($meterNumber: String!) {
    disableMeter(meterNumber: $meterNumber) {
      meterNumber
      ID
      updatedBy {
        name
        email
      }
      active
      createdBy {
        name
        email
      }
    }
  }
`;

export const EnableMeterMutation = gql`
  mutation enableMeter($meterNumber: String!) {
    enableMeter(meterNumber: $meterNumber) {
      meterNumber
      ID
      updatedBy {
        name
        email
      }
      active
      createdBy {
        name
        email
      }
    }
  }
`;

export interface GetMetersQueryInput {
  limit?: number;
  after?: string;
}

export interface NewMeterInput {
  meterNumber: string;
}

export interface MetersQueryModel {
  data?: Meter[];
  pageInfo?: PageInfoModel;
}

export interface MetersQueryResult {
  meters?: MetersQueryModel;
}
