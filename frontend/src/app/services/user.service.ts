import { Injectable } from '@angular/core';
import { AngularFireAuth } from '@angular/fire/auth';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private auth: AngularFireAuth,
    router: Router
  ) {
    auth.authState.subscribe(state => {

      if (state) {

      } else {
        if (router.routerState.snapshot.url !== "/login") {
          router.navigate(["/login"]);
        }
      }
    });
  }
}
