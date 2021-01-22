import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { AngularFireAuth } from "@angular/fire/auth";
import { Router } from "@angular/router";
import firebase from "firebase/app";
import { of, ReplaySubject } from "rxjs";
import { Apollo } from "apollo-angular";
import { catchError } from "rxjs/operators";
import dayjs from "dayjs";
import {
  DisableUserMutation,
  EnableUserMutation,
  GetUsersQueryInput,
  NewUserInput,
  NewUserMutation,
  UsersQuery,
  UsersQueryResult,
} from "app/models/gql/user.query";
import { ApolloQueryResult, FetchPolicy } from "@apollo/client/core";
import { User } from "app/models/user.model";

@Injectable({
  providedIn: "root",
})
export class UserService {
  user = new ReplaySubject<firebase.UserInfo>(1);

  constructor(
    private auth: AngularFireAuth,
    private router: Router,
    private apollo: Apollo,
    private http: HttpClient
  ) {
    this.auth.authState.subscribe((state) => {
      if (state) {
        // refresh if token doesnt exist of is expired
        const token = this.getToken();
        if (token === "" || this.tokenExpired(this.getTokenExpiry())) {
          state.getIdToken().then((t) => {
            this.fetchToken(state.email, t);
          });
        }
        this.user.next(state);
      } else {
        this.apollo.client.resetStore();
        this.destroyToken();
        if (router.routerState.snapshot.url !== "/login") {
          router.navigate(["/login"]);
        }
      }
    });
  }
  disableUser(email: string) {
    return this.apollo.mutate({
      mutation: DisableUserMutation,
      variables: { email },
    });
  }
  enableUser(email: string) {
    return this.apollo.mutate({
      mutation: EnableUserMutation,
      variables: { email },
    });
  }

  fetchToken(email: string, id: string) {
    this.http
      .post("https://atske.herokuapp.com/sessionInit", { id, email })
      .pipe(
        catchError((e) => {
          this.auth.signOut();
          if (e === "Not Authorised") {
            this.router.navigate(["/unauthorised"]);
          }
          return of({});
        })
      )
      .toPromise()
      .then((t) => {
        const token = JSON.parse(JSON.stringify(t));
        this.saveToken(token.Bearer, token.expiry);
        // console.log(token);
        this.router.navigate(["/"]);
      });
  }

  getToken(): String {
    return window.localStorage["token"] || "";
  }

  getTokenExpiry(): String {
    return window.localStorage["tokenExpiry"];
  }

  tokenExpired(expiry) {
    return dayjs().isAfter(dayjs(expiry));
  }

  saveToken(token: String, expiry: string) {
    window.localStorage["token"] = token;
    window.localStorage["tokenExpiry"] = expiry;
  }

  destroyToken() {
    window.localStorage.removeItem("token");
    window.localStorage.removeItem("tokenExpiry");
  }
  getUsers(
    inputs: GetUsersQueryInput,
    fetchPolicy: FetchPolicy
  ): Promise<ApolloQueryResult<UsersQueryResult>> {
    return this.apollo
      .query<UsersQueryResult>({
        query: UsersQuery,
        variables: { inputs },
        fetchPolicy,
      })
      .toPromise();
  }
  createUser(input: NewUserInput) {
    return this.apollo
      .mutate<User>({
        mutation: NewUserMutation,
        variables: { input },
      })
      .toPromise();
  }
}
