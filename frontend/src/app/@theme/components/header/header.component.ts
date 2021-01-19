import { Component, OnDestroy, OnInit } from "@angular/core";
import {
  NbMediaBreakpointsService,
  NbMenuItem,
  NbMenuService,
  NbSidebarService,
  NbThemeService,
} from "@nebular/theme";

// import { UserData } from '../../../@core/data/users';
import { map, takeUntil } from "rxjs/operators";
import { Subject } from "rxjs";
import { AngularFireAuth } from "@angular/fire/auth";
import { UserService } from "app/pages/shared/services/user.service";
import firebase from "firebase/app";

@Component({
  selector: "ngx-header",
  styleUrls: ["./header.component.scss"],
  templateUrl: "./header.component.html",
})
export class HeaderComponent implements OnInit, OnDestroy {
  userPictureOnly: boolean = false;
  user: firebase.UserInfo;
  themes = [
    {
      value: "default",
      name: "Light",
    },
    {
      value: "dark",
      name: "Dark",
    },
    {
      value: "cosmic",
      name: "Cosmic",
    },
    {
      value: "corporate",
      name: "Corporate",
    },
  ];
  currentTheme = "default";
  userMenu: NbMenuItem[] = [
    {
      title: "Profile",
      icon: "person-outline",
    },
    {
      title: "Log out",
      icon: "log-out-outline",
    },
  ];
  private destroy$: Subject<void> = new Subject<void>();

  constructor(
    private sidebarService: NbSidebarService,
    private menuService: NbMenuService,
    private themeService: NbThemeService,
    private userService: UserService,
    // private nbAuth: NbAuthService,
    public auth: AngularFireAuth,
    private breakpointService: NbMediaBreakpointsService
  ) {}

  ngOnInit() {
    this.currentTheme = this.themeService.currentTheme;

    this.userService.user
      .pipe(takeUntil(this.destroy$))
      .subscribe((user: any) => (this.user = user));

    const { xl } = this.breakpointService.getBreakpointsMap();
    this.themeService
      .onMediaQueryChange()
      .pipe(
        map(([, currentBreakpoint]) => currentBreakpoint.width < xl),
        takeUntil(this.destroy$)
      )
      .subscribe(
        (isLessThanXl: boolean) => (this.userPictureOnly = isLessThanXl)
      );

    this.menuService
      .onItemClick()
      .pipe(takeUntil(this.destroy$))
      .subscribe((menuBag) => {
        switch (menuBag.item.title) {
          case "Log out":
            this.auth.signOut();
            return;

          default:
            break;
        }
      });
    this.themeService
      .onThemeChange()
      .pipe(
        map(({ name }) => name),
        takeUntil(this.destroy$)
      )
      .subscribe((themeName) => (this.currentTheme = themeName));
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }

  changeTheme(themeName: string) {
    this.themeService.changeTheme(themeName);
  }

  toggleSidebar(): boolean {
    this.sidebarService.toggle(true, "menu-sidebar");

    return false;
  }

  navigateHome() {
    this.menuService.navigateHome();
    return false;
  }
}
