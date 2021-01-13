import { NgModule } from "@angular/core";
import {
  NbButtonModule,
  NbCardModule,
  NbIconModule,
  NbLayoutModule,
  NbMenuModule,
  NbSpinnerModule,
} from "@nebular/theme";

import { ThemeModule } from "../@theme/theme.module";
import { PagesComponent } from "./pages.component";
import { PagesRoutingModule } from "./pages-routing.module";
import { UserManagementComponent } from "./user-management/user-management.component";
import { MeterManagementComponent } from "./meter-management/meter-management.component";
import { TransactionsComponent } from "./transactions/transactions.component";
import { DashboardComponent } from "./dashboard/dashboard.component";
import { NbEvaIconsModule } from "@nebular/eva-icons";

@NgModule({
  imports: [
    PagesRoutingModule,
    ThemeModule,
    NbMenuModule,
    NbCardModule,
    ThemeModule,
    NbCardModule,
    NbButtonModule,
    NbEvaIconsModule,
    NbIconModule,
    ThemeModule,
    NbLayoutModule,
    NbSpinnerModule,
  ],
  declarations: [
    DashboardComponent,
    PagesComponent,
    UserManagementComponent,
    MeterManagementComponent,
    TransactionsComponent,
  ],
})
export class PagesModule {}
