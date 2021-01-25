export interface Token {
  meterNumber: string;
  tokenString: string;
  ID: string;
  status: TokenStatus;
  apiKey: string;
}

enum TokenStatus {
  New,
  Sent,
  Error,
  Loading,
  Loaded,
}
