import { RouterModule, Routes } from "@angular/router";
import { NgModule } from "@angular/core";

import { PagesComponent } from "./pages.component";
import { DashboardComponent } from "./dashboard/dashboard.component";
import { TransactionsComponent } from "./transactions/transactions.component";
import { MeterManagementComponent } from "./meter-management/meter-management.component";
import { UserManagementComponent } from "./user-management/user-management.component";

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
        path: "transactions",
        component: TransactionsComponent,
      },
      {
        path: "users",
        component: UserManagementComponent,
      },
      {
        path: "meters",
        component: MeterManagementComponent,
      },
      { path: "", redirectTo: "dashboard" },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class PagesRoutingModule {}
