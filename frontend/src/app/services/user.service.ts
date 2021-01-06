import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AngularFireAuth } from '@angular/fire/auth';
import { Router } from '@angular/router';
import firebase from 'firebase/app';
import { ReplaySubject } from 'rxjs';
import { Apollo } from 'apollo-angular';
import { take } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  user = new ReplaySubject<firebase.UserInfo>(1);

  constructor(private auth: AngularFireAuth, router: Router,private apollo: Apollo, private http: HttpClient) {
    this.auth.authState.subscribe((state) => {
      if (state) {
        if (router.routerState.snapshot.url !== '/') {
          router.navigate(['/']);
        }
        // @TODO only fetch if after validatinng the either the token isnt there locally, or is expired
        state.getIdToken().then(t=>{
          this.fetchToken(t)
        })
        this.user.next(state);
      } else {
        this.apollo.client.resetStore();
        if (router.routerState.snapshot.url !== '/login') {
          router.navigate(['/login']);
        }
      }
    });
  }
  fetchToken(id: string) {
    this.http.post('http://localhost:4242/sessionInit', {id})
    .toPromise().then(()=>{})
  }

}
