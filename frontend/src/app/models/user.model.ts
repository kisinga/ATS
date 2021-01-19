export interface User {
  name: string;
  email: string;
  createdBy: User;
  updatedBy: User;
  ID: string;
}
