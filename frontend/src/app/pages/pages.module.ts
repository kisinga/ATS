import { NgModule } from "@angular/core";
import { NbCardModule, NbMenuModule } from "@nebular/theme";

import { ThemeModule } from "../@theme/theme.module";
import { PagesComponent } from "./pages.component";
import { DashboardModule } from "./dashboard/dashboard.module";
import { PagesRoutingModule } from "./pages-routing.module";
import { UserManagementComponent } from "./user-management/user-management.component";
import { MeterManagementComponent } from "./meter-management/meter-management.component";
import { TransactionsComponent } from "./transactions/transactions.component";

@NgModule({
  imports: [
    PagesRoutingModule,
    ThemeModule,
    NbMenuModule,
    DashboardModule,
    NbCardModule,
    ThemeModule,
  ],
  declarations: [
    PagesComponent,
    UserManagementComponent,
    MeterManagementComponent,
    TransactionsComponent,
  ],
})
export class PagesModule {}
