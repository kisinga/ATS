import { of as observableOf, Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { Contacts, RecentUsers, UserData } from '../data/users';

@Injectable()
export class UserService extends UserData {

  private time: Date = new Date;

  private users = {
    kisinga: { name: 'Kamana Kisinga', picture: 'assets/images/kisinga.jpg' },
  };
  private types = {
    mobile: 'mobile',
    home: 'home',
    work: 'work',
  };
  private contacts: Contacts[] = [
    { user: this.users.kisinga, type: this.types.mobile },
  ];
  private recentUsers: RecentUsers[] = [
    { user: this.users.kisinga, type: this.types.home, time: this.time.setHours(21, 12) },
  ];

  getUsers(): Observable<any> {
    return observableOf(this.users);
  }

  getContacts(): Observable<Contacts[]> {
    return observableOf(this.contacts);
  }

  getRecentUsers(): Observable<RecentUsers[]> {
    return observableOf(this.recentUsers);
  }
}
