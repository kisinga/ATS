import { NgModule } from "@angular/core";
import {
  NbAlertModule,
  NbButtonModule,
  NbCardModule,
  NbIconModule,
  NbInputModule,
  NbLayoutModule,
  NbMenuModule,
  NbSpinnerModule,
  NbTooltipModule,
} from "@nebular/theme";

import { ThemeModule } from "../@theme/theme.module";
import { PagesComponent } from "./pages.component";
import { PagesRoutingModule } from "./pages-routing.module";
import { UserManagementComponent } from "./user-management/user-management.component";
import { MeterManagementComponent } from "./meter-management/meter-management.component";
import { TransactionsComponent } from "./transactions/transactions.component";
import { DashboardComponent } from "./dashboard/dashboard.component";
import { NbEvaIconsModule } from "@nebular/eva-icons";
import { MatTableModule } from "@angular/material/table";
import { DateFromObjectIdPipe } from "app/pages/shared/pipes/date-from-object-id.pipe";
import { NewUserComponent } from "./user-management/dialogs/new-user/new-user.component";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { NewMeterComponent } from './meter-management/dialogs/new-meter/new-meter.component';

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
    MatTableModule,
    NbInputModule,
    NbTooltipModule,
    FormsModule,
    ReactiveFormsModule,
    NbAlertModule,
  ],
  exports: [DateFromObjectIdPipe],
  declarations: [
    DashboardComponent,
    PagesComponent,
    UserManagementComponent,
    MeterManagementComponent,
    TransactionsComponent,
    DateFromObjectIdPipe,
    NewUserComponent,
    NewMeterComponent,
  ],
  providers: [],
})
export class PagesModule {}
