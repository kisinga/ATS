import { Injectable } from '@angular/core';
import { User } from 'app/models/user.model';
import { ReplaySubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class StateService {
  userManagementState =new ReplaySubject<{
    page: number,
    data: Array<User>
  }>(1)
  constructor() { }
}
