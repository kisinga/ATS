import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";

import {PagesComponent} from "./pages.component";
import {DashboardComponent} from "./dashboard/dashboard.component";
import {MeterManagementComponent} from "./meter-management/meter-management.component";
import {UserManagementComponent} from "./user-management/user-management.component";
import {TokensComponent} from "./tokens/tokens.component";
import { ProfileComponent } from "./profile/profile.component";

const routes: Routes = [
  {
    path: "",
    component: PagesComponent,
    children: [
      {
        path: "dashboard",
        component: DashboardComponent,
      },
      {
        path: "profile",
        component: ProfileComponent,
      },
      {
        path: "tokens",
        component: TokensComponent,
      },
      {
        path: "users",
        component: UserManagementComponent,
      },
      {
        path: "meters",
        component: MeterManagementComponent,
      },
      {path: "", redirectTo: "dashboard"},
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class PagesRoutingModule {
}
