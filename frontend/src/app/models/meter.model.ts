import { User } from "./user.model";

export interface Meter {
  meterNumber: String;
  ID: String;
  createdBy: User;
  updatedBy: User;
}
