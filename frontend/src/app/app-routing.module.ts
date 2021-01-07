import { ExtraOptions, RouterModule, Routes } from "@angular/router";
import { NgModule } from "@angular/core";
import {
  AngularFireAuthGuard,
  redirectLoggedInTo,
  redirectUnauthorizedTo,
} from "@angular/fire/auth-guard";
import { LoginComponent } from "./auth/login/login.component";

const redirectUnauthorizedToLogin = () => redirectUnauthorizedTo(["/login"]);
const redirectLoggedInToDashboard = () => redirectLoggedInTo([""]);

export const routes: Routes = [
  {
    path: "",
    canActivate: [AngularFireAuthGuard], // here we tell Angular to check the access with our AuthGuard
    data: { authGuardPipe: redirectUnauthorizedToLogin },
    loadChildren: () =>
      import("./pages/pages.module").then((m) => m.PagesModule),
  },
  {
    path: "login",
    component: LoginComponent,
    canActivate: [AngularFireAuthGuard],
    data: { authGuardPipe: redirectLoggedInToDashboard },
  },
  {
    path: "unauthorised",
    component: LoginComponent,
    canActivate: [AngularFireAuthGuard],
  },
  { path: "**", redirectTo: "" },
];

const config: ExtraOptions = {
  useHash: false,
};

@NgModule({
  imports: [RouterModule.forRoot(routes, config)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
