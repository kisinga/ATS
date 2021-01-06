import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AngularFireAuth } from '@angular/fire/auth';
import { Router } from '@angular/router';
import firebase from 'firebase/app';
import { ReplaySubject } from 'rxjs';
@Injectable({
  providedIn: 'root',
})
export class UserService {
  user = new ReplaySubject<firebase.UserInfo>(1);

  constructor(private auth: AngularFireAuth, router: Router, private http: HttpClient) {
    this.auth.authState.subscribe((state) => {
      if (state) {
        if (router.routerState.snapshot.url !== '/') {
          router.navigate(['/']);
        }
        this.user.next(state);
        this.http.get()
      } else {
        if (router.routerState.snapshot.url !== '/login') {
          router.navigate(['/login']);
        }
      }
    });
  }
 private fetchCookie() {

 }

}
