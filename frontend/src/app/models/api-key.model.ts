import { User } from "./user.model";

export interface APIKey {
  ID: String;
  createdBy: User;
}
