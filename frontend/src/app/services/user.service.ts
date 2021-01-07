import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {AngularFireAuth} from '@angular/fire/auth';
import {Router} from '@angular/router';
import firebase from 'firebase/app';
import {of, ReplaySubject} from 'rxjs';
import {Apollo} from 'apollo-angular';
import { take } from 'rxjs/operators';
import dayjs from 'dayjs';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  user = new ReplaySubject<firebase.UserInfo>(1);

  constructor(private auth: AngularFireAuth, private router: Router, private apollo: Apollo, private http: HttpClient) {
    this.auth.authState.subscribe((state) => {
      if (state) {
        // refresh if token doesnt exist of is expired
        const token = this.getToken();
        if (token === '' || this.tokenExpired(this.getTokenExpiry())) {
          state.getIdToken().then(t => {
            this.fetchToken(state.email, t);
          });
        }
        this.user.next(state);
      } else {
        this.apollo.client.resetStore();
        this.destroyToken();
        if (router.routerState.snapshot.url !== '/login') {
          router.navigate(['/login']);
        }
      }
    });
  }
  fetchToken(email: string, id: string) {
    this.http.post('http://localhost:4242/sessionInit', {id, email})
    .pipe(
    catchError((e) => {
      console.log(e);
      this.auth.signOut();
      return of({});
    }))
    .toPromise()
    .then(t => {
      const token = JSON.parse(JSON.stringify(t));
      this.saveToken(token.Bearer, token.expiry);
      this.router.navigate(['/']);
    });
  }

  getToken(): String {
    return window.localStorage['token'] || '';
  }

  getTokenExpiry(): String {
    return window.localStorage['tokenExpiry'];
  }

  tokenExpired(expiry) {
     return (dayjs().isAfter(dayjs(expiry)));
  }

  saveToken(token: String, expiry: string) {
    window.localStorage['token'] = token;
    window.localStorage['tokenExpiry'] = expiry;
  }

  destroyToken() {
    window.localStorage.removeItem('token');
    window.localStorage.removeItem('tokenExpiry');
  }
}
