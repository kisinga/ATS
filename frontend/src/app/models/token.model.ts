export interface Token {
  meterNumber: string;
  tokenString: string;
  ID: string;
  status: TokenStatus;
  apiKey: string;
}

export enum TokenStatus {
  New,
  Sent,
  Error,
  Loading,
  Loaded,
}
