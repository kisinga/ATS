import { Component, OnDestroy, OnInit } from "@angular/core";
import { ReplaySubject } from "rxjs";
import { takeUntil } from "rxjs/operators";
import { UserService } from "../shared/services/user.service";
import firebase from "firebase/app";

@Component({
  templateUrl: "./profile.component.html",
  styleUrls: ["./profile.component.scss"],
})
export class ProfileComponent implements OnInit, OnDestroy {
  user: firebase.UserInfo;
  comopnentDestroyed: ReplaySubject<boolean> = new ReplaySubject<boolean>();

  constructor(private userService: UserService) {
    this.userService.user
      .pipe(takeUntil(this.comopnentDestroyed))
      .subscribe((user) => {
        this.user = user;
      });
  }

  ngOnDestroy(): void {
    this.comopnentDestroyed.next(true);
  }

  ngOnInit(): void {}
}
