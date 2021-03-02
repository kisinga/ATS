import { NgModule } from "@angular/core";
import {
  NbAlertModule,
  NbButtonModule,
  NbCardModule,
  NbDialogModule,
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
import { DashboardComponent } from "./dashboard/dashboard.component";
import { NbEvaIconsModule } from "@nebular/eva-icons";
import { MatTableModule } from "@angular/material/table";
import { DateFromObjectIdPipe } from "app/pages/shared/pipes/date-from-object-id.pipe";
import { NewUserComponent } from "./user-management/dialogs/new-user/new-user.component";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { NewMeterComponent } from "./meter-management/dialogs/new-meter/new-meter.component";
import { TokensComponent } from "./tokens/tokens.component";
import { MatPaginatorModule } from "@angular/material/paginator";
import { ProfileComponent } from './profile/profile.component';
import { EditMeterComponent } from './meter-management/dialogs/edit-meter/edit-meter.component';

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
    MatPaginatorModule,
    NbInputModule,
    NbTooltipModule,
    FormsModule,
    ReactiveFormsModule,
    NbAlertModule,
    NbDialogModule.forRoot(),
  ],
  exports: [DateFromObjectIdPipe],
  declarations: [
    DashboardComponent,
    PagesComponent,
    UserManagementComponent,
    MeterManagementComponent,
    DateFromObjectIdPipe,
    NewUserComponent,
    NewMeterComponent,
    TokensComponent,
    ProfileComponent,
    EditMeterComponent,
  ],
  providers: [],
})
export class PagesModule {}
